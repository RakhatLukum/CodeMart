package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
