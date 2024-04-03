package routes

import (
	"net/http"
	"oblackserver/models"
	"github.com/gin-gonic/gin"
)

func createArticle(context *gin.Context) {
	userID := context.GetInt64("userID")
	var article models.Article

	err := context.ShouldBindJSON(&article)
	if err != nil {
		context.SecureJSON(http.StatusBadRequest, gin.H{"message": "Bad request", "success": false})
		return
	}

	if article.AuthorID != userID {
		context.SecureJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to continue", "success": false})
		return
	}

	user, _ := models.FindUserByID(userID)
	if !user.Premium {
		context.SecureJSON(http.StatusUnauthorized, gin.H{
			"message": "You need to be a premium user before to create an article",
			"success": false,
		})
		return
	}

	err = models.CreateArticle(article)
	if err != nil {
		context.SecureJSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
			"success": false,
		})
		return
	}

	context.SecureJSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Article created successfully",
	})
}

func getArticles(context *gin.Context) {
	categoryID := context.Param("category_id")

	articles, err := models.GetAllArticle(categoryID)

	if err != nil {
		context.SecureJSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
			"success": false,
		})
		return
	}

	context.SecureJSON(http.StatusOK, gin.H{
		"success": true,
		"data":    articles,
	})
}
