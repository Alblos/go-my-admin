package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-my-admin/server/connections"
	"strconv"
)

// CheckIfConnectionIdExists checks if a connection with the given ID exists
func CheckIfConnectionIdExists(c *gin.Context) (done bool, id int) {
	connectionId := c.Param("connectionId")

	id, err := strconv.Atoi(connectionId)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Invalid connection ID",
		})
		return true, 0
	}
	exists, err := connections.CheckIfConnectionWithIdExists(id)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return true, 0
	}
	if !exists {
		c.JSON(404, gin.H{
			"error": "Connection with the given ID does not exist",
		})
		return true, 0
	}
	return false, id
}