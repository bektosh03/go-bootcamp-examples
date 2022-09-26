package inventory

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"store/config"
	"store/product"
)

type PostgresInventory struct {
	db *sqlx.DB
}

func NewPostgresInventory(ctx context.Context, pscfg config.PostgresConfig) (*PostgresInventory, error) {
	db, err := sqlx.ConnectContext(ctx, "postgres", fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		pscfg.PostgresHost, pscfg.PostgresPort, pscfg.PostgresUser, pscfg.PostgresPassword, pscfg.PostgresDB,
	))
	if err != nil {
		return nil, err
	}

	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s", pscfg.PostgresMigrationsPath), "postgres", driver)
	if err != nil {
		return nil, err
	}

	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return nil, err
	}

	return &PostgresInventory{
		db: db,
	}, nil
}

func (p *PostgresInventory) DeleteProduct(ctx context.Context, id string) error {
	query := `
	DELETE FROM products WHERE id=$1
	`

	if _, err := p.db.ExecContext(ctx, query, id); err != nil {
		return err
	}
	return nil
}

func (p *PostgresInventory) AddProduct(ctx context.Context, product product.Product) error {
	query := `
	INSERT INTO products (id, name, quantity, price, original_price)
	VALUES ($1, $2, $3, $4, $5)
	`

	if _, err := p.db.ExecContext(ctx, query, product.ID, product.Name, product.Quantity, product.Price, product.OriginalPrice); err != nil {
		return err
	}

	return nil
}

func (p *PostgresInventory) ListProducts(ctx context.Context) (product.List, error) {
	query := `
	SELECT * FROM products
	`
	list := make(product.List, 0)
	if err := p.db.SelectContext(ctx, &list, query); err != nil {
		return nil, err
	}

	return list, nil
}
