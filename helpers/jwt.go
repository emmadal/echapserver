package helpers

import (
	"errors"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

// secret key
var key = EnvSecretKey()

// CreateToken creates a new JWT token with the given claims and signs it using HS256 algorithm.
func CreateToken(userID int64, phone string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"phone":  phone,
		"expiry": time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(key))
}

// VerifyToken verify the given token to get its payload.
func VerifyToken(tokenString string) (int64, error) {
	// verify the signature of the token
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Invalid signature")
		}
		return []byte(key), nil
	})

	if err != nil {
		return 0, errors.New("Couldn't handle this token")
	}

	// check the validity of token
	if tokenIsValid := parsedToken.Valid; !tokenIsValid {
		return 0, errors.New("Invalid token")
	}

	// extract claims data in the token
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("Invalid token claims")
	}
	userID := int64(claims["userID"].(float64)) // .(float64) means type checking

	return userID, nil
}
