package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"postgres-gin-crud/config"

	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Connect(cfg config.Config) (*sql.DB, error) {
	db, err := sql.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB,
		),
	)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://postgres/migrations",
		"postgres", driver,
	)
	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return nil, fmt.Errorf("failed to migrate: %v", err)
	}

	return db, nil
}
