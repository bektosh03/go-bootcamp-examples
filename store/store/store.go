package store

import (
	"fmt"
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

func (s Store) FindProduct(name string)(product.Product,bool)  {
	return s.inventory.FindProduct(name) 
}

func (s Store) Run() {
	p, exists := s.inventory.FindProduct("Olma")
	if !exists {
		fmt.Println("Olma topilmadi")
		return
	}

	fmt.Println(p)

	// err := s.inventory.AddProduct(product.Product{
	// 	Name:          "Banana",
	// 	Quantity:      23,
	// 	Price:         10,
	// 	OriginalPrice: 8,
	// })

	// if err != nil {
	// 	fmt.Println("add product err:", err)
	// }
}
