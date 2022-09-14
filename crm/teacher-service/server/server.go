package server

import (
	"context"
	"errors"
	"teacher-service/domain/subject"
	"teacher-service/domain/teacher"
	"teacher-service/service"

	"github.com/bektosh03/crmcommon/errs"
	"github.com/bektosh03/crmprotos/teacherpb"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
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

// RegisterTeacher registers a new teacher
func (s Server) RegisterTeacher(ctx context.Context, req *teacherpb.RegisterTeacherRequest) (*teacherpb.Teacher, error) {
	tch, err := s.convertRegisterTeacherRequestToDomainTeacher(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	createdTeacher, err := s.service.RegisterTeacher(ctx, tch)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return toProtoTeacher(createdTeacher), nil
}

// CreateSubject creates a new subject
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

// GetTeacher fetches teacher's data by teacher ID
func (s Server) GetTeacher(ctx context.Context, req *teacherpb.GetTeacherRequest) (*teacherpb.Teacher, error) {
	var teacherBy teacher.By
	switch by := req.By.(type) {
	case *teacherpb.GetTeacherRequest_Email:
		teacherBy = teacher.ByEmail{Email: by.Email}
	case *teacherpb.GetTeacherRequest_PhoneNumber:
		teacherBy = teacher.ByPhoneNumber{PhoneNumber: by.PhoneNumber}
	case *teacherpb.GetTeacherRequest_TeacherId:
		id, err := uuid.Parse(by.TeacherId)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "id is not uuid")
		}
		teacherBy = teacher.ByID{ID: id}
	default:
		return nil, status.Error(codes.InvalidArgument, "by is not provided")
	}

	t, err := s.service.GetTeacher(ctx, teacherBy)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return toProtoTeacher(t), nil
}

// GetSubject fetches seubject's data by subject ID
func (s Server) GetSubject(ctx context.Context, req *teacherpb.GetSubjectRequest) (*teacherpb.Subject, error) {
	id, err := uuid.Parse(req.SubjectId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "id is not uuid")
	}

	sub, err := s.service.GetSubject(ctx, id)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return toProtoSubject(sub), nil
}

// DeleteTeacher deletes teacher by ID
func (s Server) DeleteTeacher(ctx context.Context, req *teacherpb.DeleteTeacherRequest) (*emptypb.Empty, error) {
	teacherID, err := uuid.Parse(req.GetTeacherId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "id is not uuid")
	}
	err = s.service.DeleteTeacher(context.Background(), teacherID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

// DeleteSubject deletes subject by ID
func (s Server) DeleteSubject(ctx context.Context, req *teacherpb.DeleteSubjectRequest) (*emptypb.Empty, error) {
	subjectID, err := uuid.Parse(req.GetSubjectId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "id is not uuid")
	}
	err = s.service.DeleteSubject(context.Background(), subjectID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

// ListTeachers returns list of teachers
func (s Server) ListTeachers(ctx context.Context, req *teacherpb.ListTeachersRequest) (*teacherpb.ListTeachersResponse, error) {
	list, _, err := s.service.ListTeachers(ctx, req.Page, req.Limit)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	protTechs := make([]*teacherpb.Teacher, 0, len(list))
	for _, item := range list {
		prt := toProtoTeacher(item)
		protTechs = append(protTechs, prt)
	}

	return &teacherpb.ListTeachersResponse{
		Teachers: protTechs,
	}, nil
}

// ListSubjects returns list of subjects
func (s Server) ListSubjects(ctx context.Context, req *teacherpb.ListSubjectsRequest) (*teacherpb.ListSubjectsResponse, error) {
	list, _, err := s.service.ListSubjects(ctx, req.Page, req.Limit)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	protSubs := make([]*teacherpb.Subject, 0, len(list))
	for _, item := range list {
		prt := toProtoSubject(item)
		protSubs = append(protSubs, prt)
	}

	return &teacherpb.ListSubjectsResponse{
		Subjects: protSubs,
	}, nil
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
		protoTeacher.Password,
		subjectID,
	)
	if err != nil {
		return teacher.Teacher{}, status.Error(codes.InvalidArgument, err.Error())
	}

	return t, nil
}
