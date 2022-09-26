package product

type Product struct {
	ID            string  `db:"id"`
	Name          string  `db:"name"`
	Quantity      uint64  `db:"quantity"`
	Price         float32 `db:"price"`
	OriginalPrice float32 `db:"original_price"`
}

type List []Product

func (l *List) Error() string {
	//TODO implement me
	panic("implement me")
}

func (l *List) Add(p Product) {
	*l = append(*l, p)
}

func (l *List) Remove(name string) {
	index, ok := l.search(name)
	if !ok {
		return
	}

	*l = append((*l)[:index], (*l)[index+1:]...)
}

func (l *List) search(name string) (int, bool) {
	for i, p := range *l {
		if p.Name == name {
			return i, true
		}
	}

	return -1, false
}
