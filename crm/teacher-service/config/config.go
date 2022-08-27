package config

// Config defines configuration values needed for the entire service
type Config struct {
	Host string
	Port string
	PostgresConfig
}

// PostgresConfig defines variables needed for postgres
type PostgresConfig struct {
	PostgresHost           string
	PostgresPort           string
	PostgresUser           string
	PostgresPassword       string
	PostgresDB             string
	PostgresMigrationsPath string
}
