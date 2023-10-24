package connections

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/go-my-admin/server/database"
	"github.com/go-my-admin/server/routes/types"
)

type HandleCreateConnectionRequest struct {
	Name     string `json:"common_name"`
	DbName   string `json:"db_name"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	SslMode  string `json:"ssl_mode"`
}

type HandleCreateConnectionResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Id      int    `json:"id"`
}

// HandleCreateConnection
// @BasePath /
// @Summary Create a connection
// @Description Creates a connection with the given parameters
// @Tags connections
// @Accept json
// @Produce json
// @Param request body HandleCreateConnectionRequest true "Request body"
// @Success 200 {object} HandleCreateConnectionResponse "The connection was created successfully"
// @Failure 400 {object} types.ErrorResponse "The request body is invalid or that some required fields are missing"
// @Failure 502 {object} types.ErrorResponse "Error creating connection"
// @Failure 500 {object} types.ErrorResponse "Internal error"
// @Router /connections/create [post]
func HandleCreateConnection(c *gin.Context) {
	var req HandleCreateConnectionRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, types.ErrorResponse{
			Error:   true,
			Message: "Invalid request body",
		})
		return
	}

	if req.Name == "" || req.DbName == "" || req.Host == "" || req.Port == 0 || req.Username == "" || req.Password == "" {
		c.JSON(400, types.ErrorResponse{
			Error:   true,
			Message: "Missing required fields",
		})
		return
	}

	db := database.InternalDb

	var rows *sql.Rows

	if req.SslMode != "" {
		rows, err = db.RunQueryWithParams("INSERT INTO connections (common_name, database_name, host, port, username, password, ssl_mode) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id",
			req.Name, req.DbName, req.Host, req.Port, req.Username, req.Password, req.SslMode)
	} else {
		rows, err = db.RunQueryWithParams("INSERT INTO connections (common_name, database_name, host, port, username, password) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
			req.Name, req.DbName, req.Host, req.Port, req.Username, req.Password)
	}

	if err != nil {
		c.JSON(502, types.ErrorResponse{
			Error:   true,
			Message: "Error creating connection: " + err.Error(),
		})
		return
	}

	var id int
	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			c.JSON(500, types.ErrorResponse{
				Error:   true,
				Message: "Error scanning connection: " + err.Error(),
			})
			return
		}
	}

	c.JSON(200, HandleCreateConnectionResponse{
		Error:   false,
		Message: "Connection created successfully",
		Id:      id,
	})
}
