package config

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
)

func GetSESClient() *sesv2.Client {

	staticProvider := credentials.NewStaticCredentialsProvider(
		os.Getenv("AMAZON_SES_ACCESS_KEY_ID"),
		os.Getenv("AMAZON_SES_SECRET_ACCESS_KEY"),
		"",
	)

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("eu-north-1"),
		config.WithCredentialsProvider(staticProvider),
	)
	if err != nil {
		log.Fatalf("No se pudo cargar la config de AWS: %v", err)
	}

	return sesv2.NewFromConfig(cfg)
}
