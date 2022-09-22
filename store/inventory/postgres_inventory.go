package inventory

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"store/product"
)

const (
	PostgresHost     = "localhost"
	PostgresPort     = 5432
	PostgresUser     = "postgres"
	PostgresPassword = "1234"
	PostgresDBName   = "storedb"
)

type PostgresInventory struct {
	db *sqlx.DB
}

func NewPostgresInventory(ctx context.Context) (*PostgresInventory, error) {
	db, err := sqlx.ConnectContext(ctx, "postgres", fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		PostgresHost, PostgresPort, PostgresUser, PostgresPassword, PostgresDBName,
	))
	if err != nil {
		return nil, err
	}

	return &PostgresInventory{
		db: db,
	}, nil
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
