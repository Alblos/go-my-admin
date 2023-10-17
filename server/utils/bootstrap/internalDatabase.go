package bootstrap

import (
	"github.com/go-my-admin/server/database"
	"github.com/go-my-admin/server/logger"
)

// BootInternalDb bootstraps the internal database
func BootInternalDb(internalDbObject *database.DBConnection) error {
	err := internalDbObject.Connect(database.GetConnectionStringInternalDb())
	if err != nil {
		logger.Error("Error connecting to internal database: ", err)
		return err
	}

	return nil
}
