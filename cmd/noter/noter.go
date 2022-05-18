package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/d-ashesss/noter-bot/pkg/app"
	"github.com/d-ashesss/noter-bot/pkg/model"
	"log"
)

func main() {
	ctx := context.Background()
	config := app.LoadConfig()

	storeClient, err := firestore.NewClient(ctx, firestore.DetectProjectID)
	if err != nil {
		log.Fatalf("[main] failed to init Firestore: %v", err)
	}
	defer func() { _ = storeClient.Close() }()
	modelFactory := model.NewFirestoreFactory(storeClient)

	a, err := app.NewApp(config, modelFactory)
	if err != nil {
		log.Fatalf("[main] failed to init the app: %s", err)
	}
	a.Run(ctx)
}
