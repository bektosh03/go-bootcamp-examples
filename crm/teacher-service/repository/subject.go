package repository

import "github.com/google/uuid"

type Subject struct {
	ID          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
}
