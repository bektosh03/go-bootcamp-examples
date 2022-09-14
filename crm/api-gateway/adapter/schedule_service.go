package adapter

import (
	"api-gateway/pkg/httperr"
	"api-gateway/request"
	"api-gateway/response"
	"context"
	"fmt"
	"time"

	"github.com/bektosh03/crmprotos/schedulepb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ScheduleService struct {
	client schedulepb.ScheduleServiceClient
}

func NewScheduleService(client schedulepb.ScheduleServiceClient) ScheduleService {
	return ScheduleService{
		client: client,
	}
}
func (s ScheduleService) RegisterSchedule(ctx context.Context, req request.CreateScheduleRequest) (response.Schedule, error) {
	grpcRequest := &schedulepb.CreateScheduleRequest{
		GroupId:      req.GroupID,
		SubjectId:    req.SubjectID,
		TeacherId:    req.TeacherID,
		Weekday:      schedulepb.Weekday(req.WeekDay),
		LessonNumber: req.LessonNumber,
	}
	res, err := s.client.CreateSchedule(ctx, grpcRequest)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			switch st.Code() {
			case codes.InvalidArgument:
				return response.Schedule{}, fmt.Errorf("%w: %s", httperr.ErrBadRequest, st.Message())
			default:
				return response.Schedule{}, fmt.Errorf("%w: %s", httperr.ErrInternal, st.Message())
			}
		}

		return response.Schedule{}, err
	}

	return response.Schedule{
		ID:           res.Id,
		GroupId:      res.GroupId,
		SubjectID:    res.SubjectId,
		TeacherID:    res.TeacherId,
		WeekDay:      time.Weekday(res.Weekday),
		LessonNumber: res.LessonNumber,
	}, nil
}

func (s ScheduleService) GetSchedule(ctx context.Context, id string) (response.Schedule, error) {
	res, err := s.client.GetSchedule(ctx, &schedulepb.GetScheduleRequest{
		ScheduleId: id,
	})
	if err != nil {
		if sts, ok := status.FromError(err); ok {
			switch sts.Code() {
			case codes.InvalidArgument:
				return response.Schedule{}, fmt.Errorf("%w: %s", httperr.ErrBadRequest, sts.Message())
			case codes.NotFound:
				return response.Schedule{}, fmt.Errorf("%w: %s", httperr.ErrNotFound, sts.Message())
			default:
				return response.Schedule{}, fmt.Errorf("%w: %s", httperr.ErrInternal, sts.Message())
			}
		}
		return response.Schedule{}, err
	}
	return response.Schedule{
		ID:           id,
		GroupId:      res.GroupId,
		SubjectID:    res.SubjectId,
		TeacherID:    res.TeacherId,
		WeekDay:      time.Weekday(res.Weekday),
		LessonNumber: res.LessonNumber,
	}, nil
}

func (s ScheduleService) UpdateSchedule(ctx context.Context, req request.Schedule) (response.Schedule, error) {
	res, err := s.client.UpdateSchedule(ctx, &schedulepb.Schedule{
		Id:           req.ID,
		GroupId:      req.GroupID,
		SubjectId:    req.SubjectID,
		TeacherId:    req.TeacherID,
		Weekday:      schedulepb.Weekday(req.WeekDay),
		LessonNumber: req.LessonNumber,
	})
	if err != nil {
		if sts, ok := status.FromError(err); ok {
			switch sts.Code() {
			case codes.InvalidArgument:
				return response.Schedule{}, fmt.Errorf("%w: %s", httperr.ErrBadRequest, sts.Message())
			default:
				return response.Schedule{}, fmt.Errorf("%w: %s", httperr.ErrInternal, sts.Message())
			}
		}
		return response.Schedule{}, err
	}

	return response.Schedule{
		ID:           res.Id,
		GroupId:      res.GroupId,
		SubjectID:    res.SubjectId,
		TeacherID:    res.TeacherId,
		WeekDay:      time.Weekday(res.Weekday),
		LessonNumber: res.LessonNumber,
	}, nil
}

func (s ScheduleService) DeleteSchedule(ctx context.Context, id string) error {
	_, err := s.client.DeleteSchedule(ctx, &schedulepb.DeleteScheduleRequest{
		ScheduleId: id,
	})

	if err != nil {
		if sts, ok := status.FromError(err); ok {
			switch sts.Code() {
			case codes.InvalidArgument:
				return fmt.Errorf("%w: %s", httperr.ErrBadRequest, sts.Message())
			default:
				return fmt.Errorf("%w: %s", httperr.ErrInternal, sts.Message())
			}
		}
		return err
	}

	return err
}

func (s ScheduleService) GetFullScheduleForTeacher(ctx context.Context, teacherID string) ([]response.Schedule, error) {
	res, err := s.client.GetFullScheduleForTeacher(ctx, &schedulepb.GetFullScheduleForTeacherRequest{
		TeacherId: teacherID,
	})
	if err != nil {
		if sts, ok := status.FromError(err); ok {
			switch sts.Code() {
			case codes.InvalidArgument:
				return nil, fmt.Errorf("%w: %s", httperr.ErrBadRequest, sts.Message())
			default:
				return nil, fmt.Errorf("%w: %s", httperr.ErrInternal, sts.Message())
			}
		}
		return nil, err
	}

	return fromProtoScheduleListToResponseScheduleSlice(res), nil
}

func (s ScheduleService) GetFullScheduleForGroup(ctx context.Context, groupID string) ([]response.Schedule, error) {
	res, err := s.client.GetFullScheduleForGroup(ctx, &schedulepb.GetFullScheduleForGroupRequest{
		GroupId: groupID,
	})
	if err != nil {
		if sts, ok := status.FromError(err); ok {
			switch sts.Code() {
			case codes.InvalidArgument:
				return nil, fmt.Errorf("%w: %s", httperr.ErrBadRequest, sts.Message())
			default:
				return nil, fmt.Errorf("%w: %s", httperr.ErrInternal, sts.Message())
			}
		}
		return nil, err
	}

	return fromProtoScheduleListToResponseScheduleSlice(res), nil
}

func (s ScheduleService) GetSpecificDateScheduleForTeacher(ctx context.Context, teacherID string, date time.Time) ([]response.Schedule, error) {
	res, err := s.client.GetSpecificDateScheduleForTeacher(ctx, &schedulepb.GetSpecificDateScheduleForTeacherRequest{
		TeacherId: teacherID,
		Date: &timestamppb.Timestamp{
			Seconds: date.Unix(),
			Nanos:   int32(date.Nanosecond()),
		},
	})
	if err != nil {
		if sts, ok := status.FromError(err); ok {
			switch sts.Code() {
			case codes.InvalidArgument:
				return nil, fmt.Errorf("%w: %s", httperr.ErrBadRequest, sts.Message())
			default:
				return nil, fmt.Errorf("%w: %s", httperr.ErrInternal, sts.Message())
			}
		}
		return nil, err
	}

	return fromProtoScheduleListToResponseScheduleSlice(res), nil
}

func (s ScheduleService) GetSpecificDateScheduleForGroup(ctx context.Context, groupID string, date time.Time) ([]response.Schedule, error) {
	res, err := s.client.GetSpecificDateScheduleForGroup(ctx, &schedulepb.GetSpecificDateScheduleForGroupRequest{
		GroupId: groupID,
		Date:    &timestamppb.Timestamp{Seconds: date.Unix(), Nanos: int32(date.Nanosecond())},
	})
	if err != nil {
		if sts, ok := status.FromError(err); ok {
			switch sts.Code() {
			case codes.InvalidArgument:
				return nil, fmt.Errorf("%w: %s", httperr.ErrBadRequest, sts.Message())
			default:
				return nil, fmt.Errorf("%w: %s", httperr.ErrInternal, sts.Message())
			}
		}
		return nil, err
	}

	return fromProtoScheduleListToResponseScheduleSlice(res), nil
}
