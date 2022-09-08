package service

import (
	"context"
	"github.com/google/uuid"
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

func (s Service) GetJournal(ctx context.Context, id uuid.UUID) (journal.Journal, error) {
	jour, err := s.JouRepo.GetJournal(ctx, id)
	if err != nil {
		return journal.Journal{}, nil
	}
	return jour, nil
}
func (s Service) UpdateJournal(ctx context.Context, jour journal.Journal) error {
	return s.JouRepo.UpdateJournal(ctx, jour)
}
func (s Service) DeleteJournal(ctx context.Context, id uuid.UUID) error {
	return s.JouRepo.DeleteJournal(ctx, id)
}
