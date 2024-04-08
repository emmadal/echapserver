package routes

import (
	"echapserver/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func subscription(context *gin.Context) {
	userID := context.GetInt64("userID")
	var subscription models.Subscription
	err := context.ShouldBindJSON(&subscription)

	if err != nil {
		context.SecureJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	user, err := models.FindUserByID(subscription.UserID)

	if err != nil {
		context.SecureJSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Data not found",
		})
		return
	}

	if user.ID != userID {
		context.SecureJSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Not authorized to continue this action",
		})
		return
	}

	err = models.PremiumOffer(subscription.Premium, user.ID)
	if err != nil {
		context.SecureJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Something went Wrong",
		})
		return
	}

	if !subscription.Premium {
		context.SecureJSON(http.StatusOK, gin.H{
			"success": true,
			"message": "You have unsubscribe from premium service.",
		})
	} else {
		context.SecureJSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Subscription created successfully!",
		})

	}
}
