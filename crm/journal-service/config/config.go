package config

import (
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"
)

// Config defines configuration values needed for the entire service
type Config struct {
	Host string `envconfig:"host" required:"true"`
	Port string `envconfig:"port" required:"true"`
	postgres.Config
}

// PostgresConfig defines variables needed for postgres
type PostgresConfig struct {
	PostgresHost           string `envconfig:"POSTGRES_HOST" required:"true"`
	PostgresPort           string `envconfig:"POSTGRES_PORT" required:"true"`
	PostgresUser           string `envconfig:"POSTGRES_USER" required:"true"`
	PostgresPassword       string `envconfig:"POSTGRES_PASSWORD" required:"true"`
	PostgresDB             string `envconfig:"POSTGRES_DB" required:"true"`
	PostgresMigrationsPath string `envconfig:"POSTGRES_MIGRATIONS_PATH" required:"true"`
}

func Load() (Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
