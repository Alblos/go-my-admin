package connections

import (
	"github.com/go-my-admin/server/database"
	"github.com/go-my-admin/server/database/types"
)

var allConections = make(map[int]database.DBConnection)

// GetConnectionOrCreateIfNotExists Returns the connection with the given ID if it exists (creates it if it does not exist)
func GetConnectionOrCreateIfNotExists(id int) (database.DBConnection, error) {
	if allConections[id].Cnx == nil { // If it does not exist
		err := createConnection(id)
		if err != nil {
			return database.DBConnection{}, err
		}
	}
	return allConections[id], nil
}

// createConnection Creates the connection with the given ID if it does not exist
func createConnection(id int) error {
	db := database.InternalDb

	query, err := db.RunQueryWithParams("SELECT id, common_name, database_name, host, port, username, ssl_mode, password FROM connections WHERE id = $1", id)
	if err != nil {
		return err
	}

	var connectionMapped types.Connection

	for query.Next() {
		err = query.Scan(&connectionMapped.Id, &connectionMapped.CommonName, &connectionMapped.DatabaseName, &connectionMapped.Host, &connectionMapped.Port, &connectionMapped.Username, &connectionMapped.SslMode, &connectionMapped.Password)
		if err != nil {
			return err
		}
	}

	var cnx database.DBConnection

	err = cnx.Connect(database.GetConnectionString(connectionMapped))
	if err != nil {
		return err
	}

	allConections[id] = cnx
	return nil
}
