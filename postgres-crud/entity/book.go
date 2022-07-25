package entity

import (
	"github.com/google/uuid"
)

type Book struct {
	ID     string
	Name   string
	Author Author
}

func NewBook(title string, author Author) Book {
	return Book{
		ID:     uuid.NewString(),
		Name:   title,
		Author: author,
	}
}
