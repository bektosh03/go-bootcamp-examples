package store

import "fmt"

type Store struct {
	inventory Inventory
}

func New(i Inventory) *Store {
	return &Store{
		inventory: i,
	}
}

func (s Store) Run() {
	p, exists := s.inventory.FindProduct("Olma")
	if !exists {
		fmt.Println("Olma topilmadi")
		return
	}

	fmt.Println(p)
}
