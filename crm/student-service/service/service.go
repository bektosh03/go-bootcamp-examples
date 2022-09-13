package service

import (
	"context"
	"github.com/google/uuid"
	"student-service/domain/group"
	"student-service/domain/student"
)

func New(
	repo Repository, groupFactory group.Factory, studentFactory student.Factory,
) Service {
	return Service{
		repo:           repo,
		groupFactory:   groupFactory,
		studentFactory: studentFactory,
	}
}

// Service binds all the core logic together providing necessary methods for APIs
type Service struct {
	repo           Repository
	groupFactory   group.Factory
	studentFactory student.Factory
}

func (s Service) ListGroups(ctx context.Context, pages, limit int32) ([]group.Group, int, error) {
	return s.repo.ListGroups(ctx, pages, limit)
}

func (s Service) DeleteGroup(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteGroup(ctx, id)
}

func (s Service) UpdateGroup(ctx context.Context, g group.Group) error {
	return s.repo.UpdateGroup(ctx, g)
}

func (s Service) ListStudents(ctx context.Context, page, limit int32) ([]student.Student, int, error) {
	return s.repo.ListStudents(ctx, page, limit)
}

func (s Service) DeleteStudent(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteStudent(ctx, id)
}

func (s Service) UpdateStudent(ctx context.Context, st student.Student) error {
	return s.repo.UpdateStudent(ctx, st)
}

func (s Service) GetGroup(ctx context.Context, id uuid.UUID) (group.Group, error) {
	return s.repo.GetGroup(ctx, id)
}

func (s Service) GetStudent(ctx context.Context, by student.By) (student.Student, error) {
	return s.repo.GetStudent(ctx, by)
}

func (s Service) RegisterStudent(ctx context.Context, st student.Student) error {
	return s.repo.CreateStudent(ctx, st)
}

func (s Service) CreateGroup(ctx context.Context, g group.Group) error {
	return s.repo.CreateGroup(ctx, g)
}

func (s Service) GetStudentsByGroup(ctx context.Context, groupID uuid.UUID) ([]student.Student, error) {
	return s.repo.GetStudentsByGroup(ctx, groupID)
}
