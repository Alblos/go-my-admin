package auth

import (
	"github.com/go-my-admin/server/logger"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

// CreateJWT Creates a JWT token with the given username
func CreateJWT(username string) (string, error) {
	key := os.Getenv("JWT_SECRET")
	if key == "" {
		logger.Error("JWT_SECRET env variable not set", nil)
		return "", nil
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24 * 14).Unix(),
	})

	s, err := t.SignedString([]byte(key))
	if err != nil {
		logger.Error("Error signing JWT", err)
		return "", err
	}

	return s, nil
}
