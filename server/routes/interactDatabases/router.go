package interactDatabases

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine) {
	r.GET("/schema/:connectionId", HandleGetDbSchema)
}
