package connections

import (
	"github.com/gin-gonic/gin"
	"github.com/go-my-admin/server/connections"
	"github.com/go-my-admin/server/database"
	"github.com/go-my-admin/server/database/types"
	"strconv"
)

// HandleGetAllConnections
// @BasePath /
// @Summary Get all connections
// @Description Get all connections
// @Tags connections
// @Produce json
// @Success 200 {object} object "Returns the connections"
// @Failure 500 {object} object "Internal error"
// @Router /connections [get]
func HandleGetAllConnections(c *gin.Context) {
	db := database.InternalDb

	connectionsInDb, err := db.RunRawQuery("SELECT id, common_name, database_name, host, port, username, ssl_mode FROM connections")
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Error getting connections: " + err.Error(),
		})
		return
	}

	var connectionsArray []types.Connection

	for connectionsInDb.Next() {
		var connection types.Connection
		err = connectionsInDb.Scan(&connection.Id, &connection.CommonName, &connection.DatabaseName, &connection.Host, &connection.Port, &connection.Username, &connection.SslMode)
		if err != nil {
			c.JSON(500, gin.H{
				"error": "Error getting connections: " + err.Error(),
			})
			return
		}
		connectionsArray = append(connectionsArray, connection)
	}

	c.JSON(200, gin.H{
		"error":       false,
		"connections": connectionsArray,
	})
}

// HandleGetConnectionById
// @BasePath /
// @Summary Get a connection by ID
// @Description Get a connection by ID
// @Tags connections
// @Produce json
// @Param id path int true "Connection ID"
// @Success 200 {object} object "Returns the connection data"
// @Failure 400 {object} object "Returns that the connection ID is invalid"
// @Failure 404 {object} object "Returns that the connection does not exist"
// @Failure 500 {object} object "Internal error"
// @Router /connections/{id} [get]
func HandleGetConnectionById(c *gin.Context) {
	db := database.InternalDb

	connectionId := c.Param("id")

	id, err := strconv.Atoi(connectionId)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid connection ID",
		})
		return
	}
	exists, err := connections.CheckIfConnectionWithIdExists(id)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Error checking if connection exists: " + err.Error(),
		})
		return
	}
	if !exists {
		c.JSON(404, gin.H{
			"error": "Connection with the given ID does not exist",
		})
		return
	}

	connection, err := db.RunQueryWithParams("SELECT id, common_name, database_name, host, port, username, ssl_mode FROM connections WHERE id = $1", connectionId)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Error getting connection: " + err.Error(),
		})
		return
	}

	var connectionMapped types.Connection

	for connection.Next() {
		err = connection.Scan(&connectionMapped.Id, &connectionMapped.CommonName, &connectionMapped.DatabaseName, &connectionMapped.Host, &connectionMapped.Port, &connectionMapped.Username, &connectionMapped.SslMode)
		if err != nil {
			c.JSON(500, gin.H{
				"error": "Error getting connection: " + err.Error(),
			})
			return
		}
	}

	c.JSON(200, gin.H{
		"error":      false,
		"connection": connectionMapped,
	})
}
