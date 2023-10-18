package connections

// GetFullDbSchema returns a map of tables and columns in the given database
func GetFullDbSchema(connectionId int) (map[string][]Column, error) {
	var schema = make(map[string][]Column)

	tables, err := GetTablesInDatabase(connectionId)
	if err != nil {
		return nil, err
	}

	for _, table := range tables {
		columns, err := GetColumnsInTable(connectionId, table)
		if err != nil {
			return nil, err
		}
		schema[table] = columns
	}

	return schema, nil
}

// GetTablesInDatabase returns a list of tables in the given database
func GetTablesInDatabase(connectionId int) ([]string, error) {
	var tables []string

	db, err := GetConnectionOrCreateIfNotExists(connectionId)

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

// GetColumnsInTable returns a list of columns in the given table
func GetColumnsInTable(connectionId int, table string) ([]Column, error) {
	var columns = make([]Column, 0)

	db, err := GetConnectionOrCreateIfNotExists(connectionId)

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
