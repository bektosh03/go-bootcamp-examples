package store

import (
	"context"
	"store/product"
)

type Inventory interface {
	//FindProduct(name string) (product.Product, bool)
	AddProduct(ctx context.Context, p product.Product) error
	ListProducts(ctx context.Context) (product.List, error)
}
