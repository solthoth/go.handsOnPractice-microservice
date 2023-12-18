package database

import (
	"context"

	"github.com/google/uuid"
	"github.com/solthoth/go.handsonpractice/internal/models"
)

func (c Client) GetCustomers(ctx context.Context) ([]models.Customer, error) {
    var customers []models.Customer
    result := c.DB.WithContext(ctx).Find(&customers)
    return customers, result.Error
}

func (c Client) AddCustomer(ctx context.Context, customer *models.Customer) (*models.Customer, error) {
    customer.CustomerId = uuid.NewString()
    result := c.DB.WithContext(ctx).
        Create(customer)
    if result.Error != nil {
        return nil, c.handleDuplicateError(result.Error)
    }
    return customer, nil
}

func (c Client) GetCustomerById(ctx context.Context, id string) (*models.Customer, error) {
    customer := &models.Customer{}
    result := c.DB.WithContext(ctx).Where(&models.Customer{CustomerId: id}).First(&customer)
    if result.Error != nil {
        return nil, c.handleNotFoundError("customer", id, result.Error)
    }
    return customer, nil
}