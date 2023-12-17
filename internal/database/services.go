package database

import (
	"context"

	"github.com/google/uuid"
	"github.com/solthoth/go.handsonpractice/internal/models"
)

func (c Client) GetServices(ctx context.Context) ([]models.Service, error) {
    var services []models.Service
    result := c.DB.WithContext(ctx).Find(&services)
    return services, result.Error
}

func (c Client) AddService(ctx context.Context, service *models.Service) (*models.Service, error) {
    service.ServiceId = uuid.NewString()
    result := c.DB.WithContext(ctx).
        Create(service)
    if result.Error != nil {
        return nil, c.handleDuplicateError(result.Error)
    }
    return service, nil
}