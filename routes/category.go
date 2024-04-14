package routes

import (
	"echapserver/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getCategories(context *gin.Context) {
	categories, err := models.GetAllCategories()
	if err != nil {
		context.SecureJSON(http.StatusNotFound, gin.H{
			"message": "Data not found",
			"success": false,
		})
		return
	}
	context.SecureJSON(http.StatusOK, gin.H{
		"data":    categories,
		"success": true,
	})
}

func createCategory(context *gin.Context) {
	userID := context.GetInt64("userID")
	var category models.Category
	category.UserID = userID

	if err := context.ShouldBindJSON(&category); err != nil {
		context.SecureJSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
			"success": false,
		})
		return
	}

	if user, _ := models.FindUserByID(userID); !user.Role {
		context.SecureJSON(http.StatusForbidden, gin.H{
			"message": "You're not authorized to create a category",
			"success": false,
		})
		return
	}

	if err := models.CreateCategory(category); err != nil {
		context.SecureJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to create data",
			"success": false,
		})
		return
	}

	context.SecureJSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Category created successfully",
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
		"data":    category,
	})
}

func updateCategory(context *gin.Context) {
	// convert id parameter into int64
	id := context.Param("id")

	// get the userID in the context
	userID := context.GetInt64("userID")

	categoryID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		context.SecureJSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse category ID",
			"success": false,
		})
		return
	}

	// verify if the category exist in the database
	category, err := models.GetCategoryByID(categoryID)
	if err != nil {
		context.SecureJSON(http.StatusNotFound, gin.H{
			"message": "Unable to find this data",
			"success": false,
		})
		return
	}

	// We verify if it's the category owner before to update
	if userID != category.UserID {
		context.SecureJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update it", "success": false})
		return
	}

	// send the new  data to be updated on the database
	var newData models.Category
	newData.UserID = userID
	err = context.ShouldBindJSON(&newData)
	if err != nil {
		context.SecureJSON(http.StatusBadRequest, gin.H{"message": "Bad request", "success": false})
		return
	}

	// Update the Category with the new Data
	newData.ID = categoryID
	err = models.UpdateCategory(newData)
	if err != nil {
		context.SecureJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to update category data",
			"success": false,
		})
		return
	}
	context.SecureJSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Category updated",
	})
}

func deleteCategory(context *gin.Context) {
	// convert id parameter into int64
	id := context.Param("id")

	// get the userID in the context
	userID := context.GetInt64("userID")

	categoryID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		context.SecureJSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse category ID",
			"success": false,
		})
		return
	}

	// verify is the category exist in the database
	category, err := models.GetCategoryByID(categoryID)
	if err != nil {
		context.SecureJSON(http.StatusNotFound, gin.H{
			"message": "Unable to find this data",
			"success": false,
		})
		return
	}

	// We verify if it's the category owner before to update
	if userID != category.UserID {
		context.SecureJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete it", "success": false})
		return
	}

	// delete the category
	err = models.DeleteCategory(categoryID)
	if err != nil {
		context.SecureJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to delete category",
			"success": false,
		})
		return
	}

	context.SecureJSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Category deleted",
	})
}
