package nats

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/RakhatLukum/CodeMart/cart-service/internal/model"
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

type CartEvent struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

func (p *Publisher) PublishCartCreated(cart model.Cart) error {
	return p.publish("cart", CartEvent{
		Type:    "created",
		Payload: cart,
	})
}

func (p *Publisher) PublishCartDeleted(cart model.Cart) error {
	return p.publish("cart", CartEvent{
		Type:    "deleted",
		Payload: cart,
	})
}

func (p *Publisher) publish(subject string, message CartEvent) error {
	data, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal cart event: %w", err)
	}

	if err := p.conn.Publish(subject, data); err != nil {
		return fmt.Errorf("failed to publish cart event: %w", err)
	}

	log.Printf("[NATS] Published to subject %s: type=%s", subject, message.Type)
	return nil
}
