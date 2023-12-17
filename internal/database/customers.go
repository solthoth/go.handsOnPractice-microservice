package database

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/solthoth/go.handsonpractice/internal/dberrors"
	"github.com/solthoth/go.handsonpractice/internal/models"
	"gorm.io/gorm"
)

func (c Client) GetCustomers(ctx context.Context) ([]models.Customer, error) {
    var customers []models.Customer
    result := c.DB.WithContext(ctx).Find(&customers)
    return customers, result.Error
}

func (c Client) AddCustomer(ctx context.Context, customer *models.Customer) (*models.Customer, error) {
    customer.CustomerId = uuid.NewString()
    reuslt := c.DB.WithContext(ctx).
        Create(customer)
    if reuslt.Error != nil {
        if errors.Is(reuslt.Error, gorm.ErrDuplicatedKey) {
            return nil, &dberrors.ConflictError{}
        }
        return nil, reuslt.Error
    }
    return customer, nil
}