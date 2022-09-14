package service

import (
	"context"
	"journal-service/domain/journal"
	"time"

	"github.com/google/uuid"
)

type JournalRepository interface {
	CreateJournal(ctx context.Context, journal journal.Journal) error
	GetJournal(ctx context.Context, id uuid.UUID) (journal.Journal, error)
	UpdateJournal(ctx context.Context, journal journal.Journal) error
	DeleteJournal(ctx context.Context, id uuid.UUID) error
	CreateJournalStats(ctx context.Context, journalID uuid.UUID, studentIDs []uuid.UUID) error
	MarkStudent(ctx context.Context, st journal.Stat) error
	SetStudentAttendance(ctx context.Context, st journal.Stat) error
	GetStudentJournalEntries(ctx context.Context, studentID uuid.UUID, start time.Time, end time.Time) ([]journal.Entry, error)
	GetTeacherJournalEntries(ctx context.Context, teacherID uuid.UUID, start, end time.Time) ([]journal.Entry, error)
}
