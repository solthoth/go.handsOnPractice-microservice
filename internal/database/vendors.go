package database

import (
	"context"

	"github.com/google/uuid"
	"github.com/solthoth/go.handsonpractice/internal/models"
)

func (c Client) GetVendors(ctx context.Context) ([]models.Vendor, error) {
    var vendors []models.Vendor
    result := c.DB.WithContext(ctx).Find(&vendors)
    return vendors, result.Error
}

func (c Client) AddVendor(ctx context.Context, vendor *models.Vendor) (*models.Vendor, error) {
    vendor.VendorId = uuid.NewString()
    result := c.DB.WithContext(ctx).
        Create(vendor)
    if result.Error != nil {
        return nil, c.handleDuplicateError(result.Error)
    }
    return vendor, nil
}

func (c Client) GetVendorById(ctx context.Context, id string) (*models.Vendor, error) {
    vendor := new(models.Vendor)
    result := c.DB.WithContext(ctx).Where(&models.Vendor{VendorId: id}).First(&vendor)
    if result.Error != nil {
        return nil, c.handleNotFoundError("vendor", id, result.Error)
    }
    return vendor, nil
}