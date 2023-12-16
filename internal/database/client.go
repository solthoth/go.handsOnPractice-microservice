package database

import (
	"fmt"
	"time"

	"github.com/solthoth/go.handsonpractice/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type DatabaseClient interface {
	Ready() bool
    GetCustomers() []models.Customer
    GetVendors() []models.Vendor
    GetServices() []models.Service
    GetProductsByVendorId(vendorId string) []models.Product
}

type Client struct {
	DB *gorm.DB
}

func NewDatabaseClient(host, username, password, databaseName, sslMode string, port uint) (DatabaseClient, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", host, username, password, databaseName, port, sslMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "wisdom.",
		},
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
		QueryFields: true,
	})
	if err != nil {
		return nil, err
	}
	client := Client{
		DB: db,
	}
	return client, nil
}

func (c Client) Ready() bool {
	var ready string
	tx := c.DB.Raw("SELECT 1 as ready").Scan(&ready)
	if tx.Error != nil {
		return false
	}
	if ready == "1" {
		return true
	}
	return false
}

func (c Client) GetCustomers() []models.Customer {
    var customers []models.Customer
    c.DB.Table("wisdom.customers").
        Select("customer_id", "first_name", "last_name", "email", "phone", "address").
        Scan(&customers)
    return customers
}

func (c Client) GetVendors() []models.Vendor {
    var vendors []models.Vendor
    c.DB.Table("wisdom.vendors").
        Select("vendor_id", "name", "contact", "phone", "email", "address").
        Scan(&vendors)
    return vendors
}

func (c Client) GetServices() []models.Service {
    var services []models.Service
    c.DB.Table("wisdom.services").
        Select("service_id", "name", "price").
        Scan(&services)
    return services
}

func (c Client) GetProductsByVendorId(vendorId string) []models.Product {
    var products []models.Product
    c.DB.Table("wisdom.products").
        Select("product_id", "name", "price", "vendor_id").
        Where("vendor_id = ?", vendorId).
        Scan(&products)
    return products
}