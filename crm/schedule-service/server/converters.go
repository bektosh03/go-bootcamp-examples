package server

import (
	"github.com/bektosh03/crmprotos/schedulepb"
	"schedule-service/domain/schedule"
)

func toProtoScheduleList(schs []schedule.Schedule) *schedulepb.ScheduleList {
	protoSchs := make([]*schedulepb.Schedule, 0, len(schs))
	for _, item := range schs {
		protoSchs = append(protoSchs, toProtoSchedule(item))
	}
	return &schedulepb.ScheduleList{
		Schedules: protoSchs,
	}
}

func toProtoSchedule(sch schedule.Schedule) *schedulepb.Schedule {
	return &schedulepb.Schedule{
		Id:           sch.ID().String(),
		GroupId:      sch.GroupID().String(),
		SubjectId:    sch.SubjectID().String(),
		TeacherId:    sch.TeacherID().String(),
		Weekday:      schedulepb.Weekday(sch.Weekday()),
		LessonNumber: sch.LessonNumber(),
	}
}
