package config

import (
	"os"
)

// Config holds application configuration
type Config struct {
	Port      string
	DBPath    string
	AppName   string
	DebugMode bool
}

// NewConfig creates a new configuration instance
func NewConfig() *Config {
	return &Config{
		Port:      getEnv("PORT", "3000"),
		DBPath:    getEnv("DB_PATH", "users.db"),
		AppName:   getEnv("APP_NAME", "KBTG AI Backend Workshop"),
		DebugMode: getEnv("DEBUG", "false") == "true",
	}
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
