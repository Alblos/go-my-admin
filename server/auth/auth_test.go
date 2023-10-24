package auth_test

import (
	"github.com/go-my-admin/server/auth"
	"os"
	"testing"
	"time"
)

func TestCreateJWT(t *testing.T) {
	err := os.Setenv("JWT_SECRET", "test")
	if err != nil {
		t.Error("Error setting env variable: ", err)
		return
	}

	t.Parallel() // Run this test in parallel with other tests

	t.Run("Should create correctlly a JWT", func(t *testing.T) {
		t.Parallel()
		token, err := auth.CreateJWT("test")
		if err != nil || token == "" {
			t.Error("Error creating JWT: ", err)
			return
		}
	})

	t.Run("Should generate two different JWT", func(t *testing.T) {
		t.Parallel()
		tokenOne, err := auth.CreateJWT("test")
		if err != nil || tokenOne == "" {
			t.Error("Error creating JWT: ", err)
			return
		}

		time.Sleep(time.Second)

		tokenTwo, err := auth.CreateJWT("test")
		if err != nil || tokenTwo == "" {
			t.Error("Error creating JWT: ", err)
			return
		}

		if tokenOne == tokenTwo {
			t.Error("Tokens should be different")
			return
		}
	})

	t.Run("Should return the correct username", func(t *testing.T) {
		t.Parallel()
		// Generate a random username
		username := "test" + time.Now().String()
		token, err := auth.CreateJWT(username)
		if err != nil || token == "" {
			t.Error("Error creating JWT: ", err)
			return
		}

		// Verify the token
		usernameFromToken, authenticated, err := auth.VerifyJWT(token)
		if err != nil || usernameFromToken == "" {
			t.Error("Error verifying JWT: ", err)
			return
		}

		if usernameFromToken != username {
			t.Error("Username should be the same")
			return
		}

		if !authenticated {
			t.Error("JWT should be authenticated")
			return
		}
	})
}
