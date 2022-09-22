package http

type AddProductRequest struct {
	Name          string  `json:"name"`
	Quantity      uint64  `json:"quantity"`
	Price         float32 `json:"price"`
	OriginalPrice float32 `json:"originalPrice"`
}
