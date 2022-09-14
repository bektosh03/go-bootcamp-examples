package repository

import (
	"database/sql"
	"journal-service/domain/journal"
)

func toDomainEntry(entry Entry) (journal.Entry, error) {
	return journal.UnmarshalEntry(journal.UnmarshalEntryArgs{
		JournalID:  entry.JournalID,
		ScheduleID: entry.ScheduleID,
		Date:       entry.Date,
		StudentID:  entry.StudentID,
		Attended:   entry.Attended,
		Mark:       entry.Mark.Int32,
	})
}

func toDomainEntries(dbEntries []Entry) ([]journal.Entry, error) {
	entries := make([]journal.Entry, 0, len(dbEntries))
	for _, dbEntry := range dbEntries {
		entry, err := toDomainEntry(dbEntry)
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}

	return entries, nil
}

func toRepositoryJournal(jour journal.Journal) Journal {
	return Journal{
		ID:         jour.ID(),
		ScheduleID: jour.ScheduleID(),
		TeacherID:  jour.TeacherID(),
		Date:       jour.Date(),
	}
}

func toRepositoryJournalStatus(st journal.Stat) Stats {
	return Stats{
		JournalID: st.JournalID(),
		StudentID: st.StudentID(),
		Attended:  st.Attended(),
		Mark: sql.NullInt32{
			Int32: st.Mark(),
			Valid: true,
		},
	}
}
