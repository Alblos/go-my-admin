package general

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine) {
	r.GET("/ping", HandlePing)
}
