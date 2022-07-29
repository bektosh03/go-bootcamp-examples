package entity

import (
	"github.com/google/uuid"
)

type Book struct {
	ID     string
	Title  string
	Author Author
}

func NewBook(title string, author Author) Book {
	return Book{
		ID:     uuid.NewString(),
		Title:  title,
		Author: author,
	}
}
