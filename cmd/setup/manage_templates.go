package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"fmecca-notifier/internal/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
)

func main() {
	client := config.GetSESClient()

	content, err := os.ReadFile("scripts/templates/fm-biblioteca-createduser.json")
	if err != nil {
		log.Fatalf("No se pudo leer el archivo: %v", err)
	}

	var templateData struct {
		TemplateName    string `json:"TemplateName"`
		TemplateContent struct {
			Subject string `json:"Subject"`
			Html    string `json:"Html"`
			Text    string `json:"Text"`
		} `json:"TemplateContent"`
	}

	if err := json.Unmarshal(content, &templateData); err != nil {
		log.Fatalf("Error al parsear JSON: %v", err)
	}

	input := &sesv2.CreateEmailTemplateInput{
		TemplateName: aws.String(templateData.TemplateName),
		TemplateContent: &types.EmailTemplateContent{
			Subject: aws.String(templateData.TemplateContent.Subject),
			Html:    aws.String(templateData.TemplateContent.Html),
			Text:    aws.String(templateData.TemplateContent.Text),
		},
	}

	_, err = client.CreateEmailTemplate(context.TODO(), input)
	if err != nil {
		fmt.Printf("Aviso: Probablemente la plantilla ya existe, intentando actualizar... \n")
	} else {
		fmt.Println("🚀 Plantilla creada con éxito en Amazon SES!")
	}
}
