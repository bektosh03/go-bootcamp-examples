package entity

import "github.com/google/uuid"

type Author struct {
	ID   string
	Name string
}

func NewAuthor(name string) Author {
	return Author{
		ID:   uuid.NewString(),
		Name: name,
	}
}
