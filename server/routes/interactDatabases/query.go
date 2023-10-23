package interactDatabases

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/go-my-admin/server/connections"
	"github.com/go-my-admin/server/utils/routes"
	"time"
)

type requestQuery struct {
	Query string `json:"query"`
}

// HandleQuery
// @BasePath /
// @Summary Query a database
// @Description Query a database
// @Tags interactDatabases
// @Accept json
// @Produce json
// @Param connectionId path int true "Connection ID"
// @Param request body requestQuery true "Request body"
// @Success 200 {object} object "Returns the rows"
// @Failure 400 {object} object "If the request body is invalid"
// @Failure 501 {object} object "If the connection could not be established"
// @Failure 502 {object} object "If the query could not be run"
// @Failure 500 {object} object "Internal error"
// @Router /query/{connectionId} [post]
func HandleQuery(c *gin.Context) {
	done, id := routes.CheckIfConnectionIdExists(c)
	if done {
		return
	}

	var request requestQuery
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	db, err := connections.GetConnectionOrCreateIfNotExists(id)
	if err != nil {
		c.JSON(501, gin.H{
			"error": "Error getting connection: " + err.Error(),
		})
		return
	}

	startTime := time.Now()
	rows, err := db.RunRawQuery(request.Query)
	if err != nil {
		c.JSON(502, gin.H{
			"error": "Error running query: " + err.Error(),
		})
		return
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			c.JSON(500, gin.H{
				"error": "Error closing rows: " + err.Error(),
			})
			return
		}
	}(rows)
	endTime := time.Now()

	var columns []string
	columns, err = rows.Columns()
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Error getting columns: " + err.Error(),
		})
		return
	}

	var rowsData [][]interface{}
	for rows.Next() {
		values := make([]interface{}, len(columns))
		pointers := make([]interface{}, len(columns))
		for i := range values {
			pointers[i] = &values[i]
		}
		err = rows.Scan(pointers...)
		if err != nil {
			c.JSON(500, gin.H{
				"error": "Error scanning rows: " + err.Error(),
			})
			return
		}
		rowsData = append(rowsData, values)
	}

	// If no error, return the rows
	c.JSON(200, gin.H{
		"error":   false,
		"columns": columns,
		"rows":    rowsData,
		"time":    endTime.Sub(startTime).Milliseconds(),
	})
}
