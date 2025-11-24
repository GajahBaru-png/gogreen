package controller

import (
	"net/http"

	"github.com/GajahBaru-png/gogreen/database"
	"github.com/GajahBaru-png/gogreen/models"
	"github.com/gin-gonic/gin"
)

func CreateSupp(c *gin.Context) {
	var input models.Supplier

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.Products = []models.Product{}

	database.DB.Create(&input)
	c.JSON(http.StatusOK, gin.H{"data": input})
}

func GetSupp(c *gin.Context) {
	var supp []models.Supplier

	database.DB.Find(&supp)

	c.JSON(http.StatusOK, gin.H{"data": supp})
}

func DeleteSupp(c *gin.Context) {
	var supps models.Supplier
	SupplierID := c.Param("id")

	result := database.DB.First(&supps, SupplierID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data ne gak ono"})
		return
	}

	database.DB.Delete(&supps)
	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil di hapus"})
}

func FindSupp(c *gin.Context) {
	var suppliers models.Supplier

	if err := database.DB.Where("id = ?", c.Param("id")).First(&suppliers).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Cikibul"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": suppliers})
}

type SuppUpdate struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

func UpdateSupp(c *gin.Context) {
	var usupp models.Supplier

	if err := database.DB.Where("id = ?", c.Param("id")).First(&usupp).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
		return
	}

	var input SuppUpdate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedSupp := models.Supplier{Name: input.Name, Address: input.Address, Phone: input.Phone}
	database.DB.Model(&usupp).Updates(&updatedSupp)
	c.JSON(http.StatusOK, gin.H{"data": usupp})
}
