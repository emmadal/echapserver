package helpers

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)

// CreateToken creates a new JWT token with the given claims and signs it using HS256 algorithm.
func CreateToken(userID int64, phone string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"phone": phone,
		"expiry":time.Now().Add(time.Hour * 2).Unix(),
	})
	key := EnvSecretKey()
	return token.SignedString([]byte(key))
}