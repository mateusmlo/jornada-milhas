package tools

import "golang.org/x/crypto/bcrypt"

// HashPassword hashes user password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(bytes), err
}

// ValidatePasswordHash compares provided password with db stored hash
func ValidatePasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
