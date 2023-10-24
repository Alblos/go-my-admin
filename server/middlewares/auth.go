package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/go-my-admin/server/routes/types"
	"github.com/go-my-admin/server/utils/auth"
	"net/http"
)

// AuthMiddleware Middleware that checks if the user is authenticated
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, types.ErrorResponse{
				Error:   true,
				Message: "Missing Authorization header",
			})
			c.Abort()
			return
		}
		token = token[7:] // Remove the "Bearer " prefix

		username, authenticated, err := auth.VerifyJWT(token)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.ErrorResponse{
				Error:   true,
				Message: "Internal error",
			})
			c.Abort()
			return
		}

		if !authenticated {
			c.JSON(http.StatusUnauthorized, types.ErrorResponse{
				Error:   true,
				Message: "Invalid token",
			})
			c.Abort()
			return
		}

		c.Set("username", username)
		c.Next()
	}
}
