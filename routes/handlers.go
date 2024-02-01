package routes

import (
	"net/http"
	"oblackserver/models"
	"strconv"
	"github.com/gin-gonic/gin"
)

func getCategories(context *gin.Context) {
	categories, err := models.GetAllCategories()
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
			"success": false,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"data":  categories,
		"success": true,
	})
}

func createCategory(context *gin.Context) {
	var category models.Category
	err := context.ShouldBindJSON(&category)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad request", "success": false})
		return
	}
	err = models.CreateCategory(category)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to create category data", 
			"success": false,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Article created",
	})

}

func getCategoryByID(context *gin.Context) {
	id := context.Param("id")
	categoryID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse category ID",
			"success": false,
		})
		return
	}
	category, err := models.GetCategoryByID(categoryID)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Unable to fetch category data", 
			"success": false,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": category,
	})
}