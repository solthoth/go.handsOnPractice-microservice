package database

import (
	"context"

	"github.com/solthoth/go.handsonpractice/internal/models"
)

func (c Client) GetProducts(ctx context.Context, vendorId string) ([]models.Product, error) {
    var products []models.Product
    result := c.DB.WithContext(ctx).
        Where(models.Product{VendorId: vendorId}).
        Find(&products)
    return products, result.Error
}