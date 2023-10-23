package bootstrap

import (
	"github.com/go-my-admin/server/database"
	"github.com/go-my-admin/server/logger"
)

// BootInternalDb bootstraps the internal database
func BootInternalDb() error {
	internalDbObject := &database.InternalDb
	err := internalDbObject.Connect(database.GetConnectionStringInternalDb())
	if err != nil {
		logger.Error("Error connecting to internal database: ", err)
		return err
	}

	err = internalDbObject.RunSqlFile("server/database/rawSql/createTablesInternalDB.sql")
	if err != nil {
		logger.Error("Error running createTablesInternalDB.sql: ", err)
		return err
	}

	return nil
}
