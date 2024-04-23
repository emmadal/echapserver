package routes

import (
	"echapserver/helpers"
	"echapserver/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func reportIssues(context *gin.Context) {
	var issue models.Issues
	userID := context.GetInt64("userID")

	if err := context.ShouldBindJSON(&issue); err != nil {
		context.SecureJSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"success": false,
		})
		return
	}

	if userID != issue.UserID {
		context.SecureJSON(http.StatusForbidden, gin.H{
			"success": false,
			"message": "Not authorized to send a message",
		})
		return
	}

	ticketRef := helpers.GenerateRandomString(5)
	issue.TicketRef = ticketRef
	if err := models.CreateIssue(issue); err != nil {
		context.SecureJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to create a issue",
			"success": false,
		})
		return
	}

	context.SecureJSON(http.StatusOK, gin.H{
		"message": "Issue created. You will be contacted soon",
		"success": true,
	})
}
