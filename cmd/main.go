package main

import (
	"fmt"
	"log"

	"github.com/mrjxtr-dev/bmo/internal/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg.ExchangeRate)
}
