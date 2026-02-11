package domain

type EmailRequest struct {
	To           string            `json:"to"`
	TemplateName string            `json:"template_name"`
	FromName     string            `json:"from_name"`
	FromEmail    string            `json:"from_email"`
	TemplateData map[string]string `json:"template_data"`
}

type Notifier interface {
	SendNotification(event EmailRequest) error
}
