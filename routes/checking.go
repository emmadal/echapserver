package routes

import (
	"echapserver/helpers"
	"echapserver/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func verifyOTP(context *gin.Context) {
	userID := context.GetInt64("userID")

	var jsonOTP models.OTPAuth
	err := context.ShouldBindJSON(&jsonOTP)
	if err != nil {
		context.SecureJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request",
		})
		return
	}

	otp, err := models.GetOTPCode(jsonOTP.Code)
	if err != nil {
		context.SecureJSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Invalid code",
		})
		return
	}

	if otp.UserID != userID || otp.IsUsed || otp.Expiration.After(otp.Expiration) {
		context.SecureJSON(http.StatusForbidden, gin.H{
			"success": false,
			"message": "Code not authorized",
		})
		return
	}

	otp.IsUsed = true
	err = models.UpdateOTPCode(*otp)
	if err != nil {
		context.SecureJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Something went wrong. Try later",
		})
		return
	}

	user, err := models.FindUserByID(otp.UserID)
	if err != nil {
		context.SecureJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "No user found",
		})
		return
	}
	context.SecureJSON(http.StatusOK, gin.H{
		"success": true,
		"data":    user,
		"message": "Code verified successfully.",
	})
}

func getOTP(context *gin.Context) {
	userID := context.GetInt64("userID")
	user, err := models.FindUserByID(userID)

	if err != nil {
		context.SecureJSON(http.StatusInternalServerError, gin.H{
			"message": "Could not have OTP",
			"success": false,
		})
		return
	}

	if userID != user.ID {
		context.SecureJSON(http.StatusNotFound, gin.H{
			"message": "Unable to generate OTP for this user",
			"success": false,
		})
		return
	}

	// generate the otp code
	otp := helpers.GenerateOTPCode()
	var otpObject = models.OTP{
		Code:       otp,
		Expiration: time.Now().Add(time.Minute * 2), // expires in 2 minutes from now
		UserID:     user.ID,
	}

	// save the otp to database
	err = models.SaveOTPCode(otpObject)
	if err != nil {
		context.SecureJSON(http.StatusInternalServerError, gin.H{
			"message": "Could not generate otp code",
			"success": false,
		})
		return
	}
	context.SecureJSON(http.StatusOK, gin.H{
		"message": "Code generated",
		"data":    otp,
		"success": true,
	})
}
