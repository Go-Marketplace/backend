package main

import (
	"log"

	"github.com/Go-Marketplace/backend/gateway/internal/app"
	"github.com/Go-Marketplace/backend/config"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("failed to create new config: %s", err)
	}

	app.Run(cfg)
}
