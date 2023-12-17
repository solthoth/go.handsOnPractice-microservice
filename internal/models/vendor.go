package models

type Vendor struct {
    VendorId    string `gorm:"primaryKey" json:"vendor_id"`
    Name        string `json:"name"`
    Contact     string `json:"contact"`
    Phone       string `json:"phone"`
    Email       string `json:"email"`
}