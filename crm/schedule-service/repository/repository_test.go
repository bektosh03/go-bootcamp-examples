package repository

import (
	"context"
	"github.com/bektosh03/crmcommon/errs"
	"github.com/bektosh03/crmcommon/id"
	"github.com/bektosh03/crmcommon/postgres"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"schedule-service/domain/schedule"
	"testing"
	"time"
)

func TestPostgresRepository(t *testing.T) {
	// INITIALIZATION
	cfg := postgres.Config{
		PostgresHost:           "localhost",
		PostgresPort:           "5432",
		PostgresUser:           "pulat",
		PostgresPassword:       "9",
		PostgresDB:             "test_schedule_db",
		PostgresMigrationsPath: "migrations",
	}
	p, err := NewPostgresRepository(cfg)
	require.NoError(t, err)
	t.Cleanup(p.cleanUp())
	f := schedule.NewFactory(id.Generator{})

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
}

var (
	testLessonNumber = int32(3)
	testWeekday      = time.Monday
	testGroupID      = uuid.New()
	testTeacherID    = uuid.New()
	testSubjectID    = uuid.New()
)
