package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/solthoth/go.handsonpractice/internal/database"
	"github.com/solthoth/go.handsonpractice/internal/dberrors"
	"github.com/solthoth/go.handsonpractice/internal/models"
)

type Server interface {
    Start(port uint) error
    Readiness(ctx echo.Context) error
    Liveness(ctx echo.Context) error
}

type EchoServer struct {
    echo *echo.Echo
    DB database.DatabaseClient
}

func NewEchoServer(db database.DatabaseClient) Server {
    server := &EchoServer{
        echo: echo.New(),
        DB: db,
    }
    server.registerRoutes()
    return server
}

func (s *EchoServer) Start(port uint) error {
    if err := s.echo.Start(fmt.Sprintf(":%d", port)); err != nil && err != http.ErrServerClosed {
        log.Fatalf("server shutdown occurred: %s", err)
        return err
    }
    return nil
}

func (s *EchoServer) registerRoutes() {
    s.echo.GET("/readiness", s.Readiness)
    s.echo.GET("/liveness", s.Liveness)
    s.echo.GET("/customers", s.GetCustomers)
    s.echo.GET("/vendors", s.GetVendors)
    s.echo.GET("/services", s.GetServices)
    s.echo.GET("/products", s.GetProducts)
}

func (s *EchoServer) Readiness(ctx echo.Context) error {
    ready := s.DB.Ready()
    if ready {
        return ctx.JSON(http.StatusOK, models.Health{Status: "OK"})
    }
    return ctx.JSON(http.StatusInternalServerError, models.Health{Status: "Failure"})
}

func (s *EchoServer) Liveness(ctx echo.Context) error {
    return ctx.JSON(http.StatusOK, models.Health{Status: "OK"})
}

func (s *EchoServer) GetCustomers(ctx echo.Context) error {
    customers := s.DB.GetCustomers()
    if customers != nil {
        return ctx.JSON(http.StatusOK, customers)
    }
    return ctx.JSON(http.StatusInternalServerError, nil)
}

func (s *EchoServer) GetVendors(ctx echo.Context) error {
    customers := s.DB.GetVendors()
    if customers != nil {
        return ctx.JSON(http.StatusOK, customers)
    }
    return ctx.JSON(http.StatusInternalServerError, nil)
}

func (s *EchoServer) GetServices(ctx echo.Context) error {
    services := s.DB.GetServices()
    if services != nil {
        return ctx.JSON(http.StatusOK, services)
    }
    return ctx.JSON(http.StatusInternalServerError, nil)
}

func (s *EchoServer) GetProducts(ctx echo.Context) error {
    vendorId := ctx.QueryParams().Get("vendorId")
    products := s.DB.GetProductsByVendorId(vendorId)
    if products != nil {
        return ctx.JSON(http.StatusOK, products)
    }
    err := &dberrors.NotFoundError{
        Entity: "product",
        ID: vendorId,
    }
    return ctx.JSON(http.StatusInternalServerError, err)
}
