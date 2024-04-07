package middlewares

import (
	"echapserver/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Authenticate is a middleware for authorization
func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
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
