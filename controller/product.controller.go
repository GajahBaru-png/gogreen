package controller

import (
	"net/http"

	"github.com/GajahBaru-png/gogreen/database"
	"github.com/GajahBaru-png/gogreen/models"
	"github.com/gin-gonic/gin"
)

// type InputProduct struct {
// 	ProductName string `json:"product_name"`
// 	Quantity    int    `json:"quantity" `
// 	Price       int    `json:"price"`
// 	Supplier    string `json:"supplier"`
// }

func CreateProduct(c *gin.Context) {
	var input models.Product

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error Cik"})
		return
	}

	database.DB.Create(&input)

	c.JSON(http.StatusOK, gin.H{"data": input})
}

func GetProduct(c *gin.Context) {
	var products []models.Product

	database.DB.Find(&products)

	c.JSON(http.StatusOK, gin.H{"data": products})
}

func DeleteProduct(c *gin.Context) {
	var products models.Product
	ProductID := c.Param("id")

	result := database.DB.First(&products, ProductID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product tidak ditemukan"})
		return
	}

	database.DB.Delete(&products)
	c.JSON(http.StatusOK, gin.H{"pesan": "Product berhasil dihapus"})
}

func FindProduct(c *gin.Context) {
	var products models.Product

	if err := database.DB.Where("id = ?", c.Param("id")).First(&products).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Product tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": products})
}
