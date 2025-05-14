package mysql

import (
	"testing"
)

func TestNewDB_WithValidDSN(t *testing.T) {
	cfg := Config{
		DSN: "root:MyStrongPassword123!@tcp(127.0.0.1:3306)/shop?parseTime=true",
	}

	db, err := NewDB(cfg)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	defer db.Close()

	if db.Conn == nil {
		t.Fatal("expected db.Conn to be initialized")
	}
}

func TestNewDB_WithInvalidDSN(t *testing.T) {
	cfg := Config{
		DSN: "invalid:dsn@tcp(127.0.0.1:9999)/wrong",
	}

	_, err := NewDB(cfg)
	if err == nil {
		t.Fatal("expected an error, got nil")
	}
}

func TestDB_Close(t *testing.T) {
	cfg := Config{
		DSN: "root:MyStrongPassword123!@tcp(127.0.0.1:3306)/shop?parseTime=true",
	}

	db, err := NewDB(cfg)
	if err != nil {
		t.Fatalf("setup failed: %v", err)
	}

	err = db.Close()
	if err != nil {
		t.Fatalf("expected close to succeed, got %v", err)
	}
}
