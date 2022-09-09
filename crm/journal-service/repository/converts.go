package repository

import (
	"database/sql"
	"journal-service/domain/journal"
)

func toRepositoryJournal(jour journal.Journal) Journal {
	return Journal{
		ID:         jour.ID(),
		ScheduleID: jour.ScheduleID(),
		Date:       jour.Date(),
	}
}

func toRepositoryJournalStatus(st journal.Status) Status {
	return Status{
		JournalID: st.JournalID(),
		StudentID: st.StudentID(),
		Attended:  st.Attended(),
		Mark: sql.NullInt32{
			Int32: st.Mark(),
			Valid: true,
		},
	}
}
