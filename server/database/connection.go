package database

import (
	"database/sql"
	"github.com/go-my-admin/server/logger"
	_ "github.com/lib/pq" // Postgres driver
)

type DBConnection struct {
	cnx *sql.DB
}

// Connect to the database
func (DbInstance *DBConnection) Connect(connectionString string) (err error) {
	if connectionString == "" {
		logger.Error("DB_CONNECTION_STRING not set", nil)
		return
	}

	DbInstance.cnx, err = sql.Open("postgres", connectionString)
	if err != nil {
		logger.Error("Error opening database connection: ", err)
		return err
	}

	err = DbInstance.cnx.Ping()
	if err != nil {
		logger.Error("Error pinging database: ", err)
		return err
	}

	logger.Info("Connected to internal database")
	return nil
}
