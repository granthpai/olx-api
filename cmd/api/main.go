package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/granthpai/olx-api/internal/config"
	"github.com/granthpai/olx-api/internal/db"
	"github.com/granthpai/olx-api/internal/handlers"
)

func main() {

	cfg := config.MustLoad()
	db, err :=db.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("main.db.connect: %v", err)
	}

    fmt.Println("Database connection established")
	fmt.Println("Starting olx server")

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", handlers.Health)
	mux.HandleFunc("GET /listings", handlers.List(db))//closure factory pattern

	srv := http.Server{
		Addr:   ":" + cfg.Port,
		Handler: mux,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout: 60 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("server failed: %v", err)
	}
}