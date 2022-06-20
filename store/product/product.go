package product

type Product struct {
	Name          string
	Quantity      uint64
	Price         uint64
	OriginalPrice uint64
}

type List []Product
