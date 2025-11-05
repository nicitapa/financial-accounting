package config

import "os"

type Config struct {
	Port                  string
	TransactionServiceURL string
	AuthServiceURL        string
}

func Load() *Config {
	return &Config{
		Port:                  getEnv("PORT", ":8284"),
		TransactionServiceURL: getEnv("TRANSACTION_SERVICE_URL", "http://localhost:8080"),
		AuthServiceURL:        getEnv("AUTH_SERVICE_URL", "http://localhost:8090"),
	}
}
func getEnv(key, def string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return def
}
