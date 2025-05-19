package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	goredis "github.com/redis/go-redis/v9"

	"github.com/RakhatLukum/CodeMart/analytics-service/internal/model"
	"github.com/RakhatLukum/CodeMart/analytics-service/pkg/redis"
)

const (
	keyPrefix = "view:%d"
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

func (vc *Client) Set(ctx context.Context, view model.View) error {
	data, err := json.Marshal(view)
	if err != nil {
		return fmt.Errorf("failed to marshal view: %w", err)
	}

	return vc.client.Unwrap().Set(ctx, vc.key(view.ProductID), data, vc.ttl).Err()
}

func (vc *Client) SetMany(ctx context.Context, views []model.View) error {
	pipe := vc.client.Unwrap().Pipeline()
	for _, view := range views {
		data, err := json.Marshal(view)
		if err != nil {
			return fmt.Errorf("failed to marshal view: %w", err)
		}
		pipe.Set(ctx, vc.key(view.ProductID), data, vc.ttl)
	}
	_, err := pipe.Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to execute pipeline: %w", err)
	}
	return nil
}

func (vc *Client) Get(ctx context.Context, productID int) (model.View, error) {
	data, err := vc.client.Unwrap().Get(ctx, vc.key(productID)).Bytes()
	if err != nil {
		if err == goredis.Nil {
			return model.View{}, nil
		}
		return model.View{}, fmt.Errorf("failed to get view: %w", err)
	}

	var view model.View
	if err := json.Unmarshal(data, &view); err != nil {
		return model.View{}, fmt.Errorf("failed to unmarshal view: %w", err)
	}
	return view, nil
}

func (vc *Client) Delete(ctx context.Context, productID int) error {
	return vc.client.Unwrap().Del(ctx, vc.key(productID)).Err()
}

func (vc *Client) key(id int) string {
	return fmt.Sprintf(keyPrefix, id)
}
