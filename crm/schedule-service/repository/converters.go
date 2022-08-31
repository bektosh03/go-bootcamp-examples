package repository

import "schedule-service/domain/schedule"

func toRepositorySchedule(sch schedule.Schedule) Schedule {
	return Schedule{
		ID:           sch.ID(),
		GroupID:      sch.GroupID(),
		SubjectID:    sch.SubjectID(),
		TeacherID:    sch.TeacherID(),
		Weekday:      sch.Weekday(),
		LessonNumber: sch.LessonNumber(),
	}
}
