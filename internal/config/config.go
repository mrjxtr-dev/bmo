// Package config
package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/mrjxtr-dev/bmo/internal/services"
)

type Config struct {
	Port string
	ENV  string

	ExchageAPI string

	ExchangeRate float64
}

func LoadConfig() (*Config, error) {
	slog.Info("Loading app configurations")
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		Port:       os.Getenv("PORT"),
		ENV:        os.Getenv("ENV"),
		ExchageAPI: os.Getenv("EXCHANGE_API"),
	}

	er := &services.ExchangeRates{}
	rate, err := er.GetExchangeRate(cfg.ExchageAPI)
	if err != nil {
		return nil, err
	}

	cfg.ExchangeRate = rate

	return cfg, nil
}
