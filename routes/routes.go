package routes

import "github.com/gin-gonic/gin"

// RegisterRoutes register all of the routes that are needed by the application.
func RegisterRoutes(server *gin.Engine) {

	server.GET("/categories", getCategories)
	server.GET("/category/:id", getCategoryByID)
	server.POST("/category", createCategory)
	server.PUT("/category/:id", updateCategory)
	server.DELETE("/category/:id", deleteCategory)
	server.POST("/upload", uploadImage)
}