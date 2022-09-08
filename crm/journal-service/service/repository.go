package service

import (
	"context"
	"github.com/google/uuid"
	"journal-service/domain/journal"
)

type JournalRepository interface {
	CreateJournal(ctx context.Context, journal journal.Journal) error
	GetJournal(ctx context.Context, id uuid.UUID) (journal.Journal, error)
	UpdateJournal(ctx context.Context, journal journal.Journal) error
	DeleteJournal(ctx context.Context, id uuid.UUID) error
}
