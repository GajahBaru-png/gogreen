package main

import (
	"github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "github.com/GajahBaru-png/gogreen/database"
    "log"
)

func main() {
    err := godotenv.Load()

    if err != nil {
        log.Fatal("Error Loading .env File")
    }
    r := gin.Default()

    database.ConnectDB()
    
	r.Run(":8000")
}
