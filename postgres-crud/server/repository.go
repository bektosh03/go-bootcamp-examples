package server

import (
	"context"

	"postgres-gin-crud/entity"
)

type Repository interface {
	CreateBook(ctx context.Context, b entity.Book) error
	CreateAuthor(ctx context.Context, a entity.Author) error
	GetBook(ctx context.Context, id string) (entity.Book, error)
	GetAuthor(ctx context.Context, id string) (entity.Author, error)
	ListBooks(ctx context.Context) ([]entity.Book, error)
	ListAuthors(ctx context.Context) ([]entity.Author, error)
	ListBooksByAuthor(ctx context.Context, authorID string) ([]entity.Book, error)
	DeleteBook(ctx context.Context, id string) error
	DeleteAuthor(ctx context.Context, id string) error
}
