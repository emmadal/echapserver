package main

import (
	"log"
	"oblackserver/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB() // Database initialization

	server := gin.Default()

	err := server.Run(":8080") // listen server on port 8080
	if err != nil {
		log.Fatalln(err)
	}
}
