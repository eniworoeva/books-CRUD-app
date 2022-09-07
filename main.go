package main

import (
	"log"
	"os"

	"github.com/eniworoeva/books-CRUD-app/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


func main()  {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.Default()

	routes.BookRoutes(router)

	router.Run(":" + port)
}