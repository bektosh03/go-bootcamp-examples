package postgres

import (
	"context"
	"database/sql"
	"postgres-gin-crud/entity"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) CreateAuthor(ctx context.Context, a entity.Author) error {
	query := `
	INSERT INTO authors VALUES ($1, $2)
	`
	_, err := r.db.ExecContext(ctx, query, a.ID, a.Name)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) CreateBook(ctx context.Context, b entity.Book) error {
	query := `
	INSERT INTO books VALUES ($1, $2, $3)
	`
	_, err := r.db.ExecContext(ctx, query, b.ID, b.Title, b.Author.ID)
	if err != nil {
		return err
	}

	return nil
}
