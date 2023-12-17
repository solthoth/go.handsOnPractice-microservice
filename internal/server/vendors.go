package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *EchoServer) GetVendors(ctx echo.Context) error {
    vendors, err := s.DB.GetVendors(ctx.Request().Context())
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, err)
    }
    return ctx.JSON(http.StatusOK, vendors)
}
