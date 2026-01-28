package main

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	DatabaseName     string `env:"DB_NAME"`
	DatabaseHost     string `env:"DB_HOST"`
	DatabasePort     string `env:"DB_PORT"`
	DatabaseUser     string `env:"DB_USER"`
	DatabasePassword string `env:"DB_PASSWORD"`
}

func (c *Config) DatabaseUrl() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disbale",
		c.DatabaseHost,
		c.DatabaseName,
		c.DatabasePort,
		c.DatabaseUser,
		c.DatabasePassword,
	)
}

func New() (*Config, error) {
	var cfg Config
	cfg, err := env.ParseAs[Config]()
	if err != nil {
		return nil, fmt.Errorf("Failed to load config: %w", err)
	}
	return &cfg, nil

}
