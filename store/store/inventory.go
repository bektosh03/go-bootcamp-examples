package store

import "store/product"

type Inventory interface {
	FindProduct(name string) (product.Product, bool)
	AddProduct(p product.Product) error
}
