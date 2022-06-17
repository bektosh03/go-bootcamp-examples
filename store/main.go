package main

import (
	"store/inventory"
	"store/store"
)

func main() {
	i := inventory.NewInMemoryInventory()
	s := store.New(i)
	s.Run()
}
