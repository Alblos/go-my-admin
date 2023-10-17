package connections

import (
	"github.com/go-my-admin/server/database"
)

// CheckIfConnectionWithIdExists checks if a connection with the given id exists
func CheckIfConnectionWithIdExists(id int) (bool, error) {
	db := database.InternalDb.Cnx

	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM connections WHERE id = $1)", id).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
