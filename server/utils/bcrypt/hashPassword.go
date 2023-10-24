package bcrypt

import "golang.org/x/crypto/bcrypt"

// HashPassword hashes a password
func HashPassword(password string) (string, error) {
	p, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(p), err
}
