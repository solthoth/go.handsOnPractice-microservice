package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/solthoth/go.handsonpractice/internal/dberrors"
	"github.com/solthoth/go.handsonpractice/internal/models"
)

func (s *EchoServer) GetCustomers(ctx echo.Context) error {
    customers, err := s.DB.GetCustomers(ctx.Request().Context())
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, err)    
    }
    return ctx.JSON(http.StatusOK, customers)    
}

func (s *EchoServer) AddCustomer(ctx echo.Context) error {
    customer := new(models.Customer)
    if err := ctx.Bind(customer); err != nil {
        return ctx.JSON(http.StatusUnsupportedMediaType, err)
    }
    customer, err := s.DB.AddCustomer(ctx.Request().Context(), customer)
    if err != nil {
        return s.handleConflictError(ctx, err)
    }
    return ctx.JSON(http.StatusCreated, customer)
}

func (s *EchoServer) GetCustomerById(ctx echo.Context) error {
    ID := ctx.Param("id")
    customer, err := s.DB.GetCustomerById(ctx.Request().Context(), ID)
    if err != nil {
        return s.handleNotFoundError(ctx, err)
    }
    return ctx.JSON(http.StatusOK, customer)
}