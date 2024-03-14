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

	if err != nil {
		context.SecureJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(), 
			"success": false,
		})
		return
	}

	err = models.CreateArticle(article)
	if err != nil {
		context.SecureJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(), 
			"success": false,
		})
		return
	}

	context.SecureJSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Article created successfully",
	})

}

