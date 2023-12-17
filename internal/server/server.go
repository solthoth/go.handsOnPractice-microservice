package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/solthoth/go.handsonpractice/internal/database"
	"github.com/solthoth/go.handsonpractice/internal/models"
)

type Server interface {
    Start(port uint) error
    Readiness(ctx echo.Context) error
    Liveness(ctx echo.Context) error

    // Customers
    GetCustomers(ctx echo.Context) error
    AddCustomer(ctx echo.Context) error

    // Products
    GetProducts(ctx echo.Context) error

    // Services
    GetServices(ctx echo.Context) error

    // Vendors
    GetVendors(ctx echo.Context) error
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

    customers := s.echo.Group("/customers")
    customers.GET("", s.GetCustomers)
    customers.POST("", s.AddCustomer)    

    vendors := s.echo.Group("/vendors")
    vendors.GET("", s.GetVendors)

    services := s.echo.Group("/services")
    services.GET("", s.GetServices)
    
    products := s.echo.Group("/products")
    products.GET(":vendorId", s.GetProducts)
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
