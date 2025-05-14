package config

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	resetEnv := func() {
		os.Unsetenv("REDIS_ADDR")
		os.Unsetenv("REDIS_PASSWORD")
		os.Unsetenv("REDIS_DB")
		os.Unsetenv("REDIS_TTL_SECONDS")
		os.Unsetenv("MYSQL_USER")
		os.Unsetenv("MYSQL_PASSWORD")
		os.Unsetenv("MYSQL_HOST")
		os.Unsetenv("MYSQL_PORT")
		os.Unsetenv("MYSQL_DATABASE")
		os.Unsetenv("MYSQL_DSN")
		os.Unsetenv("NATS_URL")
		os.Unsetenv("GRPC_PORT")
		os.Unsetenv("MAILJET_API_KEY")
		os.Unsetenv("MAILJET_SECRET_KEY")
		os.Unsetenv("MAILJET_SENDER_EMAIL")
		os.Unsetenv("MAILJET_SENDER_NAME")
	}

	t.Run("Load with default values", func(t *testing.T) {
		resetEnv()

		cfg, err := LoadConfig()

		assert.NoError(t, err)
		assert.NotNil(t, cfg)

		assert.Equal(t, "localhost:6379", cfg.Redis.Addr)
		assert.Equal(t, "", cfg.Redis.Password)
		assert.Equal(t, 0, cfg.Redis.DB)
		assert.Equal(t, 86400*time.Second, cfg.Redis.TTL)

		assert.Equal(t, "root", cfg.MySQL.User)
		assert.Equal(t, "MyStrongPassword123!", cfg.MySQL.Password)
		assert.Equal(t, "127.0.0.1", cfg.MySQL.Host)
		assert.Equal(t, "3306", cfg.MySQL.Port)
		assert.Equal(t, "shop", cfg.MySQL.Database)
		assert.Equal(t, "root:MyStrongPassword123!@tcp(127.0.0.1:3306)/shop?parseTime=true", cfg.MySQL.DSN)

		assert.Equal(t, "nats://localhost:4222", cfg.NATS.URL)

		assert.Equal(t, 50054, cfg.GRPC.Port)

		assert.Equal(t, "", cfg.Mailjet.APIKey)
		assert.Equal(t, "", cfg.Mailjet.SecretKey)
		assert.Equal(t, "adajdzardanov@gmail.com", cfg.Mailjet.FromEmail)
		assert.Equal(t, "CodeMart", cfg.Mailjet.FromName)
	})

	t.Run("Load with custom environment variables", func(t *testing.T) {
		resetEnv()
		os.Setenv("REDIS_ADDR", "redis:6379")
		os.Setenv("REDIS_PASSWORD", "secret")
		os.Setenv("REDIS_DB", "1")
		os.Setenv("REDIS_TTL_SECONDS", "3600")
		os.Setenv("MYSQL_USER", "admin")
		os.Setenv("MYSQL_PASSWORD", "pass123")
		os.Setenv("MYSQL_HOST", "mysql")
		os.Setenv("MYSQL_PORT", "3307")
		os.Setenv("MYSQL_DATABASE", "testdb")
		os.Setenv("MYSQL_DSN", "admin:pass123@tcp(mysql:3307)/testdb?parseTime=true")
		os.Setenv("NATS_URL", "nats://nats:4222")
		os.Setenv("GRPC_PORT", "50055")
		os.Setenv("MAILJET_API_KEY", "key123")
		os.Setenv("MAILJET_SECRET_KEY", "secret456")
		os.Setenv("MAILJET_SENDER_EMAIL", "test@example.com")
		os.Setenv("MAILJET_SENDER_NAME", "TestMart")

		cfg, err := LoadConfig()

		assert.NoError(t, err)
		assert.NotNil(t, cfg)

		assert.Equal(t, "redis:6379", cfg.Redis.Addr)
		assert.Equal(t, "secret", cfg.Redis.Password)
		assert.Equal(t, 1, cfg.Redis.DB)
		assert.Equal(t, 3600*time.Second, cfg.Redis.TTL)

		assert.Equal(t, "admin", cfg.MySQL.User)
		assert.Equal(t, "pass123", cfg.MySQL.Password)
		assert.Equal(t, "mysql", cfg.MySQL.Host)
		assert.Equal(t, "3307", cfg.MySQL.Port)
		assert.Equal(t, "testdb", cfg.MySQL.Database)
		assert.Equal(t, "admin:pass123@tcp(mysql:3307)/testdb?parseTime=true", cfg.MySQL.DSN)

		assert.Equal(t, "nats://nats:4222", cfg.NATS.URL)

		assert.Equal(t, 50055, cfg.GRPC.Port)

		assert.Equal(t, "key123", cfg.Mailjet.APIKey)
		assert.Equal(t, "secret456", cfg.Mailjet.SecretKey)
		assert.Equal(t, "test@example.com", cfg.Mailjet.FromEmail)
		assert.Equal(t, "TestMart", cfg.Mailjet.FromName)
	})

	t.Run("Invalid REDIS_DB", func(t *testing.T) {
		resetEnv()
		os.Setenv("REDIS_DB", "invalid")

		cfg, err := LoadConfig()

		assert.Error(t, err)
		assert.Nil(t, cfg)
		assert.Contains(t, err.Error(), "invalid REDIS_DB")
	})

	t.Run("Invalid REDIS_TTL_SECONDS", func(t *testing.T) {
		resetEnv()
		os.Setenv("REDIS_TTL_SECONDS", "invalid")

		cfg, err := LoadConfig()

		assert.Error(t, err)
		assert.Nil(t, cfg)
		assert.Contains(t, err.Error(), "invalid REDIS_TTL_SECONDS")
	})

	t.Run("Invalid GRPC_PORT", func(t *testing.T) {
		resetEnv()
		os.Setenv("GRPC_PORT", "invalid")

		cfg, err := LoadConfig()

		assert.Error(t, err)
		assert.Nil(t, cfg)
		assert.Contains(t, err.Error(), "invalid GRPC_PORT")
	})
}

func TestGetEnv(t *testing.T) {
	t.Run("Environment variable set", func(t *testing.T) {
		key := "TEST_KEY"
		value := "test_value"
		os.Setenv(key, value)
		defer os.Unsetenv(key)

		result := getEnv(key, "default")

		assert.Equal(t, value, result)
	})

	t.Run("Environment variable unset", func(t *testing.T) {
		key := "UNSET_KEY"
		os.Unsetenv(key)

		result := getEnv(key, "default")

		assert.Equal(t, "default", result)
	})
}
