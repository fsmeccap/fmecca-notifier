package mailer

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"

	"fmecca-notifier/internal/domain"
)

type SESNotifier struct {
	client *sesv2.Client
}

func NewSESNotifier(client *sesv2.Client) *SESNotifier {
	return &SESNotifier{client: client}
}

func (n *SESNotifier) SendNotification(req domain.EmailRequest) error {
	fromName := req.FromName
	if fromName == "" {
		fromName = os.Getenv("SES_FROM_NAME")
	}

	fromEmail := req.FromEmail
	if fromEmail == "" {
		fromEmail = os.Getenv("SES_FROM_EMAIL")
	}

	jsonData, _ := json.Marshal(req.TemplateData)

	input := &sesv2.SendEmailInput{
		FromEmailAddress: aws.String(fmt.Sprintf("%s <%s>", fromName, fromEmail)),
		Destination:      &types.Destination{ToAddresses: []string{req.To}},
		Content: &types.EmailContent{
			Template: &types.Template{
				TemplateName: aws.String(req.TemplateName),
				TemplateData: aws.String(string(jsonData)),
			},
		},
	}

	_, err := n.client.SendEmail(context.TODO(), input)
	return err
}
