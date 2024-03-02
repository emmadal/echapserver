package helpers

import "golang.org/x/crypto/bcrypt"

// HashPassword hashes a password using bcrypt algorithm with default cost of 10
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// DecryptPassword decrypt a password using bcrypt algorithm
func DecryptPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}