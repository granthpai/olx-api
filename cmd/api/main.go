package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/granthpai/olx-api/internal/config"
)

func main() {

	cfg := config.MustLoad()

	fmt.Println("Starting olx server")

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)

		w.Write([]byte(`{"status": "healthy"}`))
	})

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