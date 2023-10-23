package connections

import (
	"github.com/gin-gonic/gin"
	"github.com/go-my-admin/server/connections"
	"github.com/go-my-admin/server/database"
	"strconv"
)

// HandleDeleteConnection
// @BasePath /
// @Summary Delete a connection
// @Description Delete a connection
// @Tags connections
// @Accept json
// @Produce json
// @Param id path int true "Connection ID"
// @Success 200 {object} object "Returns the success message"
// @Failure 400 {object} object "Returns that the connection ID is invalid"
// @Failure 404 {object} object "Returns that the connection does not exist"
// @Failure 500 {object} object "Internal error"
// @Router /connections/delete/{id} [delete]
func HandleDeleteConnection(c *gin.Context) {
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

	_, err = db.RunQueryWithParams("DELETE FROM connections WHERE id = $1", connectionId)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Error deleting connection: " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"error":   false,
		"message": "Connection deleted successfully",
	})
}
