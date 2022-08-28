package repository

import (
	"context"
	"teacher-service/config"
	"teacher-service/domain/subject"
	"teacher-service/domain/teacher"

	"github.com/jmoiron/sqlx"
)

func NewPostgres(cfg config.PostgresConfig) (*Postgres, error) {
	db, err := connect(cfg)
	if err != nil {
		return nil, err
	}

	return &Postgres{
		db: db,
	}, nil
}

// Postgres implements service.Repository interface
type Postgres struct {
	db *sqlx.DB
}

// CreateTeacher ...
func (p *Postgres) CreateTeacher(ctx context.Context, t teacher.Teacher) error {
	return p.createTeacher(ctx, toRepositoryTeacher(t))
}

func (p *Postgres) createTeacher(ctx context.Context, t Teacher) error {
	query := `
	INSERT INTO teachers VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := p.db.ExecContext(
		ctx, query,
		t.ID, t.FirstName, t.LastName, t.Email, t.PhoneNumber, t.SubjectID,
	)

	return err
}

func (p *Postgres) CreateSubject(ctx context.Context, s subject.Subject) error {
	return p.createSubject(ctx, toRepositorySubject(s))
}

func (p *Postgres) createSubject(ctx context.Context, s Subject) error {
	query := `
	INSERT INTO subjects VALUES ($1, $2, $3)
	`
	_, err := p.db.ExecContext(
		ctx, query,
		s.ID, s.Name, s.Description,
	)

	return err
}

func (p *Postgres) Cleanup(ctx context.Context) error {
	query := `
	DELETE FROM teachers
	`
	_, err := p.db.ExecContext(ctx, query)
	if err != nil {
		return err
	}

	query = `
	DELETE FROM subjects
	`
	_, err = p.db.ExecContext(ctx, query)
	return err
}
