package main

import (
	"log"

	"github.com/mrjxtr-dev/bmo/internal/config"
)

func main() {
	_, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
}
