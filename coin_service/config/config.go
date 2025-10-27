package config

import (
	"fmt"
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	"time"
)

type Postgres struct {
	Host               string        `env:"HOST" envDefault:"localhost"`
	Port               int           `env:"PORT" envDefault:"5432"`
	User               string        `env:"USER" envDefault:"postgres"`
	Password           string        `env:"PASSWORD" envDefault:"postgres"`
	Database           string        `env:"DATABASE"`
	SSLMode            string        `env:"SSL_MODE" envDefault:"disable"`
	MaxIdleConnections int           `env:"MAX_IDLE_CONNECTIONS" envDefault:"25"`
	MaxOpenConnections int           `env:"MAX_OPEN_CONNECTIONS" envDefault:"25"`
	ConnectionLifetime time.Duration `env:"CONNECTION_MAX_LIFETIME" envDefault:"5m"`
}

type Config struct {
	HTTPPort string   `env:"HTTP_PORT" envDefault:"7777"`
	Postgres Postgres `envPrefix:"POSTGRES_"`
}

func (p *Postgres) ConnectionURL() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		p.Host, p.Port, p.User, p.Password, p.Database, p.SSLMode)
}

func Load() Config {
	_ = godotenv.Load(".env")

	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		panic(fmt.Errorf("ошибка загрузки конфигурации: %w", err))
	}
	return cfg
}
