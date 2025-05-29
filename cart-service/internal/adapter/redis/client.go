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

const (
	keyPrefix = "cart:%d"
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

	return cc.client.Unwrap().Set(ctx, cc.key(cart.ID), data, cc.ttl).Err()
}

func (cc *Client) SetMany(ctx context.Context, carts []model.Cart) error {
	pipe := cc.client.Unwrap().Pipeline()
	for _, cart := range carts {
		data, err := json.Marshal(cart)
		if err != nil {
			return fmt.Errorf("failed to marshal cart: %w", err)
		}
		pipe.Set(ctx, cc.key(cart.ID), data, cc.ttl)
	}
	_, err := pipe.Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to execute pipeline: %w", err)
	}
	return nil
}

func (cc *Client) Get(ctx context.Context, cartID int) (model.Cart, error) {
	data, err := cc.client.Unwrap().Get(ctx, cc.key(cartID)).Bytes()
	if err != nil {
		if err == goredis.Nil {
			return model.Cart{}, nil
		}
		return model.Cart{}, fmt.Errorf("failed to get cart: %w", err)
	}

	var cart model.Cart
	if err := json.Unmarshal(data, &cart); err != nil {
		return model.Cart{}, fmt.Errorf("failed to unmarshal cart: %w", err)
	}
	return cart, nil
}

func (cc *Client) Delete(ctx context.Context, cartID int) error {
	return cc.client.Unwrap().Del(ctx, cc.key(cartID)).Err()
}

func (cc *Client) key(id int) string {
	return fmt.Sprintf(keyPrefix, id)
}

func (cc *Client) GetAll(ctx context.Context) ([]model.Cart, error) {
	pattern := "cart:*"
	keys, err := cc.client.Unwrap().Keys(ctx, pattern).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch keys: %w", err)
	}

	if len(keys) == 0 {
		return []model.Cart{}, nil
	}

	values, err := cc.client.Unwrap().MGet(ctx, keys...).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch cart values: %w", err)
	}

	carts := make([]model.Cart, 0, len(values))
	for _, val := range values {
		if val == nil {
			continue
		}
		data, ok := val.(string)
		if !ok {
			return nil, fmt.Errorf("unexpected value type in Redis")
		}
		var cart model.Cart
		if err := json.Unmarshal([]byte(data), &cart); err != nil {
			return nil, fmt.Errorf("failed to unmarshal cart: %w", err)
		}
		carts = append(carts, cart)
	}

	return carts, nil
}
