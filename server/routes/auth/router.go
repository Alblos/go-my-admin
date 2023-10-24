package auth

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine) {
	r.POST("/auth/signup", HandleSignup)
	r.POST("/auth/login", HandleLogin)
}
