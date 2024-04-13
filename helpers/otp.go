package helpers

import (
	"math/rand"
	"strconv"
	"strings"
)

const maxNum = 5

// GenerateOTPCode generate otp code for authentication
func GenerateOTPCode() string {
	var randomString []string
	for i := 1; i <= maxNum; i++ {
		random := rand.Intn(10)
		randomString = append(randomString, strconv.Itoa(random))
	}
	result := strings.Join(randomString, "")
	return result
}
