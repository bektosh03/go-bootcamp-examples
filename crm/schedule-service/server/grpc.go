package server

import (
	"context"
	"errors"
	"schedule-service/domain/schedule"
	"schedule-service/service"
	"time"

	"github.com/bektosh03/crmcommon/errs"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/bektosh03/crmprotos/schedulepb"
)

type Server struct {
	schedulepb.UnimplementedScheduleServiceServer
	svc             service.Service
	scheduleFactory schedule.Factory
}

func New(svc service.Service, fac schedule.Factory) Server {
	return Server{
		svc: svc,
		scheduleFactory: fac,
	}
}

func (s Server) GetSpecificDateScheduleForTeacher(ctx context.Context, req *schedulepb.GetSpecificDateScheduleForTeacherRequest) (*schedulepb.ScheduleList, error) {
	teacherId, err := uuid.Parse(req.GetTeacherId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "teacher id is not uuid")
	}

	schedules, err := s.svc.GetSpecificDateScheduleForTeacher(ctx, teacherId, req.Date.AsTime())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return toProtoScheduleList(schedules), nil
}

func (s Server) GetSpecificDateScheduleForGroup(ctx context.Context, req *schedulepb.GetSpecificDateScheduleForGroupRequest) (*schedulepb.ScheduleList, error) {
	groupId, err := uuid.Parse(req.GetGroupId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "group id is not uuid")
	}

	schedules, err := s.svc.GetSpecificDateScheduleForGroup(ctx, groupId, req.Date.AsTime())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return toProtoScheduleList(schedules), nil
}

func (s Server) GetFullScheduleForTeacher(ctx context.Context, req *schedulepb.GetFullScheduleForTeacherRequest) (*schedulepb.ScheduleList, error) {
	teacherId, err := uuid.Parse(req.GetTeacherId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "teacher id is not uuid")
	}
	schedules, err := s.svc.GetFullScheduleForTeacher(ctx, teacherId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return toProtoScheduleList(schedules), nil
}

func (s Server) GetFullScheduleForGroup(ctx context.Context, req *schedulepb.GetFullScheduleForGroupRequest) (*schedulepb.ScheduleList, error) {
	groupId, err := uuid.Parse(req.GetGroupId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "group id is not uuid")
	}

	schedules, err := s.svc.GetFullScheduleForGroup(ctx, groupId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return toProtoScheduleList(schedules), nil
}

func (s Server) DeleteSchedule(ctx context.Context, req *schedulepb.DeleteScheduleRequest) (*emptypb.Empty, error) {
	scheduleID, err := uuid.Parse(req.GetScheduleId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err = s.svc.DeleteSchedule(ctx, scheduleID); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s Server) UpdateSchedule(ctx context.Context, req *schedulepb.Schedule) (*schedulepb.Schedule, error) {
	sch, err := s.convertUpdateScheduleRequestToDomainSchedule(req)
	if err != nil {
		return nil, err
	}

	if err = s.svc.UpdateSchedule(ctx, sch); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return toProtoSchedule(sch), nil
}

func (s Server) GetSchedule(ctx context.Context, req *schedulepb.GetScheduleRequest) (*schedulepb.Schedule, error) {
	scheduleID, err := uuid.Parse(req.GetScheduleId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "schedule id is not uuid")
	}

	sch, err := s.svc.GetSchedule(ctx, scheduleID)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return toProtoSchedule(sch), err
}

func (s Server) CreateSchedule(ctx context.Context, req *schedulepb.CreateScheduleRequest) (*schedulepb.Schedule, error) {
	sch, err := s.convertCreateScheduleRequestToDomainSchedule(req)
	if err != nil {
		return nil, err
	}

	if err = s.svc.CreateSchedule(ctx, sch); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return toProtoSchedule(sch), nil
}

func (s Server) convertUpdateScheduleRequestToDomainSchedule(req *schedulepb.Schedule) (schedule.Schedule, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return schedule.Schedule{}, status.Error(codes.InvalidArgument, "id is not uuid")
	}

	groupID, err := uuid.Parse(req.GetGroupId())
	if err != nil {
		return schedule.Schedule{}, status.Error(codes.InvalidArgument, "group id is not uuid")
	}

	subjectID, err := uuid.Parse(req.GetSubjectId())
	if err != nil {
		return schedule.Schedule{}, status.Error(codes.InvalidArgument, "subject id is not uuid")
	}

	teacherID, err := uuid.Parse(req.GetTeacherId())
	if err != nil {
		return schedule.Schedule{}, status.Error(codes.InvalidArgument, "teacher id is not uuid")
	}

	sch, err := schedule.UnmarshalSchedule(schedule.UnmarshalArgs{
		ID:           id,
		GroupID:      groupID,
		SubjectID:    subjectID,
		TeacherID:    teacherID,
		Weekday:      time.Weekday(req.GetWeekday()),
		LessonNumber: req.GetLessonNumber(),
	})

	if err != nil {
		return schedule.Schedule{}, status.Error(codes.InvalidArgument, err.Error())
	}

	return sch, nil
}

func (s Server) convertCreateScheduleRequestToDomainSchedule(req *schedulepb.CreateScheduleRequest) (schedule.Schedule, error) {
	groupID, err := uuid.Parse(req.GetGroupId())
	if err != nil {
		return schedule.Schedule{}, status.Error(codes.InvalidArgument, "group id is not uuid")
	}

	subjectID, err := uuid.Parse(req.GetSubjectId())
	if err != nil {
		return schedule.Schedule{}, status.Error(codes.InvalidArgument, "subject id is not uuid")
	}

	teacherID, err := uuid.Parse(req.GetTeacherId())
	if err != nil {
		return schedule.Schedule{}, status.Error(codes.InvalidArgument, "teacher id is not uuid")
	}

	sch, err := s.scheduleFactory.NewSchedule(groupID, subjectID, teacherID, time.Weekday(req.Weekday), req.LessonNumber)
	if err != nil {
		return schedule.Schedule{}, status.Error(codes.InvalidArgument, err.Error())
	}

	return sch, nil
}
