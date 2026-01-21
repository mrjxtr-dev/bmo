// Package services
package services

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ExchangeRates struct {
	Date string             `json:"date"`
	USD  map[string]float64 `json:"usd"`
}

// GetExchangeRate gets USD -> PHP exchange rate
func (er *ExchangeRates) GetExchangeRate(endpoint string) (float64, error) {
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

	rate, ok := er.USD["php"]
	if !ok {
		return 0, fmt.Errorf("php rate not found")
	}

	return rate, nil
}
