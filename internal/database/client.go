package database

import (
	"context"
	"fmt"
	"time"

	"github.com/solthoth/go.handsonpractice/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type DatabaseClient interface {
	Ready() bool
    GetCustomers(ctx context.Context) ([]models.Customer, error)
    GetVendors(ctx context.Context) ([]models.Vendor, error)
    GetServices(ctx context.Context) ([]models.Service, error)
    GetProducts(ctx context.Context, vendorId string) ([]models.Product, error)
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
