package bcrypt_test

import (
	"github.com/go-my-admin/server/utils/bcrypt"
	"testing"
	"time"
)

func TestBcryptFunctions(t *testing.T) {
	t.Parallel()

	t.Run("Should hash a password", func(t *testing.T) {
		t.Parallel()
		password := "test" + time.Now().String()
		hashedPassword, err := bcrypt.HashPassword(password)
		if err != nil || hashedPassword == "" {
			t.Error("Error hashing password: ", err)
			return
		}
	})

	t.Run("Should generate two different hashes", func(t *testing.T) {
		t.Parallel()
		password := "test" + time.Now().String()
		hashedPasswordOne, err := bcrypt.HashPassword(password)
		if err != nil || hashedPasswordOne == "" {
			t.Error("Error hashing password: ", err)
			return
		}

		hashedPasswordTwo, err := bcrypt.HashPassword(password)
		if err != nil || hashedPasswordTwo == "" {
			t.Error("Error hashing password: ", err)
			return
		}

		if hashedPasswordOne == hashedPasswordTwo {
			t.Error("Hashes should be different")
			return
		}
	})

	t.Run("Should compare a password with a hash", func(t *testing.T) {
		t.Parallel()
		password := "test" + time.Now().String()
		hashedPassword, err := bcrypt.HashPassword(password)
		if err != nil || hashedPassword == "" {
			t.Error("Error hashing password: ", err)
			return
		}

		validated := bcrypt.ValidatePassword(password, hashedPassword)
		if !validated {
			t.Error("Password and hash should match")
			return
		}
	})

	t.Run("Should not compare a password with a hash", func(t *testing.T) {
		t.Parallel()
		password := "test" + time.Now().String()
		hashedPassword, err := bcrypt.HashPassword(password)
		if err != nil || hashedPassword == "" {
			t.Error("Error hashing password: ", err)
			return
		}

		validated := bcrypt.ValidatePassword("wrongPassword", hashedPassword)
		if validated {
			t.Error("Password and hash should not match")
			return
		}
	})
}
