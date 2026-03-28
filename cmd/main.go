package main

import (
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
	slog.Info("config loaded", "env", cfg.ENV)

	r := chi.NewRouter()
	srv := http.Server{
		Addr:    ":" + cfg.Port,
		Handler: r,
	}

	go func() {
		slog.Info("server started", "url", "http://localhost:"+cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
		slog.Info("server closed")
	}()
}
