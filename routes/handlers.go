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
			"message": err.Error(),
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

func createCategory(context *gin.Context) {
	var category models.Category
	err := context.ShouldBindJSON(&category)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad request", "success": false})
	}
	err = models.CreateCategory(category)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error", 
			"success": false,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Article created",
	})

}