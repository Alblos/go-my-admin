package connections

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine) {
	r.GET("/connections", HandleGetAllConnections)
	r.GET("/connections/:id", HandleGetConnectionById)
	r.POST("/connections/test", HandleTestConnection)
	r.POST("/connections", HandleCreateConnection)
	r.PUT("/connections", HandleUpdateConnection)
	r.DELETE("/connections/:id", HandleDeleteConnection)
}
