package main

import (
	"log"

	"github.com/solthoth/go.handsonpractice/internal/database"
	"github.com/solthoth/go.handsonpractice/internal/server"
)

func main() {
	db, err := database.NewDatabaseClient("db", "postgres", "postgres", "postgres", "disable", 5432)
    if err != nil {
        log.Fatalf("failed to initialize Database Client: %s", err)
    }
    srv := server.NewEchoServer(db)
    if err := srv.Start(8080); err != nil {
        log.Fatal(err.Error())
    }
}