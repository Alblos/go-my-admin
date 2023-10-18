package database

import (
	"fmt"
	"github.com/go-my-admin/server/database/types"
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

// GetConnectionString returns the connection string for the given connection
func GetConnectionString(connectionData types.Connection) string {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=",
		connectionData.Username,
		connectionData.Password,
		connectionData.Host,
		connectionData.Port,
		connectionData.DatabaseName)

	if connectionData.SslMode == "true" {
		connectionString += "require"
	} else {
		connectionString += "disable"
	}

	return connectionString
}
