package auth

import (
	"errors"
	"github.com/go-my-admin/server/logger"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

// VerifyJWT verifies the authenticity of a JWT token
func VerifyJWT(tokenString string) (username string, authenticated bool, err error) {
	key := os.Getenv("JWT_SECRET")
	if key == "" {
		logger.Error("JWT_SECRET env variable not set", nil)
		return "", false, errors.New("JWT_SECRET env variable not set")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(key), nil
	})

	if err != nil {
		logger.Error("Error parsing JWT", err)
		return "", false, nil
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", false, nil
	}

	username, ok = claims["username"].(string)
	if !ok {
		return "", false, nil
	}

	return username, true, nil
}
