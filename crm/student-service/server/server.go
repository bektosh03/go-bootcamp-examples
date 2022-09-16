package server

import (
	"context"
	"errors"
	"student-service/domain/group"
	"student-service/domain/student"
	"student-service/service"

	"github.com/bektosh03/crmcommon/errs"
	"github.com/bektosh03/crmprotos/studentpb"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func New(svc service.Service, groupFactory group.Factory, studentFactory student.Factory) Server {
	return Server{
		service:        svc,
		groupFactory:   groupFactory,
		studentFactory: studentFactory,
	}
}

type Server struct {
	studentpb.UnimplementedStudentServiceServer
	service        service.Service
	groupFactory   group.Factory
	studentFactory student.Factory
}

func (s Server) GetGroupStudents(ctx context.Context, req *studentpb.GetGroupStudentsRequest) (*studentpb.StudentList, error) {
	groupId, err := uuid.Parse(req.GetGroupId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "group id is not uuid")
	}

	students, err := s.service.GetStudentsByGroup(ctx, groupId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return toProtoStudents(students), nil
}

// ListGroups fetches list of groups
func (s Server) ListGroups(ctx context.Context, req *studentpb.ListGroupsRequest) (*studentpb.GroupList, error) {
	groups, _, err := s.service.ListGroups(ctx, req.GetPage(), req.GetLimit())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return toProtoGroups(groups), nil
}

// DeleteGroup deletes group by ID
func (s Server) DeleteGroup(ctx context.Context, req *studentpb.DeleteGroupRequest) (*emptypb.Empty, error) {
	groupId, err := uuid.Parse(req.GetGroupId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "group id is not uuid")
	}

	if err = s.service.DeleteGroup(ctx, groupId); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

// UpdateGroup updates group data
func (s Server) UpdateGroup(ctx context.Context, req *studentpb.Group) (*studentpb.Group, error) {
	gr, err := s.convertUpdateGroupRequestToDomainStudent(req)
	if err != nil {
		return nil, err
	}

	if err = s.service.UpdateGroup(ctx, gr); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return toProtoGroup(gr), nil
}

// ListStudents fetches a list of students from database
func (s Server) ListStudents(ctx context.Context, req *studentpb.ListStudentsRequest) (*studentpb.StudentList, error) {
	students, _, err := s.service.ListStudents(ctx, req.GetPage(), req.GetLimit())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return toProtoStudents(students), nil
}

// DeleteStudent deletes student by ID
func (s Server) DeleteStudent(ctx context.Context, req *studentpb.DeleteStudentRequest) (*emptypb.Empty, error) {
	studentId, err := uuid.Parse(req.GetStudentId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "student id is not uuid")
	}

	if err = s.service.DeleteStudent(ctx, studentId); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

// GetGroup fetches Group data from database by groupID
func (s Server) GetGroup(ctx context.Context, request *studentpb.GetGroupRequest) (*studentpb.Group, error) {
	id, err := uuid.Parse(request.GetGroupId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "id is not uuid")
	}

	gr, err := s.service.GetGroup(ctx, id)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "group is not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return toProtoGroup(gr), nil
}

// GetStudent fetches student data from database by studentID
func (s Server) GetStudent(ctx context.Context, req *studentpb.GetStudentRequest) (*studentpb.Student, error) {
	var studentBy student.By
	switch by := req.By.(type) {
	case *studentpb.GetStudentRequest_Email:
		studentBy = student.ByEmail{Email: by.Email}
	case *studentpb.GetStudentRequest_PhoneNumber:
		studentBy = student.ByPhoneNumber{PhoneNumber: by.PhoneNumber}
	case *studentpb.GetStudentRequest_StudentId:
		id, err := uuid.Parse(req.GetStudentId())
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "id is not uuid")
		}
		studentBy = student.ByID{ID: id}
	default:
		return nil, status.Error(codes.InvalidArgument, "by is not provided")
	}

	st, err := s.service.GetStudent(ctx, studentBy)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "student is not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return toProtoStudent(st), nil
}

// UpdateStudent updates student's info
func (s Server) UpdateStudent(ctx context.Context, req *studentpb.Student) (*studentpb.Student, error) {
	st, err := s.convertUpdateStudentRequestToDomainStudent(req)

	if err != nil {
		return nil, err
	}

	if err = s.service.UpdateStudent(ctx, st); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return toProtoStudent(st), nil
}

// RegisterStudent creates a new student
func (s Server) RegisterStudent(ctx context.Context, req *studentpb.RegisterStudentRequest) (*studentpb.Student, error) {
	st, err := s.convertRegisterStudentRequestToDomainStudent(req)
	if err != nil {
		return nil, err
	}

	if err = s.service.RegisterStudent(ctx, st); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return toProtoStudent(st), nil
}

// CreateGroup creates a new group
func (s Server) CreateGroup(ctx context.Context, req *studentpb.CreateGroupRequest) (*studentpb.Group, error) {
	mainTeacherId, err := uuid.Parse(req.MainTeacherId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "teacher id is not uuid")
	}

	gr, err := s.groupFactory.NewGroup(req.Name, mainTeacherId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err = s.service.CreateGroup(ctx, gr); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return toProtoGroup(gr), nil
}

func (s Server) convertUpdateGroupRequestToDomainStudent(protoGroup *studentpb.Group) (group.Group, error) {
	id, err := uuid.Parse(protoGroup.GetId())
	if err != nil {
		return group.Group{}, status.Error(codes.InvalidArgument, "student id is not uuid")
	}

	mainTeacherID, err := uuid.Parse(protoGroup.GetMainTeacherId())
	if err != nil {
		return group.Group{}, status.Error(codes.InvalidArgument, "teacher id is not uuid")
	}

	gr, err := group.UnmarshalGroup(group.UnmarshalGroupArgs{
		ID:            id,
		Name:          protoGroup.Name,
		MainTeacherID: mainTeacherID,
	})
	if err != nil {
		return group.Group{}, status.Error(codes.InvalidArgument, err.Error())
	}

	return gr, nil
}

func (s Server) convertUpdateStudentRequestToDomainStudent(protoStudent *studentpb.Student) (student.Student, error) {
	id, err := uuid.Parse(protoStudent.GetId())
	if err != nil {
		return student.Student{}, status.Error(codes.InvalidArgument, "student id is not uuid")
	}

	groupId, err := uuid.Parse(protoStudent.GetGroupId())
	if err != nil {
		return student.Student{}, status.Error(codes.InvalidArgument, "group id is not uuid")
	}

	st, err := student.UnmarshalStudent(student.UnmarshalStudentArgs{
		ID:          id,
		FirstName:   protoStudent.FirstName,
		LastName:    protoStudent.LastName,
		Email:       protoStudent.Email,
		PhoneNumber: protoStudent.PhoneNumber,
		Level:       protoStudent.Level,
		Password:    protoStudent.Password,
		GroupID:     groupId,
	})
	if err != nil {
		return student.Student{}, status.Error(codes.InvalidArgument, err.Error())
	}

	return st, nil
}

func (s Server) convertRegisterStudentRequestToDomainStudent(protoStudent *studentpb.RegisterStudentRequest) (student.Student, error) {
	groupID, err := uuid.Parse(protoStudent.GroupId)
	if err != nil {
		return student.Student{}, status.Error(codes.InvalidArgument, "provided group id is not uuid")
	}

	st, err := s.studentFactory.NewStudent(
		protoStudent.FirstName,
		protoStudent.LastName,
		protoStudent.Email,
		protoStudent.PhoneNumber,
		protoStudent.Password,
		protoStudent.Level,
		groupID,
	)
	if err != nil {
		return student.Student{}, status.Error(codes.InvalidArgument, err.Error())
	}

	return st, nil
}
