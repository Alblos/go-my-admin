package connections

import (
	"github.com/gin-gonic/gin"
	"github.com/go-my-admin/server/database"
	"github.com/go-my-admin/server/database/internalDbTypes"
)

func HandleGetAllConnections(c *gin.Context) {
	db := database.InternalDb

	connections, err := db.RunRawQuery("SELECT id, common_name, database_name, host, port, username, password, ssl_mode FROM connections")
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	var connectionsArray []internalDbTypes.Connection

	for connections.Next() {
		var connection internalDbTypes.Connection
		err = connections.Scan(&connection.Id, &connection.CommonName, &connection.DatabaseName, &connection.Host, &connection.Port, &connection.Username, &connection.Password, &connection.SslMode)
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
	connection, err := db.RunQueryWithParams("SELECT * FROM connections WHERE id = $1", connectionId)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"connection": connection,
	})
}
