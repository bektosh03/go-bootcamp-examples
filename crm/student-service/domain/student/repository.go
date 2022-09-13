package student

import (
	"context"
	"github.com/google/uuid"
)

type Repository interface {
	CreateStudent(ctx context.Context, s Student) error
	GetStudent(ctx context.Context, by By) (Student, error)
	UpdateStudent(ctx context.Context, s Student) error
	DeleteStudent(ctx context.Context, id uuid.UUID) error
	ListStudents(ctx context.Context, page, limit int32) ([]Student, int, error)
	GetStudentsByGroup(ctx context.Context, groupID uuid.UUID) ([]Student, error)
}
