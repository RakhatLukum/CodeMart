package mailer

import (
	"fmt"

	"github.com/RakhatLukum/CodeMart/cart-service/internal/model"
	"github.com/mailjet/mailjet-apiv3-go/v4"
)

type MailjetClient struct {
	client *mailjet.Client
	From   string
	Name   string
}

func NewMailjetClient(client *mailjet.Client, from string, name string) *MailjetClient {
	return &MailjetClient{
		client: client,
		From:   from,
		Name:   name,
	}
}

func (m *MailjetClient) SendCartItemAddedEmail(toEmail, toName string, cart model.Cart) error {
	subject := "Item Added to Your Cart"
	text := fmt.Sprintf("An item has been added to your cart:\n\nCart ID: %d\nUser ID: %d\nProduct ID: %d",
		cart.ID, cart.UserID, cart.ProductID)

	html := fmt.Sprintf(`
		<h3>Item Added to Cart</h3>
		<ul>
			<li><strong>Cart ID:</strong> %d</li>
			<li><strong>User ID:</strong> %d</li>
			<li><strong>Product ID:</strong> %d</li>
		</ul>`, cart.ID, cart.UserID, cart.ProductID)

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
