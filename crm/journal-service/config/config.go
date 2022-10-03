package config

import (
	"github.com/bektosh03/crmcommon/postgres"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"
)

// Config defines configuration values needed for the entire service
type Config struct {
	Host      string `envconfig:"host" required:"true"`
	Port      string `envconfig:"port" required:"true"`
	KafkaHost string `envconfig:"KAFKAHOST" required:"true"`
	KafkaPort string `envconfig:"KAFKAPORT" required:"true"`
	postgres.Config
}

func Load() (Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
