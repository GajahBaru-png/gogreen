package main

import (
	"log"

	"github.com/GajahBaru-png/gogreen/controller"
	"github.com/GajahBaru-png/gogreen/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error Loading .env File")
	}
	r := gin.Default()

	r.POST("/products", controller.CreateProduct)
	r.GET("/products", controller.GetProduct)
	r.GET("products/:id", controller.FindProduct)
	r.DELETE("/products/:id", controller.DeleteProduct)
	database.ConnectDB()

	r.Run(":8000")
}
