package routes

import (
	"echapserver/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getCounties(context *gin.Context) {
	countries, err := models.FindCountries()

	if err != nil {
		context.SecureJSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Countries not found",
		})
		return
	}

	context.SecureJSON(http.StatusOK, gin.H{
		"success": true,
		"data":    countries,
	})
}
