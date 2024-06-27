package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	HTTPListenPort string `env:"HTTP_LISTEN_PORT" env-default:"1234"`
}

func FromEnv() (Config, error) {
	var cfg Config
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return Config{}, fmt.Errorf("cannot parse envs: %w", err)
	}
	return cfg, nil
}
