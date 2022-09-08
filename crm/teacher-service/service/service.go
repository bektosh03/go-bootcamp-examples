package service

import (
	"context"
	"teacher-service/domain/subject"
	"teacher-service/domain/teacher"

	"github.com/google/uuid"
)

func New(
	repo Repository, subjectFactory subject.Factory, teacherFactory teacher.Factory,
) Service {
	return Service{
		repo:           repo,
		subjectFactory: subjectFactory,
		teacherFactory: teacherFactory,
	}
}

// Service binds all the core logic together providing necessary methods for APIs
type Service struct {
	repo           Repository
	subjectFactory subject.Factory
	teacherFactory teacher.Factory
}

func (s Service) RegisterTeacher(ctx context.Context, t teacher.Teacher) (teacher.Teacher, error) {
	if err := s.repo.CreateTeacher(ctx, t); err != nil {
		return teacher.Teacher{}, err
	}

	return t, nil
}

func (s Service) CreateSubject(ctx context.Context, sub subject.Subject) (subject.Subject, error) {
	if err := s.repo.CreateSubject(ctx, sub); err != nil {
		return subject.Subject{}, err
	}

	return sub, nil
}

func (s Service) UpdateTeacher(ctx context.Context, t teacher.Teacher) (teacher.Teacher, error) {
	if err := s.repo.UpdateTeacher(ctx, t); err != nil {
		return teacher.Teacher{}, err
	}

	return t, nil
}

func (s Service) UpdateSubject(ctx context.Context, sub subject.Subject) (subject.Subject, error) {
	if err := s.repo.UpdateSubject(ctx, sub); err != nil {
		return subject.Subject{}, err
	}

	return sub, nil
}

func (s Service) GetTeacher(ctx context.Context, by teacher.By) (teacher.Teacher, error) {
	t, err := s.repo.GetTeacher(ctx, by)
	if err != nil {
		return teacher.Teacher{}, err
	}

	return t, nil
}

func (s Service) GetSubject(ctx context.Context, id uuid.UUID) (subject.Subject, error) {
	sub, err := s.repo.GetSubject(ctx, id)
	if err != nil {
		return subject.Subject{}, err
	}

	return sub, nil
}

func (s Service) DeleteTeacher(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteTeacher(ctx, id)
}

func (s Service) DeleteSubject(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteSubject(ctx, id)
}

func (s Service) ListTeachers(ctx context.Context, page, limit int32) ([]teacher.Teacher, int, error) {
	teachers, count, err := s.repo.ListTeachers(ctx, page, limit)
	if err != nil {
		return nil, 0, err
	}

	return teachers, count, nil
}

func (s Service) ListSubjects(ctx context.Context, page, limit int32) ([]subject.Subject, int, error) {
	subjects, count, err := s.repo.ListSubjects(ctx, page, limit)
	if err != nil {
		return nil, 0, err
	}

	return subjects, count, nil
}
