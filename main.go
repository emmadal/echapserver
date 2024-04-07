package main

import (
	"echapserver/db"
	"echapserver/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var mode = os.Getenv("GIN_MODE")

func main() {
	db.InitDB()             // Database initialization
	server := gin.Default() // Server initialization
	gin.SetMode(mode)       // set the running mode (debug/production)

	port := os.Getenv("PORT")

	if port == "" {
		port = ":4000"
	}

	routes.RegisterRoutes(server) // Routes registration

	err := server.Run(port) // Start the server
	if err != nil {
		log.Fatalln(err)
	}
}
