package routes

import (
	"net/http"
	"oblackserver/models"

	"github.com/gin-gonic/gin"
)

func getCategories(context *gin.Context) {
	categories, err := models.GetAllCategories()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
			"success": false,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"data":  categories,
		"success": true,
		"message": "All categories",
	})
}