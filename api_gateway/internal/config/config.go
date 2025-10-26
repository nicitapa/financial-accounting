package config

import (
	"os"
)

type Config struct {
	Port                  string
	TransactionServiceURL string
}

func Load() *Config {
	return &Config{
		Port:                  getEnv("PORT", ":8283"),
		TransactionServiceURL: getEnv("TRANSACTION_SERVICE_URL", "http://localhost:8080"),
	}
}

func getEnv(key, def string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return def
}
