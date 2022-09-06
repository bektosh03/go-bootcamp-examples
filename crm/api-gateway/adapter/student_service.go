package adapter

import (
	"api-gateway/request"
	"api-gateway/response"
	"context"
	"github.com/bektosh03/crmprotos/studentpb"
)

func NewStudentService(client studentpb.StudentServiceClient) StudentService {
	return StudentService{
		client: client,
	}
}

type StudentService struct {
	client studentpb.StudentServiceClient
}

func (a StudentService) GetGroup(ctx context.Context, id string) (response.Group, error) {
	res, err := a.client.GetGroup(ctx, &studentpb.GetGroupRequest{
		GroupId: id,
	})

	if err != nil {
		return response.Group{}, err
	}

	return response.Group{
		ID:            res.Id,
		Name:          res.Name,
		MainTeacherID: res.MainTeacherId,
	}, nil
}

func (a StudentService) GetStudent(ctx context.Context, id string) (response.Student, error) {
	res, err := a.client.GetStudent(ctx, &studentpb.GetStudentRequest{
		StudentId: id})

	if err != nil {
		return response.Student{}, err
	}

	return response.Student{
		ID:          res.Id,
		FirstName:   res.FirstName,
		LastName:    res.LastName,
		Email:       res.Email,
		PhoneNumber: res.PhoneNumber,
		Level:       res.Level,
		GroupID:     res.GroupId,
	}, nil

}

func (a StudentService) CreateGroup(ctx context.Context, req request.CreateGroupRequest) (response.Group, error) {
	grpcRequest := &studentpb.CreateGroupRequest{
		Name:          req.Name,
		MainTeacherId: req.MainTeacherID,
	}

	res, err := a.client.CreateGroup(ctx, grpcRequest)
	if err != nil {
		return response.Group{}, err
	}

	return response.Group{
		ID:            res.Id,
		Name:          res.Name,
		MainTeacherID: res.MainTeacherId,
	}, nil
}

func (a StudentService) RegisterStudent(ctx context.Context, req request.RegisterStudentRequest) (response.Student, error) {
	grpcRequest := &studentpb.RegisterStudentRequest{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Level:       req.Level,
		GroupId:     req.GroupID,
	}
	res, err := a.client.RegisterStudent(ctx, grpcRequest)
	if err != nil {
		return response.Student{}, err
	}

	return response.Student{
		ID:          res.Id,
		FirstName:   res.FirstName,
		LastName:    res.LastName,
		Email:       res.Email,
		PhoneNumber: res.PhoneNumber,
		Level:       res.Level,
		GroupID:     res.GroupId,
	}, nil
}
