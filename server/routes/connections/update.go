package connections

import (
	"github.com/gin-gonic/gin"
	"github.com/go-my-admin/server/connections"
	"github.com/go-my-admin/server/database"
)

type UpdateConnectionRequest struct {
	Name     string `json:"common_name"`
	DbName   string `json:"database_name"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	SslMode  string `json:"ssl_mode"`
	Id       int    `json:"id"`
}

// HandleUpdateConnection
// @BasePath /
// @Summary Update a connection
// @Description Update a connection
// @Tags connections
// @Accept json
// @Produce json
// @Param request body UpdateConnectionRequest true "Request body"
// @Success 200 {object} object "Returns that the connection was updated successfully"
// @Failure 400 {object} object "Returns that the request body is invalid or that some required fields are missing"
// @Failure 404 {object} object "Returns that the connection with the given ID does not exist"
// @Failure 500 {object} object "Internal error"
func HandleUpdateConnection(c *gin.Context) {
	var req UpdateConnectionRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	if req.Name == "" || req.DbName == "" || req.Host == "" || req.Port == 0 || req.Username == "" || req.Id == 0 {
		c.JSON(400, gin.H{
			"error": "Missing required fields",
		})
		return
	}

	db := database.InternalDb

	exists, err := connections.CheckIfConnectionWithIdExists(req.Id)
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

	_, err = db.RunQueryWithParams("UPDATE connections SET common_name = $2, database_name = $3, host = $4, port = $5, username = $6, ssl_mode = $7 WHERE id = $1", req.Id, req.Name, req.DbName, req.Host, req.Port, req.Username, req.SslMode)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Error updating connection: " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"error":   false,
		"message": "Connection updated successfully",
	})
}
