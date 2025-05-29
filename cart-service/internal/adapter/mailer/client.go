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

type Mailer interface {
	SendCartSummaryEmail(toEmail, toName string, carts []model.Cart, products []model.Product) error
}

func (m *MailjetClient) SendCartSummaryEmail(toEmail, toName string, carts []model.Cart, products []model.Product) error {
	subject := "Your Cart Summary"
	text := "Here is the summary of your cart:\n\n"
	html := `<h3>Your Cart Summary</h3><ul>`

	for i, cart := range carts {
		product := products[i]
		text += fmt.Sprintf("- CartID: %d, Product: %s, Price: %.2f, ProductID: %d\n",
			cart.ID, product.Name, product.Price, product.ID)

		html += fmt.Sprintf(`
			<li>
				<strong>Product:</strong> %s<br/>
				<strong>Price:</strong> %.2f<br/>
				<strong>Product ID:</strong> %d<br/>
				<strong>Cart ID:</strong> %d
			</li><br/>
		`, product.Name, product.Price, product.ID, cart.ID)
	}

	html += `</ul>`

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
