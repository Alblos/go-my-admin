package interactDatabases

import (
	"github.com/gin-gonic/gin"
	"github.com/go-my-admin/server/connections"
	"github.com/go-my-admin/server/utils/routes"
)

// HandleGetDbSchema
// @BasePath /
// @Summary Get the schema of a database
// @Description Get the schema of a database
// @Tags interactDatabases
// @Accept json
// @Produce json
// @Param connectionId path int true "Connection ID"
// @Success 200 {object} object "Returns the schema"
// @Failure 400 {object} object "Returns the error"
// @Failure 501 {object} object "If the connection could not be established or the schema could not be retrieved"
// @Router /schema/{connectionId} [get]
func HandleGetDbSchema(c *gin.Context) {
	done, id := routes.CheckIfConnectionIdExists(c)
	if done {
		return
	}

	schema, err := connections.GetFullDbSchema(id, true)
	if err != nil {
		c.JSON(501, gin.H{
			"error": "Error getting schema: " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"error":  false,
		"schema": schema,
	})

}
