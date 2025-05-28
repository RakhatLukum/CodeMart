package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Redis   RedisConfig
	MySQL   MySQLConfig
	NATS    NATSConfig
	GRPC    GRPCConfig
	Mailjet MailjetConfig
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
	TTL      time.Duration
}

type MySQLConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
	DSN      string
}

type NATSConfig struct {
	URL string
}

type GRPCConfig struct {
	Port int
}

type MailjetConfig struct {
	APIKey    string
	SecretKey string
	FromEmail string
	FromName  string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load .env file: %w", err)
	}

	dbIndex, err := strconv.Atoi(getEnv("REDIS_DB", "0"))
	if err != nil {
		return nil, fmt.Errorf("invalid REDIS_DB: %w", err)
	}
	ttlSec, err := strconv.Atoi(getEnv("REDIS_TTL_SECONDS", "86400"))
	if err != nil {
		return nil, fmt.Errorf("invalid REDIS_TTL_SECONDS: %w", err)
	}
	grpcPort, err := strconv.Atoi(getEnv("GRPC_PORT", "50054"))
	if err != nil {
		return nil, fmt.Errorf("invalid GRPC_PORT: %w", err)
	}

	return &Config{
		Redis: RedisConfig{
			Addr:     getEnv("REDIS_ADDR", "localhost:6379"),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       dbIndex,
			TTL:      time.Duration(ttlSec) * time.Second,
		},
		MySQL: MySQLConfig{
			User:     getEnv("MYSQL_USER", "root"),
			Password: getEnv("MYSQL_PASSWORD", "MyStrongPassword123!"),
			Host:     getEnv("MYSQL_HOST", "127.0.0.1"),
			Port:     getEnv("MYSQL_PORT", "3306"),
			Database: getEnv("MYSQL_DATABASE", "shop"),
			DSN:      getEnv("MYSQL_DSN", "root:MyStrongPassword123!@tcp(127.0.0.1:3306)/shop?parseTime=true"),
		},
		NATS: NATSConfig{
			URL: getEnv("NATS_URL", "nats://localhost:4222"),
		},
		GRPC: GRPCConfig{
			Port: grpcPort,
		},
		Mailjet: MailjetConfig{
			APIKey:    getEnv("MAILJET_API_KEY", ""),
			SecretKey: getEnv("MAILJET_SECRET_KEY", ""),
			FromEmail: getEnv("MAILJET_SENDER_EMAIL", "adajdzardanov@gmail.com"),
			FromName:  getEnv("MAILJET_SENDER_NAME", "CodeMart"),
		},
	}, nil
}

func getEnv(key, defaultValue string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultValue
}
