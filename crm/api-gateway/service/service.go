package service

import (
	"api-gateway/request"
	"api-gateway/response"
	"context"
)

func New(teacherService TeacherServiceClient) Service {
	return Service{
		Teacher: teacherService,
	}
}

type Service struct {
	Teacher TeacherServiceClient
}

type TeacherServiceClient interface {
	RegisterTeacher(context.Context, request.RegisterTeacherRequest) (response.Teacher, error)
	GetTeacher(context.Context, string) (response.Teacher, error)
	CreateSubject(context.Context, request.CreateSubjectRequest) (response.Subject, error)
	GetSubject(context.Context, string) (response.Subject, error)
}
