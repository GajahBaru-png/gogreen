package models

type Supplier struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	Phone     string    `json:"phone"`
	Products  []Product `json:"products" gorm:"foreignKey:SupplierID"`
}