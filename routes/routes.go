package routes

import (
	"echapserver/middlewares"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes register all of the routes that are needed by the application.
func RegisterRoutes(server *gin.Engine) {

	// create a route group for protected endpoints and attach middleware to it
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/upload", uploadImage)

	// OTP related API's
	authenticated.POST("/otp-verification", verifyOTP)
	authenticated.GET("/otp", getOTP)

	// user related API's
	authenticated.GET("/user/:id", getUserByID)
	authenticated.PUT("/user/:id", updateUser)

	// Category related API's
	authenticated.POST("/category", createCategory)
	authenticated.PUT("/category/:id", updateCategory)
	authenticated.DELETE("/category/:id", deleteCategory)

	// Article
	authenticated.POST("/article", createArticle)
	authenticated.DELETE("/article/:id", deleteArticle)
	authenticated.GET("/articles/:category_id", getArticles)
	authenticated.GET("/articles/owner", getArticleByUser)

	// No need authentication
	server.GET("/categories", getCategories)
	server.GET("/category/:id", getCategoryByID)
	server.GET("/city/:countryID", getCitiesByCountry)
	server.GET("/countries", getCounties)
	server.POST("/register", register)
	server.POST("/login", login)
}
