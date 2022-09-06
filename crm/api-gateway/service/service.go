package service

import (
	"api-gateway/request"
	"api-gateway/response"
	"context"
)

func New(teacherService TeacherServiceClient, studentService StudentServiceClient) Service {
	return Service{
		Teacher: teacherService,
		Student: studentService,
	}
}

type Service struct {
	Teacher TeacherServiceClient
	Student StudentServiceClient
}

type StudentServiceClient interface {
	RegisterStudent(context.Context, request.RegisterStudentRequest) (response.Student, error)
	GetStudent(context.Context, string) (response.Student, error)
	CreateGroup(context.Context, request.CreateGroupRequest) (response.Group, error)
	GetGroup(context.Context, string) (response.Group, error)
}

type TeacherServiceClient interface {
	RegisterTeacher(context.Context, request.RegisterTeacherRequest) (response.Teacher, error)
	GetTeacher(context.Context, string) (response.Teacher, error)
	CreateSubject(context.Context, request.CreateSubjectRequest) (response.Subject, error)
	GetSubject(context.Context, string) (response.Subject, error)
}
