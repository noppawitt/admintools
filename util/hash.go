package util

import "golang.org/x/crypto/bcrypt"

// HashPassword returns hashed password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
