package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"teacher-service/config"
	"teacher-service/domain/subject"
	"teacher-service/domain/teacher"
	"teacher-service/pkg/errs"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

const (
	teachersTableName = "teachers"
	subjectsTableName = "subjects"
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

func (p *Postgres) GetSubject(ctx context.Context, id uuid.UUID) (subject.Subject, error) {
	s, err := p.getSubject(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return subject.Subject{}, errs.ErrNotFound
		}
		return subject.Subject{}, err
	}

	return subject.UnmarshalSubject(subject.UnmarshalSubjectArgs(s))
}

func (p *Postgres) getSubject(ctx context.Context, id uuid.UUID) (Subject, error) {
	query := `
	SELECT * FROM subjects WHERE id = $1
	`
	var s Subject
	if err := p.db.GetContext(ctx, &s, query, id); err != nil {
		return Subject{}, err
	}

	return s, nil
}

func (p *Postgres) UpdateTeacher(ctx context.Context, t teacher.Teacher) error {
	return p.updateTeacher(ctx, toRepositoryTeacher(t))
}

func (p *Postgres) updateTeacher(ctx context.Context, t Teacher) error {
	query := `
	UPDATE teachers
		SET first_name = $1, last_name = $2, email = $3, phone_number = $4, subject_id = $5
	WHERE id = $6
	`
	_, err := p.db.ExecContext(
		ctx, query,
		t.FirstName, t.LastName, t.Email, t.PhoneNumber, t.SubjectID, t.ID,
	)

	return err
}

func (p *Postgres) UpdateSubject(ctx context.Context, s subject.Subject) error {
	return p.updateSubject(ctx, toRepositorySubject(s))
}

func (p *Postgres) updateSubject(ctx context.Context, s Subject) error {
	query := `
	UPDATE subjects
		SET name = $1, description = $2
	WHERE id = $3
	`
	_, err := p.db.ExecContext(ctx, query, s.Name, s.Description, s.ID)

	return err
}

func (p *Postgres) DeleteTeacher(ctx context.Context, id uuid.UUID) error {
	return p.deleteTeacher(ctx, id)
}

func (p *Postgres) deleteTeacher(ctx context.Context, id uuid.UUID) error {
	query := `
	DELETE FROM teachers
	WHERE id = $1
	`
	_, err := p.db.ExecContext(ctx, query, id)

	return err
}

func (p *Postgres) DeleteSubject(ctx context.Context, id uuid.UUID) error {
	return p.deleteSubject(ctx, id)
}

func (p *Postgres) deleteSubject(ctx context.Context, id uuid.UUID) error {
	query := `
	DELETE FROM subjects
	WHERE id = $1
	`
	_, err := p.db.ExecContext(ctx, query, id)

	return err
}

func (p *Postgres) ListTeachers(ctx context.Context, page, limit int32) ([]teacher.Teacher, int, error) {
	repoTeachers, count, err := p.listTeachers(ctx, page, limit)
	if err != nil {
		return nil, 0, err
	}

	teachers, err := toDomainTeachers(repoTeachers)
	if err != nil {
		return nil, 0, err
	}

	return teachers, count, nil
}

func (p *Postgres) listTeachers(ctx context.Context, page, limit int32) ([]Teacher, int, error) {
	count, err := p.count(ctx, teachersTableName)
	if err != nil {
		return nil, 0, err
	}

	query := `
	SELECT * FROM teachers
	OFFSET $1 LIMIT $2
	`
	offset := (page - 1) * limit
	teachers := make([]Teacher, 0)
	if err := p.db.SelectContext(ctx, &teachers, query, offset, limit); err != nil {
		return nil, 0, err
	}

	return teachers, count, nil
}

func (p *Postgres) ListSubjects(ctx context.Context, page, limit int32) ([]subject.Subject, int, error) {
	repoSubjects, count, err := p.listSubjects(ctx, page, limit)
	if err != nil {
		return nil, 0, err
	}

	subjects, err := toDomainSubjects(repoSubjects)
	if err != nil {
		return nil, 0, err
	}

	return subjects, count, nil
}

func (p *Postgres) listSubjects(ctx context.Context, page, limit int32) ([]Subject, int, error) {
	count, err := p.count(ctx, subjectsTableName)
	if err != nil {
		return nil, 0, err
	}

	query := `
	SELECT * FROM subjects
	OFFSET $1 LIMIT $2
	`
	offset := (page - 1) * limit
	subjects := make([]Subject, 0)
	if err := p.db.SelectContext(ctx, &subjects, query, offset, limit); err != nil {
		return nil, 0, err
	}

	return subjects, count, nil
}

func (p *Postgres) count(ctx context.Context, table string) (int, error) {
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s", table)

	var count int
	if err := p.db.GetContext(ctx, &count, query); err != nil {
		return 0, err
	}

	return count, nil
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
