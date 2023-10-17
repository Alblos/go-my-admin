package database

import (
	"database/sql"
	"fmt"
	"github.com/go-my-admin/server/logger"
	_ "github.com/lib/pq" // Postgres driver
	"os"
)

var DBInstance *sql.DB

// Connect to the database
func Connect() (err error) {
	connectionString := os.Getenv("DB_CONNECTION_STRING")

	if connectionString == "" {
		fmt.Println("Error: DB_CONNECTION_STRING not set")
		return fmt.Errorf("DB_CONNECTION_STRING not set")
	}

	DBInstance, err = sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Println("Error opening database: ", err)
		return err
	}

	err = DBInstance.Ping()
	if err != nil {
		fmt.Println("Error pinging database, connection not established: ", err)
		return err
	}

	logger.Info("Connected to internal database")
	return nil
}
