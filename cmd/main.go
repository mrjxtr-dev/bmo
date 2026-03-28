package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mrjxtr-dev/bmo/internal/config"
)

func main() {
	slog.Info("starting server")
	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Error("failed to load config", err)
	}

	r := chi.NewRouter()
	srv := http.Server{
		Addr:    ":" + cfg.Port,
		Handler: r,
	}

	go func() {
		slog.Info("server started")
		slog.Info("server started")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
		slog.Info("server closed")
	}()

	fmt.Println(cfg.ExchangeRate)
}
