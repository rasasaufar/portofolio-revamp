package config

import (
	"fmt"
	"os"
	"time"
)

// Config holds all application configuration loaded from environment variables.
type Config struct {
	ServerPort           string
	DatabaseURL          string
	JWTSecret            string
	JWTExpiry            time.Duration
	CORSOrigin           string
	AdminDefaultPassword string
}

// Load reads configuration from environment variables and validates required fields.
// It exits with a descriptive error if any required variable is missing.
func Load() *Config {
	cfg := &Config{
		ServerPort:           getEnv("SERVER_PORT", "8080"),
		DatabaseURL:          requireEnv("DATABASE_URL"),
		JWTSecret:            requireEnv("JWT_SECRET"),
		JWTExpiry:            parseDuration(getEnv("JWT_EXPIRY", "24h")),
		CORSOrigin:           getEnv("CORS_ORIGIN", "http://localhost:5173"),
		AdminDefaultPassword: getEnv("ADMIN_DEFAULT_PASSWORD", ""),
	}
	return cfg
}

// getEnv returns the value of an environment variable or a default value.
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

// requireEnv returns the value of a required environment variable.
// It panics with a descriptive message if the variable is not set.
func requireEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Sprintf("required environment variable %s is not set", key))
	}
	return value
}

// parseDuration parses a duration string, defaulting to 24h on error.
func parseDuration(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		return 24 * time.Hour
	}
	return d
}
