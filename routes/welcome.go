package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func welcome(context *gin.Context) {
	context.JSON(http.StatusOK, "Welcome on the API")
}
