package connections

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/go-my-admin/server/database"
	"github.com/go-my-admin/server/database/types"
	"github.com/go-my-admin/server/logger"
	routeTypes "github.com/go-my-admin/server/routes/types"
)

type HandleTestConnectionRequest struct {
	DbName   string `json:"db_name"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	SslMode  string `json:"ssl_mode"`
}

type HandleTestConnectionResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

// HandleTestConnection
// @BasePath /
// @Summary Test a connection
// @Description Test a connection
// @Tags connections
// @Accept json
// @Produce json
// @Param request body HandleTestConnectionRequest true "Request body"
// @Success 200 {object} HandleTestConnectionResponse "The connection was tested successfully"
// @Failure 400 {object} types.ErrorResponse "The request body is invalid or that some required fields are missing"
// @Failure 402 {object} types.ErrorResponse "Connection could not be established"
// @Failure 500 {object} types.ErrorResponse "Internal error"
// @Router /connections/test [post]
func HandleTestConnection(c *gin.Context) {
	var req HandleTestConnectionRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, routeTypes.ErrorResponse{
			Error:   true,
			Message: "Invalid request body",
		})
		return
	}

	if req.DbName == "" || req.Host == "" || req.Port == 0 || req.Username == "" || req.Password == "" {
		c.JSON(400, routeTypes.ErrorResponse{
			Error:   true,
			Message: "Missing required fields",
		})
		return
	}
	if req.SslMode == "" {
		req.SslMode = "false"
	}

	connData := types.Connection{
		DatabaseName: req.DbName,
		Host:         req.Host,
		Port:         req.Port,
		Username:     req.Username,
		Password:     req.Password,
		SslMode:      req.SslMode,
	}

	str := database.GetConnectionString(connData)

	var connection database.DBConnection
	err = connection.Connect(str)
	defer func(Cnx *sql.DB) {
		err := Cnx.Close()
		if err != nil {
			logger.Error("Error closing connection", err)
		}
	}(connection.Cnx)

	if err != nil {
		c.JSON(402, routeTypes.ErrorResponse{
			Error:   true,
			Message: "Connection could not be established",
		})
		return
	}

	c.JSON(200, HandleTestConnectionResponse{
		Error:   false,
		Message: "Connection tested successfully",
	})

}
