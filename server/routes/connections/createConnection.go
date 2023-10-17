package connections

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/go-my-admin/server/database"
)

type CreateConnectionRequest struct {
	Name     string `json:"common_name"`
	DbName   string `json:"db_name"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	SslMode  string `json:"ssl_mode"`
}

func HandleCreateConnection(c *gin.Context) {
	var req CreateConnectionRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	if req.Name == "" || req.DbName == "" || req.Host == "" || req.Port == 0 || req.Username == "" || req.Password == "" {
		c.JSON(400, gin.H{
			"error": "Missing required fields",
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
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	var id int
	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	c.JSON(200, gin.H{
		"message": "Connection created successfully",
		"id":      id,
	})
}
