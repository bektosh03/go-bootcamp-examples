package product

type Product struct {
	Name     string
	price    int
	quantity int
}

func (p Product) Price() int {
	return p.price
}
