package adapter

import (
	"api-gateway/request"
	"api-gateway/response"
	"context"

	"github.com/bektosh03/crmprotos/teacherpb"
)

func NewTeacherService(client teacherpb.TeacherServiceClient) TeacherService {
	return TeacherService{
		client: client,
	}
}

type TeacherService struct {
	client teacherpb.TeacherServiceClient
}

func (a TeacherService) RegisterTeacher(ctx context.Context, req request.RegisterTeacherRequest) (response.Teacher, error) {
	grpcRequest := &teacherpb.RegisterTeacherRequest{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		SubjectId:   req.SubjectID,
	}
	res, err := a.client.RegisterTeacher(ctx, grpcRequest)
	if err != nil {
		return response.Teacher{}, err
	}

	return response.Teacher{
		ID:          res.Id,
		FirstName:   res.FirstName,
		LastName:    res.LastName,
		Email:       res.Email,
		PhoneNumber: res.PhoneNumber,
		SubjectID:   res.SubjectId,
	}, nil
}

func (a TeacherService) GetTeacher(ctx context.Context, id string) (response.Teacher, error) {
	res, err := a.client.GetTeacher(ctx, &teacherpb.GetTeacherRequest{TeacherId: id})
	if err != nil {
		return response.Teacher{}, err
	}

	return response.Teacher{
		ID:          res.Id,
		FirstName:   res.FirstName,
		LastName:    res.LastName,
		Email:       res.Email,
		PhoneNumber: res.PhoneNumber,
		SubjectID:   res.SubjectId,
	}, nil
}

func (a TeacherService) CreateSubject(ctx context.Context, req request.CreateSubjectRequest) (response.Subject, error) {
	res, err := a.client.CreateSubject(
		ctx,
		&teacherpb.CreateSubjectRequest{
			Name:        req.Name,
			Description: req.Description,
		},
	)
	if err != nil {
		return response.Subject{}, err
	}

	return response.Subject{
		ID:          res.Id,
		Name:        res.Name,
		Description: res.Description,
	}, nil
}

func (a TeacherService) GetSubject(ctx context.Context, id string) (response.Subject, error) {
	res, err := a.client.GetSubject(ctx, &teacherpb.GetSubjectRequest{
		SubjectId: id,
	})
	if err != nil {
		return response.Subject{}, err
	}

	return response.Subject{
		ID:          res.Id,
		Name:        res.Name,
		Description: res.Description,
	}, nil
}
