package repository

import "schedule-service/domain/schedule"

func convertListToDomainSchedules(schedules []Schedule) ([]schedule.Schedule, error) {
	schs := make([]schedule.Schedule, 0, len(schedules))
	for _, item := range schedules {
		sch, err := schedule.UnmarshalSchedule(schedule.UnmarshalArgs(item))
		if err != nil {
			return nil, err
		}
		schs = append(schs, sch)
	}
	return schs, nil
}

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
