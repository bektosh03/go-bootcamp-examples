package repository

import "journal-service/domain/journal"

func toRepositoryJournal(jour journal.Journal) Journal {
	return Journal{
		ID:         jour.Id(),
		ScheduleID: jour.ScheduleId(),
		StudentID:  jour.StudentId(),
		Attended:   jour.Attended(),
		Mark:       jour.Mark(),
	}
}
