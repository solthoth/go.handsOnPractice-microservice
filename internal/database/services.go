package database

import (
	"context"

	"github.com/solthoth/go.handsonpractice/internal/models"
)

func (c Client) GetServices(ctx context.Context) ([]models.Service, error) {
    var services []models.Service
    result := c.DB.WithContext(ctx).Find(&services)
    return services, result.Error
}
