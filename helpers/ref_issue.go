package helpers

import (
	"fmt"
	"math/rand"
)

// GenerateRandomString generates a random string of specified length
func GenerateRandomString(length int) string {
	// Define the character set from which to generate the random string
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"


	// Create a byte slice of the specified length
	randomString := make([]byte, length)

	// Fill the byte slice with random characters from the charset
	for i := range randomString {
		randomString[i] = charset[rand.Intn(len(charset))]
	}
	value := fmt.Sprintf("ISSUE-%s", string(randomString))

	return value
}