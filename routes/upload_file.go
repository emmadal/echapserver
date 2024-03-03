package routes

import (
	"net/http"
	"oblackserver/helpers"

	"github.com/gin-gonic/gin"
)

func uploadImage(context *gin.Context) {
	// return the first file for the provided form key.
	header, err := context.FormFile("file") 

	if err != nil {
		context.SecureJSON(http.StatusBadRequest, gin.H{"message": "Bad image format", "success": false})
		return
	}
	url, err := helpers.UploadHelper(header)
	if err != nil {
		context.SecureJSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to uploaded file",
			"success": false,
		})
		return
	}
	context.SecureJSON(http.StatusOK, gin.H{
		"message": "file uploaded",
		"success": true,
		"data":    url,
	})
}
