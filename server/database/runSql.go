package database

import (
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
	_, err = DbInstance.cnx.Exec(string(file))
	if err != nil {
		return err
	}

	return nil
}
