package database

import (
	"context"

	"github.com/google/uuid"
	"github.com/solthoth/go.handsonpractice/internal/models"
)

func (c Client) GetProducts(ctx context.Context, vendorId string) ([]models.Product, error) {
    var products []models.Product
    result := c.DB.WithContext(ctx).
        Where(models.Product{VendorId: vendorId}).
        Find(&products)
    return products, result.Error
}

func (c Client) AddProduct(ctx context.Context, product *models.Product) (*models.Product, error) {
    product.ProductId = uuid.NewString()
    result := c.DB.WithContext(ctx).Create(product)
    if result.Error != nil {
        return nil, c.handleDuplicateError(result.Error)
    }
    return product, nil
}