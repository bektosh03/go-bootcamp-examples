package repository

import (
	"context"
	"schedule-service/domain/schedule"
	"testing"
	"time"

	"github.com/bektosh03/crmcommon/errs"
	"github.com/bektosh03/crmcommon/id"
	"github.com/bektosh03/crmcommon/postgres"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPostgresRepository(t *testing.T) {
	// INITIALIZATION
	cfg := postgres.Config{
		PostgresHost:           "localhost",
		PostgresPort:           "5432",
		PostgresUser:           "postgres",
		PostgresPassword:       "1234",
		PostgresDB:             "test_schedule_db",
		PostgresMigrationsPath: "migrations",
	}
	p, err := NewPostgresRepository(cfg)
	require.NoError(t, err)
	t.Cleanup(p.cleanUp())
	f := schedule.NewFactory(id.Generator{})

	t.Run("crud", func(t *testing.T) {
		t.Cleanup(p.cleanUp())
		// TESTING CREATE

		sch, err := f.NewSchedule(testGroupID, testSubjectID, testTeacherID, testWeekday, testLessonNumber)
		require.NoError(t, err)

		err = p.CreateSchedule(context.Background(), sch)
		require.NoError(t, err)

		// TESTING GET

		got, err := p.GetSchedule(context.Background(), sch.ID())
		require.NoError(t, err)
		require.Equal(t, sch, got)

		// TESTING UPDATE
		err = sch.SetLessonNumber(4)
		require.NoError(t, err)

		err = p.UpdateSchedule(context.Background(), sch)
		require.NoError(t, err)

		got, err = p.GetSchedule(context.Background(), sch.ID())
		require.NoError(t, err)
		require.Equal(t, sch, got)

		// TEST DELETE

		err = p.DeleteSchedule(context.Background(), sch.ID())
		require.NoError(t, err)

		got, err = p.GetSchedule(context.Background(), sch.ID())
		require.ErrorIs(t, err, errs.ErrNotFound)
		require.Empty(t, got)
	})

	t.Run("get full schedule for teacher", func(t *testing.T) {
		t.Cleanup(p.cleanUp())

		schedules := genFullTestSchedulesForTeacher(f, testTeacherID)
		fillWithTestSchedules(p, schedules)

		got, err := p.GetFullScheduleForTeacher(context.Background(), testTeacherID)
		require.NoError(t, err)

		assert.Equal(t, schedules, got)
	})

	t.Run("get specific weekday for teacher", func(t *testing.T) {
		t.Cleanup(p.cleanUp())

		schedules := genTestSchedulesForTeacherOnWeekday(f, testTeacherID, time.Monday)
		fillWithTestSchedules(p, schedules)

		got, err := p.GetScheduleForTeacher(context.Background(), testTeacherID, time.Monday)
		require.NoError(t, err)

		assert.Equal(t, schedules, got)
	})

	t.Run("get full schedule for group", func(t *testing.T) {
		t.Cleanup(p.cleanUp())
		schedules := genFullTestSchedulesForGroup(f, testGroupID)
		fillWithTestSchedules(p, schedules)

		got, err := p.GetFullScheduleForGroup(context.Background(), testGroupID)
		require.NoError(t, err)

		assert.Equal(t, schedules, got)
	})

	t.Run("get spicific weekday schedule for group", func(t *testing.T) {
		t.Cleanup(p.cleanUp())

		schedules := genTestSchedulesForGroupOnWeekday(f, testGroupID, time.Monday)
		fillWithTestSchedules(p, schedules)

		got, err := p.GetScheduleForGroup(context.Background(), testGroupID, time.Monday)
		require.NoError(t, err)

		assert.Equal(t, schedules, got)
	})
}

func fillWithTestSchedules(r *PostgresRepository, schedules []schedule.Schedule) {
	for _, sch := range schedules {
		err := r.CreateSchedule(context.Background(), sch)
		if err != nil {
			panic(err)
		}
	}
}

func genFullTestSchedulesForTeacher(scheduleFactory schedule.Factory, teacherID uuid.UUID) []schedule.Schedule {
	schedules := make([]schedule.Schedule, 0)
	for i := time.Weekday(1); i < 6; i++ {
		schs := genTestSchedulesForTeacherOnWeekday(scheduleFactory, teacherID, i)
		schedules = append(schedules, schs...)
	}

	return schedules
}

func genTestSchedulesForTeacherOnWeekday(scheduleFactory schedule.Factory, teacherID uuid.UUID, weekday time.Weekday) []schedule.Schedule {
	schedules := make([]schedule.Schedule, 0, 3)
	for i := int32(1); i < 4; i++ {
		sch, err := scheduleFactory.NewSchedule(
			uuid.New(),
			uuid.New(),
			teacherID,
			weekday,
			i,
		)
		if err != nil {
			panic(err)
		}

		schedules = append(schedules, sch)
	}

	return schedules
}

func genFullTestSchedulesForGroup(scheduleFactory schedule.Factory, groupID uuid.UUID) []schedule.Schedule {
	schedules := make([]schedule.Schedule, 0)
	for i := time.Weekday(1); i < 6; i++ {
		schs := genTestSchedulesForGroupOnWeekday(scheduleFactory, groupID, i)
		schedules = append(schedules, schs...)
	}

	return schedules
}

func genTestSchedulesForGroupOnWeekday(scheduleFactory schedule.Factory, groupID uuid.UUID, weekday time.Weekday) []schedule.Schedule {
	schedules := make([]schedule.Schedule, 0, 3)
	for i := int32(1); i < 4; i++ {
		sch, err := scheduleFactory.NewSchedule(
			groupID,
			uuid.New(),
			uuid.New(),
			weekday,
			i,
		)
		if err != nil {
			panic(err)
		}

		schedules = append(schedules, sch)
	}

	return schedules
}

var (
	testLessonNumber = int32(3)
	testWeekday      = time.Monday
	testGroupID      = uuid.New()
	testTeacherID    = uuid.New()
	testSubjectID    = uuid.New()
)
