package database

import (
	"context"

	"github.com/solthoth/go.handsonpractice/internal/models"
)

func (c Client) GetVendors(ctx context.Context) ([]models.Vendor, error) {
    var vendors []models.Vendor
    result := c.DB.WithContext(ctx).Find(&vendors)
    return vendors, result.Error
}
