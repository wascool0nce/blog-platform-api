package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	_ "modernc.org/sqlite"
)

type dbConfig struct {
	Driver string `env:"DB_DRIVER" envDefault:"sqlite"`
	Path   string `env:"DB_PATH" envDefault:"./sqlite3.db"`
}

type HTTPConfig struct {
	Port string `env:"HTTP_PORT" envDefault:":8080"`
}

type Config struct {
	DB   dbConfig
	HTTP HTTPConfig
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}
