package redis

import (
	"context"
	"testing"
	"time"
)

func TestNewClient_Success(t *testing.T) {
	cfg := Config{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}

	client, err := NewClient(cfg)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if client.Conn == nil {
		t.Fatal("expected redis client to be initialized")
	}
	defer client.Close()
}

func TestClient_Ping(t *testing.T) {
	cfg := Config{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}

	client, err := NewClient(cfg)
	if err != nil {
		t.Fatalf("failed to create redis client: %v", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := client.Ping(ctx); err != nil {
		t.Fatalf("expected ping to succeed, got %v", err)
	}
}

func TestClient_Close(t *testing.T) {
	cfg := Config{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}

	client, err := NewClient(cfg)
	if err != nil {
		t.Fatalf("failed to create redis client: %v", err)
	}

	if err := client.Close(); err != nil {
		t.Fatalf("expected close to succeed, got %v", err)
	}
}
