package bcrypt

import "golang.org/x/crypto/bcrypt"

// ValidatePassword validates a password against a hash
func ValidatePassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
