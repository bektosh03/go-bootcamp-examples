package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"store/product"
	"store/store"
)

func NewServer(s *store.Store) Server {
	return Server{
		store: s,
	}
}

type Server struct {
	store *store.Store
}

func (s Server) GetProduct(c *gin.Context) {

}

func (s Server) UpdateProduct(c *gin.Context) {

}

func (s Server) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if err := s.store.DeleteProduct(context.Background(), id); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"okay": true})
}

func (s Server) AddProduct(c *gin.Context) {
	var request AddProductRequest
	if err := c.BindJSON(&request); err != nil {
		return
	}

	p, err := s.store.AddProduct(c.Request.Context(), product.Product{
		Name:          request.Name,
		Quantity:      request.Quantity,
		Price:         request.Price,
		OriginalPrice: request.OriginalPrice,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, p)
}

func (s Server) Run(addr string) error {
	r := gin.Default()

	r.POST("/product", s.AddProduct)
	r.DELETE("/delete_product/:id", s.DeleteProduct)

	return r.Run(addr)
}
