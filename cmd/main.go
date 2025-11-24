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

	// products router
	r.POST("/products", controller.CreateProduct)
	r.GET("/products", controller.GetProduct)
	r.GET("products/:id", controller.FindProduct)
	r.DELETE("/products/:id", controller.DeleteProduct)
	r.PATCH("/products/:id", controller.UpdateProduct)

	// suppliers router
	r.POST("/suppliers", controller.CreateSupp)
	r.GET("/suppliers", controller.GetSupp)
	r.GET("/suppliers/:id", controller.FindSupp)
	r.DELETE("/suppliers/:id", controller.DeleteSupp)
	r.PATCH("/suppliers/:id", controller.UpdateSupp)

	database.ConnectDB()

	r.Run(":8000")
}
