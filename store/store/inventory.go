package store

import (
	"context"
	"store/product"
)

type Inventory interface {
	//FindProduct(name string) (product.Product, bool)
	AddProduct(ctx context.Context, p product.Product) error
	ListProducts(ctx context.Context) (product.List, error)
	//GetProduct(ctx context.Context,id string)(product.Product,error)
	//UpdateProduct(ctx context.Context,p product.Product)(error)
	DeleteProduct(ctx context.Context, id string) error
}
