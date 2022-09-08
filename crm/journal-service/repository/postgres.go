package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/bektosh03/crmcommon/errs"
	"github.com/bektosh03/crmcommon/postgres"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"journal-service/domain/journal"
)

type PostgresRepository struct {
	db *sqlx.DB
}

func NewPostgresRepository(cfg postgres.Config) (*PostgresRepository, error) {
	db, err := postgres.Connect(cfg)
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{
		db: db,
	}, nil
}

func (r *PostgresRepository) CreateJournal(ctx context.Context, jour journal.Journal) error {
	return r.createJournal(ctx, toRepositoryJournal(jour))
}

func (r *PostgresRepository) createJournal(ctx context.Context, jour Journal) error {
	query := ` insert into journals values ($1,$2,$3,$4)`
	_, err := r.db.ExecContext(ctx, query, jour.ID, jour.ScheduleID, jour.StudentID, jour.Attended, jour.Mark)
	return err
}

func (r *PostgresRepository) GetJournal(ctx context.Context, id uuid.UUID) (journal.Journal, error) {
	j, err := r.getJournal(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return journal.Journal{}, errs.ErrNotFound
		}
		return journal.Journal{}, err
	}
	return journal.UnmarshalJournal(journal.UnmarshalJournalArgs(j))
}

func (r *PostgresRepository) getJournal(ctx context.Context, id uuid.UUID) (Journal, error) {
	query := `SELECT * FROM journals where id = $1`
	var j Journal
	if err := r.db.GetContext(ctx, &j, query, id); err != nil {
		return Journal{}, err
	}
	return j, nil
}

func (r *PostgresRepository) UpdateJournal(ctx context.Context, jour journal.Journal) error {
	return r.updateJournal(ctx, toRepositoryJournal(jour))
}

func (r *PostgresRepository) updateJournal(ctx context.Context, jour Journal) error {
	query := `UPDATE journals SET schedule_id = $1, student_id = $2, attended = $3, mark = $4 where id = $5`
	_, err := r.db.ExecContext(ctx, query, jour.ScheduleID, jour.StudentID, jour.Attended, jour.Mark, jour.ID)
	return err
}

func (r *PostgresRepository) DeleteJournal(ctx context.Context, id uuid.UUID) error {
	return r.deleteJournal(ctx, id)
}

func (r *PostgresRepository) deleteJournal(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM journals WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
