package main

import (
	"context"
	"github.com/d-ashesss/noter-bot/pkg/app"
)

func main() {
	config := app.LoadConfig()
	a := app.NewApp(config)
	a.Run(context.Background())
}
