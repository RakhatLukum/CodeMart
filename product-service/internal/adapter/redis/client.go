package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	goredis "github.com/redis/go-redis/v9"

	"github.com/RakhatLukum/CodeMart/product-service/internal/model"
	"github.com/RakhatLukum/CodeMart/product-service/pkg/redis"
)

const (
	keyPrefix = "product:%d"
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

func (pc *Client) Set(ctx context.Context, product model.Product) error {
	data, err := json.Marshal(product)
	if err != nil {
		return fmt.Errorf("failed to marshal product: %w", err)
	}

	return pc.client.Unwrap().Set(ctx, pc.key(product.ID), data, pc.ttl).Err()
}

func (pc *Client) SetMany(ctx context.Context, products []model.Product) error {
	pipe := pc.client.Unwrap().Pipeline()
	for _, product := range products {
		data, err := json.Marshal(product)
		if err != nil {
			return fmt.Errorf("failed to marshal product: %w", err)
		}
		pipe.Set(ctx, pc.key(product.ID), data, pc.ttl)
	}
	_, err := pipe.Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to execute pipeline: %w", err)
	}
	return nil
}

func (pc *Client) Get(ctx context.Context, productID int) (model.Product, error) {
	data, err := pc.client.Unwrap().Get(ctx, pc.key(productID)).Bytes()
	if err != nil {
		if err == goredis.Nil {
			return model.Product{}, nil
		}
		return model.Product{}, fmt.Errorf("failed to get product: %w", err)
	}

	var product model.Product
	if err := json.Unmarshal(data, &product); err != nil {
		return model.Product{}, fmt.Errorf("failed to unmarshal product: %w", err)
	}
	return product, nil
}

func (pc *Client) Delete(ctx context.Context, productID int) error {
	return pc.client.Unwrap().Del(ctx, pc.key(productID)).Err()
}

func (pc *Client) key(id int) string {
	return fmt.Sprintf(keyPrefix, id)
}

func (pc *Client) GetAll(ctx context.Context) ([]model.Product, error) {
	pattern := "product:*"
	keys, err := pc.client.Unwrap().Keys(ctx, pattern).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch keys: %w", err)
	}

	if len(keys) == 0 {
		return []model.Product{}, nil
	}

	values, err := pc.client.Unwrap().MGet(ctx, keys...).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch products: %w", err)
	}

	products := make([]model.Product, 0, len(values))
	for _, val := range values {
		if val == nil {
			continue
		}
		data, ok := val.(string)
		if !ok {
			return nil, fmt.Errorf("unexpected value type in Redis")
		}
		var product model.Product
		if err := json.Unmarshal([]byte(data), &product); err != nil {
			return nil, fmt.Errorf("failed to unmarshal product: %w", err)
		}
		products = append(products, product)
	}

	return products, nil
}
