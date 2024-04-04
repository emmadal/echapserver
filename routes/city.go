package routes

import (
	"net/http"
	"oblackserver/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getCitiesByCountry(context *gin.Context) {
	countryID, err := strconv.ParseInt(context.Param("countryID"), 10, 64)
	if err != nil {
		context.SecureJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Could not parse country ID",
		})
		return
	}
	cities, err := models.FindCitiesByCountryID(countryID)
	if err != nil {
		context.SecureJSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Cities not found",
		})
		return
	}
	context.SecureJSON(http.StatusOK, gin.H{
		"success": true,
		"data":    cities,
	})
}

