package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"journal-service/domain/journal"
	"time"

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
	query := `INSERT INTO journals VALUES ($1, $2, $3, $4)`
	_, err := r.db.ExecContext(ctx, query, jour.ID, jour.ScheduleID, jour.Date, jour.TeacherID)
	return err
}

func (r *PostgresRepository) CreateJournalStats(ctx context.Context, journalID uuid.UUID, studentIDs []uuid.UUID) error {
	return r.createJournalStatuses(ctx, journalID, studentIDs)
}

func (r *PostgresRepository) createJournalStatuses(ctx context.Context, journalID uuid.UUID, studentIDs []uuid.UUID) error {
	query := `
	INSERT INTO journal_stats VALUES ($1, $2)
	`
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		tx.Commit()
		fmt.Println("COMMITTED TRANSACTION")
	}()

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
	query := `SELECT * FROM journals WHERE id = $1`
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
	query := `UPDATE journals SET schedule_id = $1, date = $2 where id = $3`
	_, err := r.db.ExecContext(ctx, query, jour.ScheduleID, jour.Date, jour.ID)
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

func (r *PostgresRepository) MarkStudent(ctx context.Context, st journal.Stat) error {
	return r.markStudent(ctx, toRepositoryJournalStatus(st))
}

func (r *PostgresRepository) markStudent(ctx context.Context, st Stats) error {
	query := `
	UPDATE journal_stats SET mark = $1, attended = TRUE
	WHERE student_id = $2 AND journal_id = $3
	`
	_, err := r.db.ExecContext(ctx, query, st.Mark, st.StudentID, st.JournalID)
	return err
}

func (r *PostgresRepository) SetStudentAttendance(ctx context.Context, st journal.Stat) error {
	return r.setStudentAttendance(ctx, st.StudentID(), st.JournalID(), st.Attended())
}

func (r *PostgresRepository) setStudentAttendance(ctx context.Context, studentID, journalID uuid.UUID, attended bool) error {
	query := `
	UPDATE journal_stats SET attended = $1
	WHERE journal_id = $2 AND student_id = $3
	`
	_, err := r.db.ExecContext(ctx, query, attended, journalID, studentID)
	return err
}

func (r *PostgresRepository) GetStudentJournalEntries(ctx context.Context, studentID uuid.UUID, start, end time.Time) ([]journal.Entry, error) {
	dbEntries, err := r.getStudentJournalEntries(ctx, studentID, start, end)
	if err != nil {
		return nil, err
	}

	return toDomainEntries(dbEntries)
}

func (r *PostgresRepository) getStudentJournalEntries(ctx context.Context, studentID uuid.UUID, start, end time.Time) ([]Entry, error) {
	query := `
	SELECT j.id as journal_id, j.schedule_id, j.date, js.student_id, js.attended, js.mark
	FROM journals j
	JOIN journal_stats js ON j.id = js.journal_id
	WHERE js.student_id = $1 AND (j.date BETWEEN $2 AND $3)
	`
	entries := make([]Entry, 0)
	if err := r.db.SelectContext(ctx, &entries, query, studentID, start, end); err != nil {
		return nil, err
	}

	return entries, nil
}

func (r *PostgresRepository) GetTeacherJournalEntries(ctx context.Context, teacherID uuid.UUID, start, end time.Time) ([]journal.Entry, error) {
	dbEntries, err := r.getTeacherJournalEntries(ctx, teacherID, start, end)
	if err != nil {
		return nil, err
	}

	return toDomainEntries(dbEntries)
}

func (r *PostgresRepository) getTeacherJournalEntries(ctx context.Context, teacherID uuid.UUID, start, end time.Time) ([]Entry, error) {
	query := `
	SELECT j.id as journal_id, j.schedule_id, j.date, js.student_id, js.attended, js.mark
	FROM journals j
	JOIN journal_stats js ON j.id = js.journal_id
	WHERE j.teacher_id = $1 AND (j.date BETWEEN $2 AND $3)
	`
	entries := make([]Entry, 0)
	if err := r.db.SelectContext(ctx, &entries, query, teacherID, start, end); err != nil {
		return nil, err
	}

	return entries, nil
}

func (r *PostgresRepository) cleanup(ctx context.Context) error {

	query := `DELETE FROM journal_stats`
	_, err := r.db.ExecContext(ctx, query)
	if err != nil {
		return err
	}
	query = `DELETE FROM journals`

	_, err = r.db.ExecContext(ctx, query)
	if err != nil {
		return err
	}

	return nil
}
