package connections

import (
	"database/sql"
	"encoding/json"
	"github.com/go-my-admin/server/cache"
	"github.com/go-my-admin/server/database"
	"github.com/go-my-admin/server/logger"
	"github.com/redis/go-redis/v9"
	"strconv"
	"sync"
)

// GetFullDbSchema returns a map of tables and columns in the given database
func GetFullDbSchema(connectionId int, useCache bool) (map[string][]Column, error) {

	if useCache {
		exists, schema, err := GetSchemaFromCache(connectionId)
		if err != nil {
			return nil, err
		} else if exists {
			return schema, nil
		}
	}

	schema := make(map[string][]Column)

	db, err := GetConnectionReplica(connectionId)
	defer func(Cnx *sql.DB) {
		err := Cnx.Close()
		if err != nil {
			logger.Error("Could not close connection to database", err)
		}
	}(db.Cnx)

	tables, err := getTablesInDatabase(&db)
	if err != nil {
		return nil, err
	}

	// Create a channel
	columnsChannel := make(chan map[string][]Column)

	var wg = sync.WaitGroup{}

	for _, table := range tables {
		wg.Add(1)
		go func(table string) {
			defer wg.Done()
			columns, err := getColumnsInTable(table, &db)
			if err != nil {
				logger.Error("Could not get columns in table", err)
				return
			}
			columnsChannel <- map[string][]Column{table: columns}
		}(table)
	}
	go func() {
		wg.Wait()
		close(columnsChannel)
	}()

	for columns := range columnsChannel {
		for table, column := range columns {
			schema[table] = column
		}
	}

	SaveSchemaToCache(schema, connectionId)
	return schema, nil
}

// getTablesInDatabase returns a list of tables in the given database
func getTablesInDatabase(db *database.DBConnection) ([]string, error) {
	var tables []string

	rows, err := db.Cnx.Query("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var table string
		err = rows.Scan(&table)
		if err != nil {
			return nil, err
		}
		tables = append(tables, table)
	}

	return tables, nil
}

type Column struct {
	Name string
	Type string
}

// getColumnsInTable returns a list of columns in the given table
func getColumnsInTable(table string, db *database.DBConnection) ([]Column, error) {
	var columns = make([]Column, 0)

	rows, err := db.Cnx.Query("SELECT column_name, data_type FROM information_schema.columns WHERE table_name = $1", table)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var oneColumn Column
		err = rows.Scan(&oneColumn.Name, &oneColumn.Type)
		if err != nil {
			return nil, err
		}
		columns = append(columns, Column{
			Name: oneColumn.Name,
			Type: oneColumn.Type,
		})
	}

	return columns, nil
}

// SaveSchemaToCache saves the schema of a database in the internal cache
func SaveSchemaToCache(schema map[string][]Column, connectionId int) {
	key := "schema:" + strconv.Itoa(connectionId)
	jsonSchema, err := json.Marshal(schema)
	if err != nil {
		logger.Error("Could not marshal schema to JSON", err)
		return
	}
	cache.InternalCache.Cnx.Set(cache.InternalCache.Ctx, key, jsonSchema, 0)
}

// GetSchemaFromCache returns the schema of a database from the internal cache
func GetSchemaFromCache(connectionId int) (exists bool, schema map[string][]Column, err error) {
	var jsonSchema []byte
	key := "schema:" + strconv.Itoa(connectionId)
	err = cache.InternalCache.Cnx.Get(cache.InternalCache.Ctx, key).Scan(&jsonSchema)
	if err != nil {
		if err == redis.Nil {
			return false, nil, nil
		} else {
			return false, nil, err
		}
	}

	err = json.Unmarshal(jsonSchema, &schema)

	return true, schema, nil
}
