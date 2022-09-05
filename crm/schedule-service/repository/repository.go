package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/bektosh03/crmcommon/errs"
	"github.com/google/uuid"
	"log"
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

func (r *PostgresRepository) GetFullScheduleForGroup(ctx context.Context, groupId uuid.UUID) ([]schedule.Schedule, error) {
	schedules, err := r.getFullScheduleForGroup(ctx, groupId)
	if err != nil {
		return nil, err
	}

	return convertListToDomainSchedules(schedules)
}

func (r *PostgresRepository) getFullScheduleForGroup(ctx context.Context, groupId uuid.UUID) ([]Schedule, error) {
	query := `SELECT * FROM schedules WHERE group_id = $1`

	schedules := make([]Schedule, 0)
	if err := r.db.SelectContext(ctx, &schedules, query, groupId); err != nil {
		return nil, err
	}
	return schedules, nil
}

func (r *PostgresRepository) DeleteSchedule(ctx context.Context, id uuid.UUID) error {
	return r.deleteSchedule(ctx, id)
}

func (r *PostgresRepository) deleteSchedule(ctx context.Context, id uuid.UUID) error {
	query := `
	DELETE FROM schedules WHERE id = $1
`

	_, err := r.db.ExecContext(ctx, query, id)

	return err
}

func (r *PostgresRepository) UpdateSchedule(ctx context.Context, sch schedule.Schedule) error {
	return r.updateSchedule(ctx, sch)
}

func (r *PostgresRepository) updateSchedule(ctx context.Context, sch schedule.Schedule) error {
	query := `
	UPDATE schedules SET group_id = $1, subject_id = $2, teacher_id = $3, weekday = $4, lesson_number = $5 WHERE id = $6
`
	_, err := r.db.ExecContext(ctx, query, sch.GroupID(), sch.SubjectID(), sch.TeacherID(), sch.Weekday(), sch.LessonNumber(), sch.ID())

	return err
}

func (r *PostgresRepository) GetSchedule(ctx context.Context, id uuid.UUID) (schedule.Schedule, error) {
	sch, err := r.getSchedule(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return schedule.Schedule{}, errs.ErrNotFound
		}
		return schedule.Schedule{}, err
	}

	return schedule.UnmarshalSchedule(schedule.UnmarshalArgs(sch))
}

func (r *PostgresRepository) getSchedule(ctx context.Context, id uuid.UUID) (Schedule, error) {
	query := `
	SELECT * FROM schedules WHERE id = $1
`
	var sch Schedule
	if err := r.db.GetContext(ctx, &sch, query, id); err != nil {
		return Schedule{}, err
	}

	return sch, nil
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

func (r *PostgresRepository) cleanUp() func() {
	return func() {
		query := `DELETE FROM schedules`
		_, err := r.db.Exec(query)
		if err != nil {
			log.Println(err)
		}
	}
}
