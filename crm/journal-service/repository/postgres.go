package repository

import (
	"context"
	"github.com/bektosh03/crmcommon/postgres"
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

/*
func (p *Postgres) GetTeacher(ctx context.Context, id uuid.UUID) (teacher.Teacher, error) {
	t, err := p.getTeacher(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return teacher.Teacher{}, errs.ErrNotFound
		}
		return teacher.Teacher{}, err
	}

	return teacher.UnmarshalTeacher(teacher.UnmarshalTeacherArgs(t))
}

func (p *Postgres) getTeacher(ctx context.Context, id uuid.UUID) (Teacher, error) {
	query := `
	SELECT * FROM teachers WHERE id = $1
	`
	var t Teacher
	if err := p.db.GetContext(ctx, &t, query, id); err != nil {
		return Teacher{}, err
	}

	return t, nil
}
*/
