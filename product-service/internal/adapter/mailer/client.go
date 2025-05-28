package mail

import (
	"fmt"

	"github.com/RakhatLukum/CodeMart/product-service/internal/model"
	"github.com/mailjet/mailjet-apiv3-go/v4"
)

type MailjetClient struct {
	client *mailjet.Client
	From   string
	Name   string
}

type Mailer interface {
	SendProductCreatedEmail(toEmail, toName string, product model.Product) error
}

func NewMailjetClient(client *mailjet.Client, from string, name string) *MailjetClient {
	return &MailjetClient{
		client: client,
		From:   from,
		Name:   name,
	}
}

func (m *MailjetClient) SendProductCreatedEmail(toEmail, toName string, product model.Product) error {
	subject := fmt.Sprintf("New Product Created: %s", product.Name)
	text := fmt.Sprintf("A new product has been added:\n\nID: %d\nName: %s\nPrice: %.2f\nTags: %s",
		product.ID, product.Name, product.Price, product.Tags)

	html := fmt.Sprintf(`
		<h3>New Product Added</h3>
		<ul>
			<li><strong>ID:</strong> %d</li>
			<li><strong>Name:</strong> %s</li>
			<li><strong>Price:</strong> %.2f</li>
			<li><strong>Tags:</strong> %s</li>
		</ul>`, product.ID, product.Name, product.Price, product.Tags)

	return m.SendEmail(toEmail, toName, subject, text, html)
}

func (m *MailjetClient) SendEmail(toEmail, toName, subject, textBody, htmlBody string) error {
	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: m.From,
				Name:  m.Name,
			},
			To: &mailjet.RecipientsV31{
				{
					Email: toEmail,
					Name:  toName,
				},
			},
			Subject:  subject,
			TextPart: textBody,
			HTMLPart: htmlBody,
		},
	}

	messages := mailjet.MessagesV31{Info: messagesInfo}
	_, err := m.client.SendMailV31(&messages)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}
	return nil
}
