package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig_DefaultValues(t *testing.T) {
	// Arrange
	// Clear environment variables
	os.Unsetenv("PORT")
	os.Unsetenv("DB_PATH")
	os.Unsetenv("APP_NAME")
	os.Unsetenv("DEBUG")

	// Act
	cfg := NewConfig()

	// Assert
	assert.Equal(t, "3000", cfg.Port)
	assert.Equal(t, "users.db", cfg.DBPath)
	assert.Equal(t, "KBTG AI Backend Workshop", cfg.AppName)
	assert.False(t, cfg.DebugMode)
}

func TestNewConfig_CustomValues(t *testing.T) {
	// Arrange
	os.Setenv("PORT", "8080")
	os.Setenv("DB_PATH", "custom.db")
	os.Setenv("APP_NAME", "Custom App")
	os.Setenv("DEBUG", "true")

	defer func() {
		// Cleanup
		os.Unsetenv("PORT")
		os.Unsetenv("DB_PATH")
		os.Unsetenv("APP_NAME")
		os.Unsetenv("DEBUG")
	}()

	// Act
	cfg := NewConfig()

	// Assert
	assert.Equal(t, "8080", cfg.Port)
	assert.Equal(t, "custom.db", cfg.DBPath)
	assert.Equal(t, "Custom App", cfg.AppName)
	assert.True(t, cfg.DebugMode)
}

func TestGetEnv(t *testing.T) {
	// Test with existing environment variable
	os.Setenv("TEST_VAR", "test_value")
	result := getEnv("TEST_VAR", "default")
	assert.Equal(t, "test_value", result)

	// Test with non-existing environment variable
	result = getEnv("NON_EXISTING_VAR", "default")
	assert.Equal(t, "default", result)

	// Cleanup
	os.Unsetenv("TEST_VAR")
}
