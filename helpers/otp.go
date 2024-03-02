package helpers

import (
	"math/rand"
	"strconv"
)

const maxNum = 6

func sliceNumToString(num []int) string {
	var result string
	for _, v := range num {
		result += strconv.Itoa(v)
	}
	return result
}

// GenerateOTPCode generate otp code for authentication
func GenerateOTPCode () string {
	sliceNumber := make([]int, 0)
	for  i:=1; i <= maxNum ;i++ {
		random := rand.Intn(10)
		sliceNumber = append(sliceNumber, random)
	}
	result := sliceNumToString(sliceNumber)
	return  result
}