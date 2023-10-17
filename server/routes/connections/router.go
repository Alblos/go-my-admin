package connections

import "github.com/gin-gonic/gin"

func RouterConnections(r *gin.Engine) {
	r.GET("/connections", HandleGetAllConnections)
	r.GET("/connections/:id", HandleGetConnectionById)
	r.POST("/connections", HandleCreateConnection)
	r.PUT("/connections", HandleUpdateConnection)
	r.DELETE("/connections/:id", HandleDeleteConnection)
}
