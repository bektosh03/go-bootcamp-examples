package teacher

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	CreateTeacher(ctx context.Context, t Teacher) error
	GetTeacher(ctx context.Context, by By) (Teacher, error)
	UpdateTeacher(ctx context.Context, t Teacher) error
	DeleteTeacher(ctx context.Context, id uuid.UUID) error
	ListTeachers(ctx context.Context, page, limit int32) ([]Teacher, int, error)
}
