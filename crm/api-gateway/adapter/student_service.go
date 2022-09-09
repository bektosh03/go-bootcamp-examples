package adapter

import (
	"api-gateway/pkg/httperr"
	"api-gateway/request"
	"api-gateway/response"
	"context"
	"fmt"
	"github.com/bektosh03/crmprotos/studentpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewStudentService(client studentpb.StudentServiceClient) StudentService {
	return StudentService{
		client: client,
	}
}

type StudentService struct {
	client studentpb.StudentServiceClient
}

func (a StudentService) ListGroups(ctx context.Context, page, limit int32) ([]response.Group, error) {
	res, err := a.client.ListGroups(ctx, &studentpb.ListGroupsRequest{
		Page:  page,
		Limit: limit,
	})
	if err != nil {
		return nil, err
	}

	return fromProtoToResponseGroups(res)
}

func (a StudentService) DeleteGroup(ctx context.Context, groupID string) error {
	if _, err := a.client.DeleteGroup(ctx, &studentpb.DeleteGroupRequest{
		GroupId: groupID,
	}); err != nil {
		return err
	}

	return nil
}

func (a StudentService) UpdateGroup(ctx context.Context, gr request.Group) (response.Group, error) {
	grpcRequest := &studentpb.Group{
		Id:            gr.ID,
		Name:          gr.Name,
		MainTeacherId: gr.MainTeacherID,
	}
	res, err := a.client.UpdateGroup(ctx, grpcRequest)
	if err != nil {
		return response.Group{}, err
	}

	return response.Group{
		ID:            res.Id,
		Name:          res.Name,
		MainTeacherID: res.MainTeacherId,
	}, nil
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

func (a StudentService) ListStudents(ctx context.Context, page, limit int32) ([]response.Student, error) {
	students, err := a.client.ListStudents(ctx, &studentpb.ListStudentsRequest{
		Page:  page,
		Limit: limit,
	})
	if err != nil {
		return nil, err
	}

	return fromProtoToResponseStudents(students)
}

func (a StudentService) DeleteStudent(ctx context.Context, studentID string) error {
	if _, err := a.client.DeleteStudent(ctx, &studentpb.DeleteStudentRequest{
		StudentId: studentID,
	}); err != nil {
		return err
	}

	return nil
}

func (a StudentService) UpdateStudent(ctx context.Context, st request.Student) (response.Student, error) {
	grpcRequest := &studentpb.Student{
		Id:          st.ID,
		FirstName:   st.FirstName,
		LastName:    st.LastName,
		Email:       st.Email,
		PhoneNumber: st.PhoneNumber,
		Level:       st.Level,
		GroupId:     st.GroupID,
	}
	res, err := a.client.UpdateStudent(ctx, grpcRequest)
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

func (a StudentService) GetStudent(ctx context.Context, req request.GetStudentRequest) (response.Student, error) {
	var (
		res *studentpb.Student
		err error
	)
	if req.StudentID != "" {
		res, err = a.client.GetStudent(ctx, &studentpb.GetStudentRequest{
			By: &studentpb.GetStudentRequest_StudentId{
				StudentId: req.StudentID,
			},
		})
	} else if req.Email != "" {
		res, err = a.client.GetStudent(ctx, &studentpb.GetStudentRequest{
			By: &studentpb.GetStudentRequest_Email{
				Email: req.Email,
			},
		})
	} else if req.PhoneNumber != "" {
		res, err = a.client.GetStudent(ctx, &studentpb.GetStudentRequest{
			By: &studentpb.GetStudentRequest_PhoneNumber{
				PhoneNumber: req.PhoneNumber,
			},
		})
	} else {
		return response.Student{}, fmt.Errorf("%w: %s", httperr.ErrBadRequest, "searching data is not provided")
	}

	if err != nil {
		if sts, ok := status.FromError(err); ok {
			switch sts.Code() {
			case codes.InvalidArgument:
				return response.Student{}, fmt.Errorf("%w: %s", httperr.ErrBadRequest, sts.Message())
			case codes.NotFound:
				return response.Student{}, fmt.Errorf("%w: %s", httperr.ErrNotFound, sts.Message())
			default:
				return response.Student{}, fmt.Errorf("%w: %s", httperr.ErrInternal, sts.Message())
			}
		}
		return response.Student{}, err
	}

	return fromProtoToResponseStudent(res), nil

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
