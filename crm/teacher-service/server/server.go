package server

import (
	"context"
	"teacher-service/domain/subject"
	"teacher-service/domain/teacher"
	"teacher-service/service"

	"github.com/bektosh03/crmprotos/teacherpb"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func New(svc service.Service, subjectFactory subject.Factory, teacherFactory teacher.Factory) Server {
	return Server{
		service:        svc,
		subjectFactory: subjectFactory,
		teacherFactory: teacherFactory,
	}
}

type Server struct {
	teacherpb.UnimplementedTeacherServiceServer
	service        service.Service
	subjectFactory subject.Factory
	teacherFactory teacher.Factory
}

func (s Server) RegisterTeacher(ctx context.Context, req *teacherpb.RegisterTeacherRequest) (*teacherpb.Teacher, error) {
	tch, err := s.convertRegisterTeacherRequestToDomainTeacher(req)
	if err != nil {
		return nil, err
	}

	createdTeacher, err := s.service.RegisterTeacher(ctx, tch)
	if err != nil {
		return nil, err
	}

	return toProtoTeacher(createdTeacher), nil
}

func (s Server) convertRegisterTeacherRequestToDomainTeacher(protoTeacher *teacherpb.RegisterTeacherRequest) (teacher.Teacher, error) {
	subjectID, err := uuid.Parse(protoTeacher.SubjectId)
	if err != nil {
		return teacher.Teacher{}, status.Error(codes.InvalidArgument, "provided subject id is not uuid")
	}

	t, err := s.teacherFactory.NewTeacher(
		protoTeacher.FirstName,
		protoTeacher.LastName,
		protoTeacher.Email,
		protoTeacher.PhoneNumber,
		subjectID,
	)
	if err != nil {
		return teacher.Teacher{}, status.Error(codes.InvalidArgument, err.Error())
	}

	return t, nil
}
