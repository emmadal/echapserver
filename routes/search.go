package routes

import (
	"echapserver/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func searchArticle(context *gin.Context) {
	categoryID, err := strconv.ParseInt(context.Query("category"), 10, 64)
	title := context.Query("title")

	if err != nil {
		context.SecureJSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
			"success": false,
		})
		return
	}

	articles, err := models.SearchArticleByName(title, categoryID)
	if err != nil {
		context.SecureJSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong. Try again",
			"success": false,
		})
		return
	}

	context.SecureJSON(http.StatusOK, gin.H{
		"success": true,
		"data":    articles,
	})
}
