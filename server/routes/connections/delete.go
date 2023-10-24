package connections

import (
	"github.com/gin-gonic/gin"
	"github.com/go-my-admin/server/connections"
	"github.com/go-my-admin/server/database"
	"github.com/go-my-admin/server/routes/types"
	"strconv"
)

type HandleDeleteConnectionResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

// HandleDeleteConnection
// @BasePath /
// @Summary Delete a connection
// @Description Delete a connection with the given ID
// @Tags connections
// @Accept json
// @Produce json
// @Param id path int true "Connection ID"
// @Success 200 {object} HandleDeleteConnectionResponse "The connection was deleted successfully"
// @Failure 400 {object} types.ErrorResponse "Invalid connection ID"
// @Failure 404 {object} types.ErrorResponse "Connection with the given ID does not exist"
// @Failure 500 {object} types.ErrorResponse "Internal error"
// @Router /connections/delete/{id} [delete]
func HandleDeleteConnection(c *gin.Context) {
	db := database.InternalDb

	connectionId := c.Param("id")

	id, err := strconv.Atoi(connectionId)
	if err != nil {
		c.JSON(400, types.ErrorResponse{
			Error:   true,
			Message: "Invalid connection ID",
		})
		return
	}
	exists, err := connections.CheckIfConnectionWithIdExists(id)
	if err != nil {
		c.JSON(500, types.ErrorResponse{
			Error:   true,
			Message: "Error checking if connection exists: " + err.Error(),
		})
		return
	}
	if !exists {
		c.JSON(404, types.ErrorResponse{
			Error:   true,
			Message: "Connection with the given ID does not exist",
		})
		return
	}

	_, err = db.RunQueryWithParams("DELETE FROM connections WHERE id = $1", connectionId)
	if err != nil {
		c.JSON(500, types.ErrorResponse{
			Error:   true,
			Message: "Error deleting connection: " + err.Error(),
		})
		return
	}

	c.JSON(200, HandleDeleteConnectionResponse{
		Error:   false,
		Message: "Connection deleted successfully",
	})
}
