package routes

import "github.com/gin-gonic/gin"

// RegisterRoutes register all of the routes that are needed by the application.
func RegisterRoutes(server *gin.Engine) {

	server.GET("/article", getCategories)

}