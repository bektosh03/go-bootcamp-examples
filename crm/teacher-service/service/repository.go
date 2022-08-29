package service

import (
	"context"
	"teacher-service/domain/subject"
	"teacher-service/domain/teacher"

	"github.com/google/uuid"
)

// Repository defines methods that should be present in our storage provider
type Repository interface {
	TeacherRepository
	SubjectRepository
}

type TeacherRepository interface {
	CreateTeacher(ctx context.Context, t teacher.Teacher) error
	GetTeacher(ctx context.Context, id uuid.UUID) (teacher.Teacher, error)
	UpdateTeacher(ctx context.Context, t teacher.Teacher) error
	DeleteTeacher(ctx context.Context, id uuid.UUID) error
	ListTeachers(ctx context.Context, page, limit int32) ([]teacher.Teacher, int, error)
}

type SubjectRepository interface {
	CreateSubject(ctx context.Context, s subject.Subject) error
	GetSubject(ctx context.Context, id uuid.UUID) (subject.Subject, error)
	UpdateSubject(ctx context.Context, s subject.Subject) error
	DeleteSubject(ctx context.Context, id uuid.UUID) error
	ListSubjects(ctx context.Context, page, limit int32) ([]subject.Subject, int, error)
}
