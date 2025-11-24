package controller

import (
	"net/http"

	"github.com/GajahBaru-png/gogreen/database"
	"github.com/GajahBaru-png/gogreen/models"
	"github.com/gin-gonic/gin"
)

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

	database.DB.Preload("Supplier").Find(&products)

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

	if err := database.DB.Preload("Supplier").Where("id = ?", c.Param("id")).First(&products).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Product tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": products})
}

type InputProduct struct {
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity" binding:"gte=0" gorm:"not null;default:0;check:quantity >= 0"`
	Price       int    `json:"price" binding:"gte=0" gorm:"not null;default:0;check:price >= 0"`
	SupplierID  uint   `json:"supplier_id"`
}

func UpdateProduct(c *gin.Context) {
	var product models.Product

	if err := database.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
		return
	}

	var input InputProduct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedProduct := models.Product{ProductName: input.ProductName, Quantity: input.Quantity, Price: input.Price, SupplierID: input.SupplierID}
	database.DB.Model(&product).Updates(&updatedProduct)
	c.JSON(http.StatusOK, gin.H{"data": product})
}
