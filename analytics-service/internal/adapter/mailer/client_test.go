package mailjet

import (
	"errors"
	"testing"

	"github.com/mailjet/mailjet-apiv3-go/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockMailjetClient struct {
	mock.Mock
}

func (m *mockMailjetClient) SendMailV31(messages *mailjet.MessagesV31) (*mailjet.ResultsV31, error) {
	args := m.Called(messages)
	return args.Get(0).(*mailjet.ResultsV31), args.Error(1)
}

func (m *mockMailjetClient) APIKeyPublic() string {
	return "test-api-key"
}

func (m *mockMailjetClient) APIKeyPrivate() string {
	return "test-secret-key"
}

func (m *mockMailjetClient) Client() interface{} {
	return nil
}

func (m *mockMailjetClient) SetClient(client interface{}) {}

func TestNewMailjetClient(t *testing.T) {
	mockClient := &mockMailjetClient{}
	client := NewMailjetClient(&mailjet.Client{}, "test@example.com", "Test Sender")
	assert.Equal(t, "test@example.com", client.From)
	assert.Equal(t, "Test Sender", client.Name)
	assert.Equal(t, mockClient, client.client)
}

func TestSendEmail_Success(t *testing.T) {
	mockClient := &mockMailjetClient{}
	mockClient.On("SendMailV31", mock.Anything).Return(&mailjet.ResultsV31{}, nil)

	client := NewMailjetClient(&mailjet.Client{}, "from@example.com", "Sender")
	err := client.SendEmail("to@example.com", "Recipient", "Subject", "Text", "<html>HTML</html>")

	assert.NoError(t, err)
	mockClient.AssertExpectations(t)
}

func TestSendEmail_Failure(t *testing.T) {
	mockClient := &mockMailjetClient{}
	mockClient.On("SendMailV31", mock.Anything).Return(&mailjet.ResultsV31{}, errors.New("send error"))

	client := NewMailjetClient(&mailjet.Client{}, "from@example.com", "Sender")
	err := client.SendEmail("to@example.com", "Recipient", "Subject", "Text", "<html>HTML</html>")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to send email")
	mockClient.AssertExpectations(t)
}

func TestSendDailyReportEmail_Success(t *testing.T) {
	mockClient := &mockMailjetClient{}
	mockClient.On("SendMailV31", mock.Anything).Return(&mailjet.ResultsV31{}, nil)

	client := NewMailjetClient(&mailjet.Client{}, "from@example.com", "Sender")
	report := "Sample report content"
	err := client.SendDailyReportEmail("to@example.com", "Recipient", report)

	assert.NoError(t, err)
	mockClient.AssertExpectations(t)
}

func TestSendDailyReportEmail_Failure(t *testing.T) {
	mockClient := &mockMailjetClient{}
	mockClient.On("SendMailV31", mock.Anything).Return(&mailjet.ResultsV31{}, errors.New("send error"))

	client := NewMailjetClient(&mailjet.Client{}, "from@example.com", "Sender")
	report := "Sample report content"
	err := client.SendDailyReportEmail("to@example.com", "Recipient", report)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to send email")
	mockClient.AssertExpectations(t)
}

func TestSendEmail_ValidatesInput(t *testing.T) {
	mockClient := &mockMailjetClient{}
	client := NewMailjetClient(&mailjet.Client{}, "from@example.com", "Sender")

	tests := []struct {
		name      string
		toEmail   string
		toName    string
		subject   string
		textBody  string
		htmlBody  string
		expectErr bool
	}{
		{"EmptyToEmail", "", "Name", "Subject", "Text", "HTML", true},
		{"EmptySubject", "to@test.com", "Name", "", "Text", "HTML", true},
		{"ValidInput", "to@test.com", "Name", "Subject", "Text", "HTML", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.expectErr {
				mockClient.On("SendMailV31", mock.Anything).Return(&mailjet.ResultsV31{}, nil)
			}

			err := client.SendEmail(tt.toEmail, tt.toName, tt.subject, tt.textBody, tt.htmlBody)

			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				mockClient.AssertExpectations(t)
			}
		})
	}
}
