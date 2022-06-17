package product

type Product struct {
	Name  string
	price int
}

func (p Product) Price() int {
	return p.price
}
