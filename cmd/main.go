package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/mrjxtr-dev/bmo/internal/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	er := &ExchangeRates{}
	rate, err := er.getExchangeRate(cfg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rate)
}

type ExchangeRates struct {
	Date string             `json:"date"`
	USD  map[string]float64 `json:"usd"`
}

// getExchangeRate gets USD -> PHP exchange rate
func (er *ExchangeRates) getExchangeRate(cfg *config.Config) (float64, error) {
	endpoint := cfg.ExchageAPI
	resp, err := http.Get(endpoint)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(&er); err != nil {
		return 0, err
	}

	php, ok := er.USD["php"]
	if !ok {
		return 0, fmt.Errorf("php rate not found")
	}

	return php, nil
}
