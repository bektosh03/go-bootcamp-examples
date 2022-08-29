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

func (s Server) CreateSubject(ctx context.Context, req *teacherpb.CreateSubjectRequest) (*teacherpb.Subject, error) {
	sub, err := s.subjectFactory.NewSubject(req.Name, req.Description)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	sub, err = s.service.CreateSubject(ctx, sub)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return toProtoSubject(sub), nil
}

func (s Server) GetTeacher(ctx context.Context, req *teacherpb.GetTeacherRequest) (*teacherpb.Teacher, error) {
	id, err := uuid.Parse(req.TeacherId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "id is not uuid")
	}

	t, err := s.service.GetTeacher(ctx, id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return toProtoTeacher(t), nil
}

func (s Server) GetSubject(ctx context.Context, req *teacherpb.GetSubjectRequest) (*teacherpb.Subject, error) {
	id, err := uuid.Parse(req.SubjectId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "id is not uuid")
	}

	sub, err := s.service.GetSubject(ctx, id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return toProtoSubject(sub), nil
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
