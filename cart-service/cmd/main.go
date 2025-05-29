package main

import (
	"context"
	"log"

	"github.com/RakhatLukum/CodeMart/cart-service/config"
	"github.com/RakhatLukum/CodeMart/cart-service/internal/app"
)

func main() {
	ctx := context.Background()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Printf("failed to parse config: %v", err)
		return
	}

	application, err := app.New(ctx, cfg)
	if err != nil {
		log.Printf("failed to setup application: %v", err)
		return
	}

	err = application.Run()
	if err != nil {
		log.Printf("failed to run application: %v", err)
		return
	}
}
