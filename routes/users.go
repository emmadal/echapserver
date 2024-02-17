package routes

import (
	"net/http"
	"oblackserver/models"
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
		return;
	}
	err = models.CreateUser(user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"success": false,
		})
		return;
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "User created",
		"success": true,
	})
}