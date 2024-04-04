package routes

import (
	"net/http"
	"oblackserver/helpers"
	"oblackserver/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

var domain = helpers.EnvDomainNameKey()
var mode = helpers.EnvMode()

func register(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.SecureJSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
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

	context.SecureJSON(http.StatusOK, gin.H{
		"message": "Login  Successful",
		"data":    token,
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
	if userID != contextUserID {
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

func updateUser(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	userID := context.GetInt64("userID")

	if err != nil {
		context.SecureJSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse user ID",
			"success": false,
		})
		return
	}

	// verify if the user exist in the database
	user, err := models.FindUserByID(id)
	if err != nil {
		context.SecureJSON(http.StatusNotFound, gin.H{
			"message": "No data found",
			"success": false,
		})
		return
	}

	// We verify if it's the user owner before to update
	if userID != user.ID {
		context.SecureJSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized to update",
			"success": false,
		})
		return
	}

	// send the new  data to be updated on the database
	var newData models.User
	err = context.ShouldBindJSON(&newData)
	if err != nil {
		context.SecureJSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
			"success": false,
		})
		return
	}

	// Update the user with the new Data
	newData.ID = user.ID
	newData.CreatedAt = user.CreatedAt
	err = models.UpdateUser(newData)
	if err != nil {
		context.SecureJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to update user profile",
			"success": false,
		})
		return
	}
	context.SecureJSON(http.StatusOK, gin.H{
		"success": true,
		"data":    newData,
		"message": "User profile updated",
	})
}
