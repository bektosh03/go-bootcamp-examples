package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	KafkaHost string `envconfig:"KAFKA_HOST"`
	KafkaPort string `envconfig:"KAFKA_PORT"`
}

func Load() (Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
