package main

import (
	"log"

	"github.com/Go-Marketplace/backend/config"
	"github.com/Go-Marketplace/backend/order/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("failed to create new config: %s", err)
	}

	app.Run(cfg)
}
