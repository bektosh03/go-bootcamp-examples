package adapter

import (
	"api-gateway/request"
	"api-gateway/response"
	"context"
	"time"

	"github.com/bektosh03/crmprotos/schedulepb"
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
