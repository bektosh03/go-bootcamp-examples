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

// @Summary      Create author
// @Description  creates a author with provided info
// @Tags         authors
// @Accept       json
// @Produce      json
// @Param        request body CreateAuthorRequest true  "Author info"
// @Success      200 {object} CreateAuthorResponse
// @Failure      400
// @Failure      500
// @Router       /author [POST]
func (s Server) CreateAuthor(c *gin.Context) {
	var request CreateAuthorRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := s.repo.CreateAuthor(c.Request.Context(), entity.NewAuthor(request.Name)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})
}

// @Summary      Create book
// @Description  creates a book with provided info
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        request body CreateBookRequest true  "Book info"
// @Success      200
// @Failure      400
// @Failure      500
// @Router       /book [POST]
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
