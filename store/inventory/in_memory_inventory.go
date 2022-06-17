package inventory

import "store/product"

type InMemoryInventory struct {
	products []product.Product
}

func NewInMemoryInventory() *InMemoryInventory {
	return &InMemoryInventory{
		products: make([]product.Product, 0),
	}
}

func (i *InMemoryInventory) FindProduct(name string) (product.Product, bool) {
	return product.Product{}, false
}
