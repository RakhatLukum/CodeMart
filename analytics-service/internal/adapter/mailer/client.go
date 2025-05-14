package mailjet

import (
	"fmt"

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

func (m *MailjetClient) SendDailyReportEmail(toEmail, toName, report string) error {
	subject := "Daily View Report"
	text := report
	html := "<h3>Daily View Report</h3><pre>" + report + "</pre>"

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
