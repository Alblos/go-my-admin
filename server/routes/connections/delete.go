package connections

import (
	"github.com/gin-gonic/gin"
	"github.com/go-my-admin/server/connections"
	"github.com/go-my-admin/server/database"
	"strconv"
)

func HandleDeleteConnection(c *gin.Context) {
	db := database.InternalDb

	connectionId := c.Param("id")

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

	_, err = db.RunQueryWithParams("DELETE FROM connections WHERE id = $1", connectionId)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Connection deleted successfully",
	})
}
