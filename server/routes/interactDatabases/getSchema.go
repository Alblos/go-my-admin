package interactDatabases

import (
	"github.com/gin-gonic/gin"
	"github.com/go-my-admin/server/connections"
	"strconv"
)

func HandleGetDbSchema(c *gin.Context) {
	connectionId := c.Param("connectionId")

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

	tables, err := connections.GetTablesInDatabase(id)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"tables": tables,
	})

}
