package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"oblackserver/db"
	"oblackserver/routes"
)

func main() {
	db.InitDB() // Database initialization

	server := gin.Default() // Server initialization

	routes.RegisterRoutes(server) // Routes registration

	err := server.Run(":8080") // Start the server on port 8080
	if err != nil {
		log.Fatalln(err)
	}
}
