package interactDatabases

import (
	"github.com/gin-gonic/gin"
	"github.com/go-my-admin/server/connections"
	"github.com/go-my-admin/server/utils/routes"
)

// HandleGetDbSchema gets the schema of a database
func HandleGetDbSchema(c *gin.Context) {
	done, id := routes.CheckIfConnectionIdExists(c)
	if done {
		return
	}

	schema, err := connections.GetFullDbSchema(id, true)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"error":  false,
		"schema": schema,
	})

}
