package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/solthoth/go.handsonpractice/internal/models"
)

func (s *EchoServer) GetServices(ctx echo.Context) error {
    services, err := s.DB.GetServices(ctx.Request().Context())
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, err)
    }
    return ctx.JSON(http.StatusOK, services)    
}

func (s *EchoServer) AddService(ctx echo.Context) error {
    service := new(models.Service)
    if err := ctx.Bind(service); err != nil {
        return ctx.JSON(http.StatusUnsupportedMediaType, err)
    }
    service, err := s.DB.AddService(ctx.Request().Context(), service)
    if err != nil {
        return s.handleConflictError(ctx, err)
    }
    return ctx.JSON(http.StatusCreated, service)
}