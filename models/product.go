package models

type Product struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity" binding:"gte=0" gorm:"not null;default:0check:quantity >= 0"`
	Price       int    `json:"price" binding:"gte=0" gorm:"not null;default:0check:price >= 0"`
	Supplier    string `json:"supplier" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	SupplierID  uint   `json:"supplier_id"`
}
