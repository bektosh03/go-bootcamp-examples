package service

import (
	"context"
	"journal-service/domain/journal"
)

type Service struct {
	JouRepo JournalRepository
	JourFac journal.Factory
}

func New(jouRepo JournalRepository, journalFactory journal.Factory) Service {
	return Service{
		JouRepo: jouRepo,
		JourFac: journalFactory,
	}
}

func (s Service) RegisterJournal(ctx context.Context, journal journal.Journal) error {
	return s.JouRepo.CreateJournal(ctx, journal)
}
