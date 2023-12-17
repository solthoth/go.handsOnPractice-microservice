package database

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/solthoth/go.handsonpractice/internal/dberrors"
	"github.com/solthoth/go.handsonpractice/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type DatabaseClient interface {
	Ready() bool

    // Customers
    GetCustomers(ctx context.Context) ([]models.Customer, error)
    AddCustomer(ctx context.Context, customer *models.Customer) (*models.Customer, error)
    
	// Vendors
    GetVendors(ctx context.Context) ([]models.Vendor, error)
	AddVendor(ctx context.Context, vendor *models.Vendor) (*models.Vendor, error)

	// Services
    GetServices(ctx context.Context) ([]models.Service, error)
	AddService(ctx context.Context, service *models.Service) (*models.Service, error)

	// Products
    GetProducts(ctx context.Context, vendorId string) ([]models.Product, error)
	AddProduct(ctx context.Context, product *models.Product) (*models.Product, error)
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

func (c Client) handleDuplicateError(resultError error) error {
	if errors.Is(resultError, gorm.ErrDuplicatedKey) {
		return &dberrors.ConflictError{}
	}
	return resultError
}