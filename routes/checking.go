package routes

import (
	"net/http"
	"oblackserver/models"

	"github.com/gin-gonic/gin"
)

func verifyOTP(context *gin.Context) {
	userID := context.GetInt64("userID")

	var jsonOTP models.OTPAuth
	err := context.ShouldBindJSON(&jsonOTP)
	if err != nil {
		context.SecureJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	otp, err := models.GetOTPCodeByUserID(userID)
	if err != nil {
		context.SecureJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	if otp.UserID != userID || otp.IsUsed || otp.Expiration.After(otp.Expiration) {
		context.SecureJSON(http.StatusForbidden, gin.H{
			"success": false,
			"message": "Not authorized to use this code",
		})
		return
	}

	otp.IsUsed = true
	err = models.UpdateOTPCode(*otp)
	if err != nil {
		context.SecureJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	context.SecureJSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Code verified successfully.",
	})

}
