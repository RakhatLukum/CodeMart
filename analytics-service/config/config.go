package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Redis RedisConfig
	MySQL MySQLConfig
	NATS  NATSConfig
	GRPC  GRPCConfig
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
	Port string
}

func LoadConfig() *Config {
	_ = godotenv.Load()

	dbIndex, _ := strconv.Atoi(getEnv("REDIS_DB", "0"))
	ttlSec, _ := strconv.Atoi(getEnv("REDIS_TTL_SECONDS", "86400"))

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
			Port: getEnv("GRPC_PORT", "50054"),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultValue
}
