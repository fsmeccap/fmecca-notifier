package main

import (
	"context"
	"fmecca-notifier/internal/config"
	"fmecca-notifier/internal/mailer"
	"fmecca-notifier/internal/worker"
	"log"

	"cloud.google.com/go/pubsub"
	"github.com/joho/godotenv" // Asegúrate de importar esto
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Aviso: No se pudo cargar el archivo .env")
	}

	ctx := context.Background()

	sesClient := config.GetSESClient()
	sesNotifier := mailer.NewSESNotifier(sesClient)

	projectID := "franco-297918"
	psClient, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Error al conectar con Pub/Sub: %v", err)
	}
	defer psClient.Close()

	subscriptionName := "fmecca-email-dispatcher-sub"
	paymentWorker := worker.NewWorker(sesNotifier, psClient, subscriptionName)

	paymentWorker.Start(ctx)
}
