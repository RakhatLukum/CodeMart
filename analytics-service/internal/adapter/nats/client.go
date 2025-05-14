package nats

import (
	"encoding/json"
	"fmt"
	"log"

	"CodeMart/analytics-service/internal/model"
	"CodeMart/analytics-service/internal/usecase"

	"github.com/nats-io/nats.go"
)

type Subscriber struct {
	conn         *nats.Conn
	viewUsecase  usecase.ViewUsecase
	subjectProds string
	subjectUsers string
	subjectCarts string
}

func NewSubscriber(
	nc *nats.Conn,
	viewUC usecase.ViewUsecase,
	subjectProds string,
	subjectUsers string,
	subjectCarts string,
) *Subscriber {
	return &Subscriber{
		conn:         nc,
		viewUsecase:  viewUC,
		subjectProds: subjectProds,
		subjectUsers: subjectUsers,
		subjectCarts: subjectCarts,
	}
}

func (s *Subscriber) Subscribe() error {
	if _, err := s.conn.Subscribe(s.subjectProds, func(msg *nats.Msg) {
		var product model.Product
		if err := json.Unmarshal(msg.Data, &product); err != nil {
			log.Printf("Failed to unmarshal product: %v", err)
			return
		}
		log.Printf("Product received: ID=%d, Name=%s, Price=%.2f", product.ID, product.Name, product.Price)
	}); err != nil {
		return fmt.Errorf("failed to subscribe to products: %w", err)
	}

	if _, err := s.conn.Subscribe(s.subjectUsers, func(msg *nats.Msg) {
		var user model.User
		if err := json.Unmarshal(msg.Data, &user); err != nil {
			log.Printf("Failed to unmarshal user: %v", err)
			return
		}
		log.Printf("User received: ID=%d, Email=%s", user.ID, user.Email)
	}); err != nil {
		return fmt.Errorf("failed to subscribe to users: %w", err)
	}

	if _, err := s.conn.Subscribe(s.subjectCarts, func(msg *nats.Msg) {
		var cart model.Cart
		if err := json.Unmarshal(msg.Data, &cart); err != nil {
			log.Printf("Failed to unmarshal cart: %v", err)
			return
		}
		log.Printf("Cart received: ID=%d, UserID=%d, ProductID=%d", cart.ID, cart.UserID, cart.ProductID)
	}); err != nil {
		return fmt.Errorf("failed to subscribe to carts: %w", err)
	}

	return nil
}
