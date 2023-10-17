package database

import (
	"github.com/go-my-admin/server/logger"
	"os"
)

// GetConnectionStringInternalDb returns the connection string for the internal database
func GetConnectionStringInternalDb() string {
	connectionString := os.Getenv("DB_CONNECTION_STRING")

	if connectionString == "" {
		logger.Error("DB_CONNECTION_STRING not set", nil)
		return ""
	}

	return connectionString
}
