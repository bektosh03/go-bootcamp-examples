package main

import (
	"store/inventory"
	"store/store"
)

func main() {
	i, err := inventory.NewFileInventory("data/inventory.txt")
	if err != nil {
		panic(err)
	}
	defer i.Close()

	s := store.New(i)
	s.Run()
}
