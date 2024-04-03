package routes

import (
	"fmt"
	"net/http"
	"oblackserver/models"
	"strconv"

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

func deleteArticle(context *gin.Context) {
	// get the param ID
	articleID, err := strconv.ParseInt(context.Param("id"), 10, 64)

	// get the userID in the context
	userID := context.GetInt64("userID")

	if err != nil {
		context.SecureJSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse article ID",
			"success": false,
		})
		return
	}

	fmt.Println("ArticleID:", articleID)

	// verify is the article exist in the database
	article, err := models.FindArticleByID(articleID)
	if err != nil {
		context.SecureJSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
			"success": false,
		})
		return
	}

	// We verify if it's the article owner before to update
	if userID != article.AuthorID {
		context.SecureJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete it", "success": false})
		return
	}

	// delete article
	err = models.DeleteArticle(articleID)
	if err != nil {
		context.SecureJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to delete article",
			"success": false,
		})
		return
	}

	context.SecureJSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Article deleted",
	})
}
