package main

import (
	"context"
	"github.com/d-ashesss/noter-bot/pkg/app"
	"log"
)

func main() {
	config := app.LoadConfig()
	a, err := app.NewApp(config)
	if err != nil {
		log.Fatalf("[main] failed to init the app: %s", err)
	}
	a.Run(context.Background())
}
