package service

import (
	"context"
	"journal-service/domain/journal"
)

type JournalRepository interface {
	CreateJournal(ctx context.Context, j journal.Journal) error
}
