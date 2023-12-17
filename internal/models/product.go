package models

type Product struct {
    ProductId   string `gorm:"primaryKey" json:"product_id"`
    Name        string `json:"name"`
    Price       float32 `json:"price" sql:"type:decimal(12,2)"`
    VendorId    string `json:"vendor_id"`
}