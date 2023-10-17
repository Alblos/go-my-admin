package connections

import (
	"github.com/gin-gonic/gin"
	"github.com/go-my-admin/server/connections"
	"github.com/go-my-admin/server/database"
	"github.com/go-my-admin/server/database/internalDbTypes"
	"strconv"
)

func HandleGetAllConnections(c *gin.Context) {
	db := database.InternalDb

	connectionsInDb, err := db.RunRawQuery("SELECT id, common_name, database_name, host, port, username, ssl_mode FROM connections")
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	var connectionsArray []internalDbTypes.Connection

	for connectionsInDb.Next() {
		var connection internalDbTypes.Connection
		err = connectionsInDb.Scan(&connection.Id, &connection.CommonName, &connection.DatabaseName, &connection.Host, &connection.Port, &connection.Username, &connection.SslMode)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		connectionsArray = append(connectionsArray, connection)
	}

	c.JSON(200, gin.H{
		"connections": connectionsArray,
	})
}

func HandleGetConnectionById(c *gin.Context) {
	db := database.InternalDb

	connectionId := c.Param("id")
	connection, err := db.RunQueryWithParams("SELECT id, common_name, database_name, host, port, username, ssl_mode FROM connections WHERE id = $1", connectionId)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	var connectionMapped internalDbTypes.Connection

	id, err := strconv.Atoi(connectionId)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Invalid connection ID",
		})
		return
	}
	exists, err := connections.CheckIfConnectionWithIdExists(id)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	if !exists {
		c.JSON(404, gin.H{
			"error": "Connection with the given ID does not exist",
		})
		return
	}

	for connection.Next() {
		err = connection.Scan(&connectionMapped.Id, &connectionMapped.CommonName, &connectionMapped.DatabaseName, &connectionMapped.Host, &connectionMapped.Port, &connectionMapped.Username, &connectionMapped.SslMode)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	c.JSON(200, connectionMapped)
}
