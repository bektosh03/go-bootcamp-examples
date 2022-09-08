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

	return fromProtoToResponseTeacher(res), nil
}

func (a TeacherService) GetTeacher(ctx context.Context, id string) (response.Teacher, error) {
	res, err := a.client.GetTeacher(ctx, &teacherpb.GetTeacherRequest{TeacherId: id})
	if err != nil {
		return response.Teacher{}, err
	}

	return fromProtoToResponseTeacher(res), nil
}

func (a TeacherService) DeleteTeacher(ctx context.Context, id string) error {
	_, err := a.client.DeleteTeacher(ctx, &teacherpb.DeleteTeacherRequest{
		TeacherId: id,
	})
	if err != nil {
		return err
	}

	return nil
}

func (a TeacherService) ListTeachers(ctx context.Context, page, limit int32) ([]response.Teacher, error) {
	res, err := a.client.ListTeachers(ctx, &teacherpb.ListTeachersRequest{
		Page:  page,
		Limit: limit,
	})

	if err != nil {
		return nil, err
	}
	return fromProtoToTeacherList(res)
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

	return fromProtoToResponseSubject(res), nil
}

func (a TeacherService) GetSubject(ctx context.Context, id string) (response.Subject, error) {
	res, err := a.client.GetSubject(ctx, &teacherpb.GetSubjectRequest{
		SubjectId: id,
	})
	if err != nil {
		return response.Subject{}, err
	}

	return fromProtoToResponseSubject(res), nil
}

func (a TeacherService) DeleteSubject(ctx context.Context, id string) error {
	_, err := a.client.DeleteSubject(ctx, &teacherpb.DeleteSubjectRequest{
		SubjectId: id,
	})

	if err != nil {
		return err
	}

	return nil
}

func (a TeacherService) ListSubjects(ctx context.Context, page, limit int32) ([]response.Subject, error) {
	resp, err := a.client.ListSubjects(ctx, &teacherpb.ListSubjectsRequest{
		Page:  page,
		Limit: limit,
	})

	if err != nil {
		return nil, err
	}

	return fromProtoToSubjectList(resp)
}
