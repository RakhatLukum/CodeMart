package nats

import "github.com/RakhatLukum/CodeMart/product-service/internal/model"

type PublisherInterface interface {
	PublishProductCreated(product model.Product) error
	PublishProductUpdated(product model.Product) error
	PublishProductDeleted(productID int) error
}
