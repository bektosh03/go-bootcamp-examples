package server

import (
	"context"
	"github.com/bektosh03/crmprotos/studentpb"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"student-service/domain/group"
	"student-service/domain/student"
	"student-service/service"
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

func (s Server) ListGroups(ctx context.Context, req *studentpb.ListGroupsRequest) (*studentpb.GroupList, error) {
	groups, _, err := s.service.ListGroups(ctx, req.GetPage(), req.GetLimit())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return toProtoGroups(groups), nil
}

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

func (s Server) ListStudents(ctx context.Context, req *studentpb.ListStudentsRequest) (*studentpb.StudentList, error) {
	students, _, err := s.service.ListStudents(ctx, req.GetPage(), req.GetLimit())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return toProtoStudents(students), nil
}

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

func (s Server) GetGroup(ctx context.Context, request *studentpb.GetGroupRequest) (*studentpb.Group, error) {
	id, err := uuid.Parse(request.GetGroupId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "id is not uuid")
	}

	gr, err := s.service.GetGroup(ctx, id)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return toProtoGroup(gr), nil
}

func (s Server) GetStudent(ctx context.Context, req *studentpb.GetStudentRequest) (*studentpb.Student, error) {
	id, err := uuid.Parse(req.GetStudentId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "id is not uuid")
	}
	st, err := s.service.GetStudent(ctx, id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return toProtoStudent(st), nil
}

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

func (s Server) RegisterStudent(ctx context.Context, req *studentpb.RegisterStudentRequest) (*studentpb.Student, error) {
	st, err := s.convertRegisterStudentRequestToDomainStudent(req)
	if err != nil {
		return nil, err
	}

	if err = s.service.RegisterStudent(ctx, st); err != nil {
		return nil, err
	}

	return toProtoStudent(st), nil
}

func (s Server) CreateGroup(ctx context.Context, req *studentpb.CreateGroupRequest) (*studentpb.Group, error) {
	mainTeacherId, err := uuid.Parse(req.MainTeacherId)
	if err != nil {
		return nil, err
	}

	gr, err := s.groupFactory.NewGroup(req.Name, mainTeacherId)

	if err = s.service.CreateGroup(ctx, gr); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return toProtoGroup(gr), nil
}

func (s Server) convertUpdateGroupRequestToDomainStudent(protoGroup *studentpb.Group) (group.Group, error) {
	id, err := uuid.Parse(protoGroup.GetId())
	if err != nil {
		return group.Group{}, err
	}

	mainTeacherID, err := uuid.Parse(protoGroup.GetMainTeacherId())
	if err != nil {
		return group.Group{}, err
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
		return student.Student{}, err
	}

	groupId, err := uuid.Parse(protoStudent.GetGroupId())
	if err != nil {
		return student.Student{}, err
	}

	st, err := student.UnmarshalStudent(student.UnmarshalStudentArgs{
		ID:          id,
		FirstName:   protoStudent.FirstName,
		LastName:    protoStudent.LastName,
		Email:       protoStudent.Email,
		PhoneNumber: protoStudent.PhoneNumber,
		Level:       protoStudent.Level,
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
		return student.Student{}, status.Error(codes.InvalidArgument, "provided group uuid is not uuid")
	}

	st, err := s.studentFactory.NewStudent(
		protoStudent.FirstName,
		protoStudent.LastName,
		protoStudent.Email,
		protoStudent.PhoneNumber,
		protoStudent.Level,
		groupID,
	)
	if err != nil {
		return student.Student{}, status.Error(codes.InvalidArgument, err.Error())
	}

	return st, nil
}
