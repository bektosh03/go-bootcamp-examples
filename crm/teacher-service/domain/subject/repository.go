package subject

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	CreateSubject(ctx context.Context, s Subject) error
	GetSubject(ctx context.Context, id uuid.UUID) (Subject, error)
	UpdateSubject(ctx context.Context, s Subject) error
	DeleteSubject(ctx context.Context, id uuid.UUID) error
	ListSubjects(ctx context.Context, page, limit int32) ([]Subject, int, error)
}
