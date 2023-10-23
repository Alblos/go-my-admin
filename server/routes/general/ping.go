package general

import "github.com/gin-gonic/gin"

// HandlePing
// @BasePath /
// @Summary Ping the server
// @Description Ping the server
// @Tags general
// @Produce json
// @Success 200 {object} object "Returns pong"
// @Router /ping [get]
func HandlePing(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}
