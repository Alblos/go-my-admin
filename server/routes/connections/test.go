package connections

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/go-my-admin/server/database"
	"github.com/go-my-admin/server/database/types"
	"github.com/go-my-admin/server/logger"
)

type TestConnectionRequest struct {
	DbName   string `json:"db_name"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	SslMode  string `json:"ssl_mode"`
}

// HandleTestConnection
// @BasePath /
// @Summary Test a connection
// @Description Test a connection
// @Tags connections
// @Accept json
// @Produce json
// @Param request body TestConnectionRequest true "Request body"
// @Success 200 {object} object "Returns if the connection could be established or not"
// @Failure 400 {object} object "Returns that the request body is invalid or that some required fields are missing"
// @Failure 402 {object} object "If the connection could not be established"
// @Failure 500 {object} object "Internal error"
// @Router /connections/test [post]
func HandleTestConnection(c *gin.Context) {
	var req TestConnectionRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	if req.DbName == "" || req.Host == "" || req.Port == 0 || req.Username == "" || req.Password == "" {
		c.JSON(400, gin.H{
			"error": "Missing required fields",
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
		c.JSON(402, gin.H{
			"error":   true,
			"success": "Connection could not be established",
		})
		return
	}

	c.JSON(200, gin.H{
		"error":   false,
		"success": "Connection established",
	})

}
