package general

import "github.com/gin-gonic/gin"

func RouterGeneral(r *gin.Engine) {
	r.GET("/ping", HandlePing)
}
