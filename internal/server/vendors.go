package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/solthoth/go.handsonpractice/internal/models"
)

func (s *EchoServer) GetVendors(ctx echo.Context) error {
    vendors, err := s.DB.GetVendors(ctx.Request().Context())
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, err)
    }
    return ctx.JSON(http.StatusOK, vendors)
}

func (s *EchoServer) AddVendor(ctx echo.Context) error {
    vendor := new(models.Vendor)
    if err := ctx.Bind(vendor); err != nil {
        return ctx.JSON(http.StatusUnsupportedMediaType, err)
    }
    vendor, err := s.DB.AddVendor(ctx.Request().Context(), vendor)
    if err != nil {
        return s.handleConflictError(ctx, err)
    }
    return ctx.JSON(http.StatusCreated, vendor)
}

func (s *EchoServer) GetVendorById(ctx echo.Context) error {
    ID := ctx.Param("id")
    vendor, err := s.DB.GetVendorById(ctx.Request().Context(), ID)
    if err != nil {
        return s.handleNotFoundError(ctx, err)
    }
    return ctx.JSON(http.StatusOK, vendor)
}