package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *EchoServer) GetCustomers(ctx echo.Context) error {
    customers, err := s.DB.GetCustomers(ctx.Request().Context())
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, err)    
    }
    return ctx.JSON(http.StatusOK, customers)    
}