package repository

import (
	"context"
	"schedule-service/domain/schedule"

	"github.com/bektosh03/crmcommon/postgres"
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

func (r *PostgresRepository) CreateSchedule(ctx context.Context, sch schedule.Schedule) error {
	return r.createSchedule(ctx, toRepositorySchedule(sch))
}

func (r *PostgresRepository) createSchedule(ctx context.Context, sch Schedule) error {
	query := `
	INSERT INTO schedules VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := r.db.ExecContext(
		ctx, query,
		sch.ID, sch.GroupID, sch.SubjectID, sch.TeacherID, sch.Weekday, sch.LessonNumber,
	)

	return err
}
