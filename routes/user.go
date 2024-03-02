package routes

import (
	"fmt"
	"net/http"
	"oblackserver/helpers"
	"oblackserver/models"
	"time"

	"github.com/gin-gonic/gin"
)

func register(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"success": false,
		})
		return
	}
	err = models.CreateUser(user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"success": false,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "User created",
		"success": true,
	})
}

func login(context *gin.Context) {
	var auth models.AuthStruct
	err := context.ShouldBindJSON(&auth)
	if err != nil {
		context.SecureJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"success": false,
		})
		return
	}

	user, err := models.LoginUser(auth)
	if err != nil {
		context.SecureJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"success": false,
		})
		return
	}

	// create the jwt token
	token, err := helpers.CreateToken(user.ID, user.Phone)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not generate user token",
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
	fmt.Println(otpObject)
	err = models.SaveOTPCode(otpObject) // save the otp to database
	if err != nil {
		context.SecureJSON(http.StatusInternalServerError, gin.H{
			"message": "Could not generate otp code",
			"success": false,
		})
		return
	}

	// get the cookie options for setting up the httpOnly and secure flags
	_, err = context.Cookie("tkauth")
	if err != nil {
		maxAge := 2 * 60 * 60
		context.SetCookie("tkauth", token, maxAge, "/", "localhost", false, true)
	}

	context.SecureJSON(http.StatusOK, gin.H{
		"message": "Login  Successful",
		"otp":     otp,
		"success": true,
	})
}
