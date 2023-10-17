package database

import (
	"database/sql"
	"os"
)

// RunSqlFile runs a SQL file
func (DbInstance *DBConnection) RunSqlFile(filePath string) error {
	// Read file
	file, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	// Run SQL
	_, err = DbInstance.Cnx.Exec(string(file))
	if err != nil {
		return err
	}

	return nil
}

// RunRawQuery runs a raw SQL query
func (DbInstance *DBConnection) RunRawQuery(query string) (rows *sql.Rows, err error) {
	rows, err = DbInstance.Cnx.Query(query)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

// RunQueryWithParams runs a SQL query with params
func (DbInstance *DBConnection) RunQueryWithParams(query string, params ...interface{}) (rows *sql.Rows, err error) {
	rows, err = DbInstance.Cnx.Query(query, params...)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
