package repository

import (
	"context"
	"database/sql"
	"errors"
	"journal-service/domain/journal"

	"github.com/bektosh03/crmcommon/errs"
	"github.com/bektosh03/crmcommon/postgres"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
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
	query := `INSERT INTO journal VALUES ($1,$2,$3)`
	_, err := r.db.ExecContext(ctx, query, jour.ID, jour.ScheduleID, jour.Date)
	return err
}

func (r *PostgresRepository) CreateJournalStatuses(ctx context.Context, journalID uuid.UUID, studentIDs []uuid.UUID) error {
	return r.createJournalStatuses(ctx, journalID, studentIDs)
}

func (r *PostgresRepository) createJournalStatuses(ctx context.Context, journalID uuid.UUID, studentIDs []uuid.UUID) error {
	query := `
	INSERT INTO journal_status VALUES ($1, $2)
	`
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()

	for _, studentID := range studentIDs {
		_, err := tx.ExecContext(ctx, query, journalID, studentID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return nil
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
	query := `SELECT * FROM journal where id = $1`
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
	query := `UPDATE journal SET schedule_id = $1, date = $2 where id = $3`
	_, err := r.db.ExecContext(ctx, query, jour.ScheduleID, jour.Date, jour.ID)
	return err
}

func (r *PostgresRepository) DeleteJournal(ctx context.Context, id uuid.UUID) error {
	return r.deleteJournal(ctx, id)
}

func (r *PostgresRepository) deleteJournal(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM journal WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *PostgresRepository) MarkStudent(ctx context.Context, st journal.Status) error {
	return r.markStudent(ctx, toRepositoryJournalStatus(st))
}

func (r *PostgresRepository) markStudent(ctx context.Context, st Status) error {
	query := `
	UPDATE journal_status SET mark = $1, attended = TRUE
	WHERE student_id = $2 AND journal_id = $3
	`
	_, err := r.db.ExecContext(ctx, query, st.Mark, st.StudentID, st.JournalID)
	return err
}
