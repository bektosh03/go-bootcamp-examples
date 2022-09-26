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

func (s Store) DeleteProduct(ctx context.Context, id string) error {
	if err := s.inventory.DeleteProduct(ctx, id); err != nil {
		return err
	}
	return nil
}

func (s Store) AddProduct(ctx context.Context, p product.Product) (product.Product, error) {
	p.ID = uuid.NewString()
	if err := s.inventory.AddProduct(ctx, p); err != nil {
		return product.Product{}, err
	}

	return p, nil
}

func (s Store) ListProducts(ctx context.Context) (product.List, error) {
	if _, err := s.inventory.ListProducts(ctx); err != nil {
		return nil, err
	}
	return nil, nil
}

//
//func (s Store) FindProduct(ctx context.Context,name string) (product.Product,bool) {
//	return product.Product{},true
//}
//
//func (s Store) GetProduct(ctx context.Context, id string) (product.Product,error) {
//	return product.Product{},nil
//}
//
//func (s Store) UpdateProduct(ctx context.Context,p product.Product) (error) {
//	return nil
//}
