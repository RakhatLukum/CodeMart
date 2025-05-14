package nats

import (
	"testing"

	"github.com/nats-io/nats.go"
)

func TestNewClient_Success(t *testing.T) {
	client, err := NewClient(nats.DefaultURL)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if client.Conn == nil {
		t.Fatal("expected client.Conn to be initialized")
	}
	client.Close()
}

func TestConnect_Success(t *testing.T) {
	client, err := Connect(nats.DefaultURL)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if client.Conn == nil {
		t.Fatal("expected client.Conn to be initialized")
	}
	client.Close()
}

func TestClose(t *testing.T) {
	client, err := NewClient(nats.DefaultURL)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	client.Close()
}
