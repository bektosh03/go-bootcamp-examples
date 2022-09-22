package store

import (
	"context"
	"github.com/google/uuid"
	"store/product"
)

type Store struct {
	inventory Inventory
}

func New(i Inventory) *Store {
	return &Store{
		inventory: i,
	}
}

func (s Store) AddProduct(ctx context.Context, p product.Product) (product.Product, error) {
	p.ID = uuid.NewString()
	if err := s.inventory.AddProduct(ctx, p); err != nil {
		return product.Product{}, err
	}

	return p, nil
}

func (s Store) ListProducts(ctx context.Context) (product.List, error) {
	return s.inventory.ListProducts(ctx)
}
