package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/bektosh03/crmcommon/errs"
	"github.com/bektosh03/crmcommon/postgres"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
	"log"
	"student-service/domain/group"
	"student-service/domain/student"
)

const (
	studentsTableName = "students"
	groupsTableName   = "groups"
)

func NewPostgres(cfg postgres.Config) (*Postgres, error) {
	db, err := postgres.Connect(cfg)
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

func (p *Postgres) GetStudentsByGroup(ctx context.Context, groupID uuid.UUID) ([]student.Student, error) {
	repoStudents, err := p.getStudentsByGroup(ctx, groupID)
	if err != nil {
		return nil, err
	}

	return toDomainStudents(repoStudents)
}

func (p *Postgres) getStudentsByGroup(ctx context.Context, groupID uuid.UUID) ([]Student, error) {
	query := `
	SELECT * FROM students WHERE group_id = $1
	`
	students := make([]Student, 0)
	if err := p.db.SelectContext(ctx, &students, query, groupID); err != nil {
		return nil, err
	}

	return students, nil
}

// ListGroups ...
func (p *Postgres) ListGroups(ctx context.Context, page, limit int32) ([]group.Group, int, error) {
	repoGroups, count, err := p.listGroups(ctx, page, limit)
	if err != nil {
		return nil, 0, err
	}

	groups, err := toDomainGroups(repoGroups)
	if err != nil {
		return nil, 0, err
	}

	return groups, count, nil
}

func (p *Postgres) listGroups(ctx context.Context, page, limit int32) ([]Group, int, error) {
	count, err := p.count(ctx, groupsTableName)
	if err != nil {
		return nil, 0, err
	}

	query := `select * from groups offset $1 limit $2`

	offset := (page - 1) * limit
	groups := make([]Group, 0)

	if err = p.db.SelectContext(ctx, &groups, query, offset, limit); err != nil {
		return nil, 0, err
	}

	return groups, count, nil
}

// DeleteGroup ...
func (p *Postgres) DeleteGroup(ctx context.Context, id uuid.UUID) error {
	return p.deleteGroup(ctx, id)
}

func (p *Postgres) deleteGroup(ctx context.Context, id uuid.UUID) error {
	query := `delete from groups where id = $1`

	_, err := p.db.ExecContext(ctx, query, id)

	return err
}

// UpdateGroup ...
func (p *Postgres) UpdateGroup(ctx context.Context, g group.Group) error {
	return p.updateGroup(ctx, g)
}

func (p *Postgres) updateGroup(ctx context.Context, g group.Group) error {
	query := `update groups set name = $1, main_teacher_id = $2, id = $3`
	_, err := p.db.ExecContext(ctx, query, g.Name(), g.MainTeacherID(), g.ID())

	return err
}

// ListStudents ...
func (p *Postgres) ListStudents(ctx context.Context, page, limit int32) ([]student.Student, int, error) {
	repoStudent, count, err := p.listStudents(ctx, page, limit)
	if err != nil {
		return nil, 0, err
	}

	students, err := toDomainStudents(repoStudent)
	if err != nil {
		return nil, 0, err
	}

	return students, count, nil
}

func (p *Postgres) listStudents(ctx context.Context, page, limit int32) ([]Student, int, error) {
	count, err := p.count(ctx, studentsTableName)
	if err != nil {
		return nil, 0, err
	}

	query := `select * from students offset $1 limit $2`

	offset := (page - 1) * limit
	students := make([]Student, 0)
	if err = p.db.SelectContext(ctx, &students, query, offset, limit); err != nil {
		return nil, 0, err
	}

	return students, count, nil
}

// DeleteStudent ...
func (p *Postgres) DeleteStudent(ctx context.Context, id uuid.UUID) error {
	return p.deleteStudent(ctx, id)
}

func (p *Postgres) deleteStudent(ctx context.Context, id uuid.UUID) error {
	query := `delete from students where id = $1`

	_, err := p.db.ExecContext(ctx, query, id)

	return err
}

// UpdateStudent ...
func (p *Postgres) UpdateStudent(ctx context.Context, s student.Student) error {
	return p.updateStudent(ctx, toRepositoryStudent(s))
}

func (p *Postgres) updateStudent(ctx context.Context, s Student) error {
	bp, err := bcrypt.GenerateFromPassword([]byte(s.Password), bcrypt.MinCost)
	query := `update students 
	set first_name = $1, last_name = $2, email = $3, phone_number = $4, level = $5, password = $6, group_id = $7 
	where id = $8`
	_, err = p.db.ExecContext(
		ctx, query,
		s.FirstName, s.LastName, s.Email, s.PhoneNumber, s.Level, string(bp), s.GroupID, s.ID,
	)

	return err
}

// GetGroup ...
func (p *Postgres) GetGroup(ctx context.Context, id uuid.UUID) (group.Group, error) {
	gr, err := p.getGroup(ctx, id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return group.Group{}, errs.ErrNotFound
		}
		return group.Group{}, err
	}

	return group.UnmarshalGroup(group.UnmarshalGroupArgs(gr))
}

func (p *Postgres) getGroup(ctx context.Context, id uuid.UUID) (Group, error) {
	query := `select * from groups where id = $1`

	var g Group
	if err := p.db.GetContext(ctx, &g, query, id); err != nil {
		return Group{}, err
	}

	return g, nil
}

// GetStudent ...
func (p *Postgres) GetStudent(ctx context.Context, by student.By) (student.Student, error) {
	s, err := p.getStudent(ctx, by)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return student.Student{}, errs.ErrNotFound
		}
		return student.Student{}, err
	}

	return student.UnmarshalStudent(student.UnmarshalStudentArgs(s))
}

func (p *Postgres) getStudent(ctx context.Context, by student.By) (Student, error) {
	var (
		query string
		arg   interface{}
	)

	switch b := by.(type) {
	case student.ByID:
		query = `select * from students where id = $1`
		arg = b.ID
	case student.ByEmail:
		query = `select * from students where email = $1`
		arg = b.Email
	case student.ByPhoneNumber:
		query = `select * from students where phone_number = $1`
		arg = b.PhoneNumber
	}

	var s Student
	if err := p.db.GetContext(ctx, &s, query, arg); err != nil {
		return Student{}, err
	}

	return s, nil
}

// CreateStudent ...
func (p *Postgres) CreateStudent(ctx context.Context, s student.Student) error {
	return p.createStudent(ctx, toRepositoryStudent(s))
}

func (p *Postgres) createStudent(ctx context.Context, s Student) error {
	bp, err := bcrypt.GenerateFromPassword([]byte(s.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}

	query := `insert into students 
    (id, first_name, last_name, email, phone_number, level, password, group_id)
	values ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err = p.db.ExecContext(
		ctx, query,
		s.ID, s.FirstName, s.LastName, s.Email, s.PhoneNumber, s.Level, string(bp), s.GroupID,
	)

	return err
}

// CreateGroup ...
func (p *Postgres) CreateGroup(ctx context.Context, g group.Group) error {
	return p.createGroup(ctx, toRepositoryGroup(g))
}

func (p *Postgres) createGroup(ctx context.Context, g Group) error {
	query := `insert into groups values ($1, $2, $3)`

	_, err := p.db.ExecContext(ctx, query, g.ID, g.Name, g.MainTeacherID)
	return err
}

func (p *Postgres) count(ctx context.Context, table string) (int, error) {
	query := fmt.Sprintf("select count(*) from %s", table)

	var count int
	if err := p.db.GetContext(ctx, &count, query); err != nil {
		return 0, err
	}

	return count, nil
}

func (p *Postgres) cleanUp() func() {
	return func() {
		query := `delete from groups`
		_, err := p.db.Exec(query)
		if err != nil {
			log.Println(err)
		}

		query = `delete from students`
		_, err = p.db.Exec(query)
		if err != nil {
			log.Println(err)
		}
	}
}
