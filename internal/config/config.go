// Package config
package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	ExchageAPI string
}

func LoadConfig() (*Config, error) {
	slog.Info("Loading app configurations")
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		Port:       os.Getenv("PORT"),
		ExchageAPI: os.Getenv("EXCHANGE_API"),
	}

	return cfg, nil
}
