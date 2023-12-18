package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/solthoth/go.handsonpractice/internal/models"
)

func (s *EchoServer) GetProducts(ctx echo.Context) error {
    vendorId := ctx.QueryParam("vendorId")
    products, err := s.DB.GetProducts(ctx.Request().Context(), vendorId)
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, err)
    }
    return ctx.JSON(http.StatusOK, products)
}

func (s *EchoServer) AddProduct(ctx echo.Context) error {
    product := new(models.Product)
    if err := ctx.Bind(product); err != nil {
        return ctx.JSON(http.StatusUnsupportedMediaType, err)
    }
    product, err := s.DB.AddProduct(ctx.Request().Context(), product)
    if err != nil {
        return s.handleConflictError(ctx, err)
    }
    return ctx.JSON(http.StatusCreated, product)
}

func (s *EchoServer) GetProductById(ctx echo.Context) error {
    ID := ctx.Param("id")
    product, err := s.DB.GetProductById(ctx.Request().Context(), ID)
    if err != nil {
        return s.handleNotFoundError(ctx, err)
    }
    return ctx.JSON(http.StatusOK, product)
}