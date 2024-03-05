package middlewares

import (
	"net/http"
	"oblackserver/helpers"
	"github.com/gin-gonic/gin"
)

// Authenticate is a middleware for authorization
func Authenticate(context *gin.Context) {
	token, err := context.Cookie("tkauth")

	if err != nil || token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	//  validate the token and get user data from it
	userID, err := helpers.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	// Attach the userID to the request
	context.Set("userID", userID)

	context.Next()
}
