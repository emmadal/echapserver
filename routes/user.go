package routes

import (
	"net/http"
	"oblackserver/helpers"
	"oblackserver/models"
	"strconv"
	"github.com/gin-gonic/gin"
)

func register(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.SecureJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid credentials",
			"success": false,
		})
		return
	}
	err = models.CreateUser(user)
	if err != nil {
		context.SecureJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to create new account",
			"success": false,
		})
		return
	}
	context.SecureJSON(http.StatusOK, gin.H{
		"message": "User created",
		"success": true,
	})
}

func login(context *gin.Context) {
	var auth models.AuthStruct
	err := context.ShouldBindJSON(&auth)
	if err != nil {
		context.SecureJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid credentials",
			"success": false,
		})
		return
	}

	user, err := models.LoginUser(auth)
	if err != nil {
		context.SecureJSON(http.StatusNotFound, gin.H{
			"message": "No user found with provided credentials",
			"success": false,
		})
		return
	}

	// create the jwt token
	token, err := helpers.CreateToken(user.ID, user.Phone)
	if err != nil {
		context.SecureJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to satisfy your request",
			"success": false,
		})
		return
	}

	// get the cookie options for setting up the httpOnly and secure flags
	_, err = context.Cookie("tkauth")
	if err != nil {
		maxAge := 2 * 60 * 60
		domain := helpers.EnvDomainNameKey()
		context.SetCookie("tkauth", token, maxAge, "/", domain, false, true)
	}

	user, _ = models.FindUserByID(user.ID)

	context.SecureJSON(http.StatusOK, gin.H{
		"message": "Login  Successful",
		"data":    user,
		"success": true,
	})
}

func getUserByID(context *gin.Context) {
	userID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	contextUserID := context.GetInt64("userID")

	if err != nil {
		context.SecureJSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse user ID",
			"success": false,
		})
		return
	}

	// Verify if the requesting user is trying to access another users data
	if userID != contextUserID  {
		context.SecureJSON(http.StatusNotFound, gin.H{
			"message": "User Not Found",
			"success": false,
		})
		return
	}

	user, err := models.FindUserByID(userID)
	if err != nil {
		context.SecureJSON(http.StatusNotFound, gin.H{
			"message": "Unable to find user",
			"success": false,
		})
		return
	}
	context.SecureJSON(http.StatusOK, gin.H{
		"message": "user details",
		"data":    user,
		"success": true,
	})
}
