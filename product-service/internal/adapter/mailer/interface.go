package mailer

import "github.com/RakhatLukum/CodeMart/product-service/internal/model"

type Mailer interface {
	SendProductCreatedEmail(toEmail, toName string, product model.Product) error
}
