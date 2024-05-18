package main

import (
	"echapserver/db"
	"echapserver/routes"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()                // Database initialization
	gin.SetMode(gin.DebugMode) // set the running mode (debug/production)
	server := gin.Default()    // Server initialization
	server.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}))

	port := os.Getenv("PORT")

	if port == "" {
		port = ":4000"
	}

	routes.RegisterRoutes(server) // Routes registration

	log.Fatalln(server.Run(port)) // Start the server
}
