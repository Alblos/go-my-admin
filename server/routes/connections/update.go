package connections

import (
	"github.com/gin-gonic/gin"
	"github.com/go-my-admin/server/connections"
	"github.com/go-my-admin/server/database"
	"github.com/go-my-admin/server/routes/types"
)

type HandleUpdateConnectionRequest struct {
	Name     string `json:"common_name"`
	DbName   string `json:"database_name"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	SslMode  string `json:"ssl_mode"`
	Id       int    `json:"id"`
}

type HandleUpdateConnectionResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

// HandleUpdateConnection
// @BasePath /
// @Summary Update a connection
// @Description Update a connection
// @Tags connections
// @Accept json
// @Produce json
// @Param request body HandleUpdateConnectionRequest true "Request body"
// @Success 200 {object} HandleUpdateConnectionResponse "The connection was updated successfully"
// @Failure 400 {object} types.ErrorResponse "The request body is invalid or that some required fields are missing"
// @Failure 404 {object} types.ErrorResponse "Connection with the given ID does not exist"
// @Failure 500 {object} types.ErrorResponse "Internal error"
// @Router /connections/update [put]
func HandleUpdateConnection(c *gin.Context) {
	var req HandleUpdateConnectionRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, types.ErrorResponse{
			Error:   true,
			Message: "Invalid request body",
		})
		return
	}

	if req.Name == "" || req.DbName == "" || req.Host == "" || req.Port == 0 || req.Username == "" || req.Id == 0 {
		c.JSON(400, types.ErrorResponse{
			Error:   true,
			Message: "Missing required fields",
		})
		return
	}

	db := database.InternalDb

	exists, err := connections.CheckIfConnectionWithIdExists(req.Id)
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

	_, err = db.RunQueryWithParams("UPDATE connections SET common_name = $2, database_name = $3, host = $4, port = $5, username = $6, ssl_mode = $7 WHERE id = $1", req.Id, req.Name, req.DbName, req.Host, req.Port, req.Username, req.SslMode)
	if err != nil {
		c.JSON(500, types.ErrorResponse{
			Error:   true,
			Message: "Error updating connection: " + err.Error(),
		})
		return
	}

	c.JSON(200, HandleUpdateConnectionResponse{
		Error:   false,
		Message: "Connection updated successfully",
	})
}
