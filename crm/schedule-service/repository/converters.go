package repository

import "schedule-service/domain/schedule"

func convertListToDomainSchedules(list []Schedule) ([]schedule.Schedule, error) {
	schs := make([]schedule.Schedule, 0, len(list))
	for _, item := range list {
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
