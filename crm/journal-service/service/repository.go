package service

import (
	"context"
	"journal-service/domain/journal"

	"github.com/google/uuid"
)

type JournalRepository interface {
	CreateJournal(ctx context.Context, journal journal.Journal) error
	GetJournal(ctx context.Context, id uuid.UUID) (journal.Journal, error)
	UpdateJournal(ctx context.Context, journal journal.Journal) error
	DeleteJournal(ctx context.Context, id uuid.UUID) error
	CreateJournalStatuses(ctx context.Context, journalID uuid.UUID, studentIDs []uuid.UUID) error
}
