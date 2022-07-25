package server

import (
	"net/http"
	"postgres-gin-crud/entity"

	"github.com/gin-gonic/gin"
)

type Server struct {
	repo Repository
}

func New(repo Repository) Server {
	return Server{
		repo: repo,
	}
}

func (s Server) CreateBook(c *gin.Context) {
	var request CreateBookRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	b := entity.NewBook(request.Title, entity.Author{
		ID: request.AuthorID,
	})
	if err := s.repo.CreateBook(c.Request.Context(), b); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})
}
