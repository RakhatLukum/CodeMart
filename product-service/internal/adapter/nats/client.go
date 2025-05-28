package nats

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/RakhatLukum/CodeMart/product-service/internal/model"
	"github.com/nats-io/nats.go"
)

type Publisher struct {
	conn *nats.Conn
}

func NewPublisher(nc *nats.Conn) *Publisher {
	return &Publisher{
		conn: nc,
	}
}

type ProductEvent struct {
	Type    string `json:"type"`
	Payload any    `json:"payload"`
}

func (p *Publisher) PublishProductCreated(product model.Product) error {
	return p.publish("product", ProductEvent{
		Type:    "created",
		Payload: product,
	})
}

func (p *Publisher) PublishProductUpdated(product model.Product) error {
	return p.publish("product", ProductEvent{
		Type:    "updated",
		Payload: product,
	})
}

func (p *Publisher) PublishProductDeleted(productID int) error {
	return p.publish("product", ProductEvent{
		Type:    "deleted",
		Payload: map[string]int{"id": productID},
	})
}

func (p *Publisher) publish(subject string, message ProductEvent) error {
	data, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal product event: %w", err)
	}

	if err := p.conn.Publish(subject, data); err != nil {
		return fmt.Errorf("failed to publish product event: %w", err)
	}

	log.Printf("[NATS] Published to subject %s: type=%s", subject, message.Type)
	return nil
}
