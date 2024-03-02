package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"oblackserver/helpers"
	"oblackserver/models"
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
		context.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"success": false,
		})
		return
	}

	user, err := models.LoginUser(auth)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"success": false,
		})
		return
	}

	token, err := helpers.CreateToken(user.ID, user.Phone)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"success": false,
		})
		return
	}

	otp := helpers.GenerateOTPCode()

	_, err = context.Cookie("tkauth")

	if err != nil {
		maxAge := 2 * 60 * 60
		context.SetCookie("tkauth", token, maxAge, "/", "localhost", false, true)
	}

	context.SecureJSON(http.StatusOK, gin.H{
		"message": "Login  Successful",
		"otp":     otp,
		"token":   token,
		"success": true,
	})
}
