package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	goredis "github.com/redis/go-redis/v9"

	"github.com/RakhatLukum/CodeMart/cart-service/internal/model"
	"github.com/RakhatLukum/CodeMart/cart-service/pkg/redis"
)

type Client struct {
	client *redis.Client
	ttl    time.Duration
}

func NewClient(client *redis.Client, ttl time.Duration) *Client {
	return &Client{
		client: client,
		ttl:    ttl,
	}
}

func (cc *Client) Set(ctx context.Context, cart model.Cart) error {
	data, err := json.Marshal(cart)
	if err != nil {
		return fmt.Errorf("failed to marshal cart: %w", err)
	}

	return cc.client.Unwrap().Set(ctx, cc.key(cart.UserID, cart.ProductID), data, cc.ttl).Err()
}

func (cc *Client) SetMany(ctx context.Context, carts []model.Cart) error {
	pipe := cc.client.Unwrap().Pipeline()
	for _, cart := range carts {
		data, err := json.Marshal(cart)
		if err != nil {
			return fmt.Errorf("failed to marshal cart: %w", err)
		}
		pipe.Set(ctx, cc.key(cart.UserID, cart.ProductID), data, cc.ttl)
	}
	_, err := pipe.Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to execute pipeline: %w", err)
	}
	return nil
}

func (cc *Client) Get(ctx context.Context, userID, productID int) (model.Cart, error) {
	data, err := cc.client.Unwrap().Get(ctx, cc.key(userID, productID)).Bytes()
	if err != nil {
		if err == goredis.Nil {
			return model.Cart{}, nil
		}
		return model.Cart{}, fmt.Errorf("failed to get cart item: %w", err)
	}

	var cart model.Cart
	if err := json.Unmarshal(data, &cart); err != nil {
		return model.Cart{}, fmt.Errorf("failed to unmarshal cart: %w", err)
	}
	return cart, nil
}

func (cc *Client) GetAll(ctx context.Context, userID int) ([]model.Cart, error) {
	pattern := fmt.Sprintf("cart:%d:*", userID)
	keys, err := cc.client.Unwrap().Keys(ctx, pattern).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to list keys: %w", err)
	}

	var carts []model.Cart
	for _, key := range keys {
		data, err := cc.client.Unwrap().Get(ctx, key).Bytes()
		if err != nil {
			continue
		}
		var cart model.Cart
		if err := json.Unmarshal(data, &cart); err == nil {
			carts = append(carts, cart)
		}
	}
	return carts, nil
}

func (cc *Client) Delete(ctx context.Context, userID, productID int) error {
	return cc.client.Unwrap().Del(ctx, cc.key(userID, productID)).Err()
}

func (cc *Client) key(userID, productID int) string {
	return fmt.Sprintf("%s%d:%d", "cart:", userID, productID)
}
