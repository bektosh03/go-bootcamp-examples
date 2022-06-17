package store

type Store struct {
	inventory Inventory
}

func New(i Inventory) *Store {
	return &Store{
		inventory: i,
	}
}

func (s Store) Run() {

}
