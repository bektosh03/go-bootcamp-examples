package service

import (
	"context"
	"github.com/google/uuid"
	"student-service/domain/group"
	"student-service/domain/student"
)

// Repository defines methods that should be present in our storage provider
type Repository interface {
	StudentRepository
	GroupRepository
}

type StudentRepository interface {
	CreateStudent(ctx context.Context, s student.Student) error
	GetStudent(ctx context.Context, id uuid.UUID) (student.Student, error)
	UpdateStudent(ctx context.Context, s student.Student) error
	DeleteStudent(ctx context.Context, id uuid.UUID) error
	ListStudents(ctx context.Context, page, limit int32) ([]student.Student, int, error)
}

type GroupRepository interface {
	CreateGroup(ctx context.Context, s group.Group) error
	GetGroup(ctx context.Context, id uuid.UUID) (group.Group, error)
	UpdateGroup(ctx context.Context, s group.Group) error
	DeleteGroup(ctx context.Context, id uuid.UUID) error
	ListGroups(ctx context.Context, page, limit int32) ([]group.Group, int, error)
}
