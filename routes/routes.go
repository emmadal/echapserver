package routes

import (
	"github.com/gin-gonic/gin"
	"oblackserver/middlewares"
)

// RegisterRoutes register all of the routes that are needed by the application.
func RegisterRoutes(server *gin.Engine) {

	// create a route group for protected endpoints and attach middleware to it
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/category", createCategory)
	authenticated.PUT("/category/:id", updateCategory)
	authenticated.DELETE("/category/:id", deleteCategory)
	authenticated.POST("/upload", uploadImage)
	authenticated.POST("/otp-verification", verifyOTP)
	authenticated.GET("/user/:id", getUserByID)
	authenticated.GET("/otp", getOTP)
	authenticated.POST("/article", createArticle)



	server.GET("/categories", getCategories)
	server.GET("/category/:id", getCategoryByID)
	server.POST("/register", register)
	server.POST("/login", login)
}