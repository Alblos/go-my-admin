package connections

type SchemaQueryResponse struct {
	TableCatalog     string
	TableSchema      string
	TableName        string
	TableType        string
	IsInsertableInto string
	IsTyped          string
}

// GetTablesInDatabase returns a list of tables in the given database
func GetTablesInDatabase(connectionId int) ([]SchemaQueryResponse, error) {
	var tables []SchemaQueryResponse

	db, err := GetConnectionOrCreateIfNotExists(connectionId)

	rows, err := db.Cnx.Query("SELECT table_catalog, table_schema, table_name, table_type, is_insertable_into, is_typed FROM information_schema.tables WHERE table_schema = 'public'")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var table SchemaQueryResponse
		err = rows.Scan(&table.TableCatalog, &table.TableSchema, &table.TableName, &table.TableType, &table.IsInsertableInto, &table.IsTyped)
		if err != nil {
			return nil, err
		}
		tables = append(tables, table)
	}

	return tables, nil
}
