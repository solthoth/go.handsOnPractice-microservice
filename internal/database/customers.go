package database

import (
	"context"

	"github.com/solthoth/go.handsonpractice/internal/models"
)

func (c Client) GetCustomers(ctx context.Context) ([]models.Customer, error) {
    var customers []models.Customer
    result := c.DB.WithContext(ctx).Find(&customers)
    return customers, result.Error
}