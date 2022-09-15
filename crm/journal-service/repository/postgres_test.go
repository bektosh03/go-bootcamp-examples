package repository

import (
	"context"
	"github.com/bektosh03/crmcommon/errs"
	"github.com/stretchr/testify/assert"
	"journal-service/domain/journal"
	"log"
	"testing"
	"time"

	"github.com/bektosh03/crmcommon/id"
	"github.com/bektosh03/crmcommon/postgres"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

var testPostgresCfg = postgres.Config{
	PostgresHost:           "localhost",
	PostgresPort:           "5432",
	PostgresUser:           "postgres",
	PostgresPassword:       "2307",
	PostgresDB:             "crm_test",
	PostgresMigrationsPath: "migrations",
}

var (
	testStudentID  = uuid.New()
	testScheduelID = uuid.New()
	testTeacherID = uuid.New()
)

func TestPostgresRepository(t *testing.T) {

	p, err := NewPostgresRepository(testPostgresCfg)
	require.NoError(t, err)

	t.Cleanup(cleanup(p))

	journalFac := journal.NewFactory(id.Generator{})
	testDate, _ := time.Parse("2006-01-02 15:04:05 -0700", "2022-09-20 00:00:00 +0000")

	t.Run("create journal", func(t *testing.T) {
		t.Cleanup(cleanup(p))
		want, err := journalFac.NewJournal(testScheduelID, testTeacherID,testDate)
		require.NoError(t, err)

		err = p.CreateJournal(context.Background(), want)
		require.NoError(t, err)

		got, err := p.GetJournal(context.Background(), want.ID())

		require.NoError(t, err)
		assert.Equal(t, want, got)
	})

	t.Run("create journal_stats", func(t *testing.T) {
		t.Cleanup(cleanup(p))
		j, err := journalFac.NewJournal(testScheduelID, testTeacherID, testDate)
		require.NoError(t, err)
		err = p.CreateJournal(context.Background(), j)
		require.NoError(t, err)

		err = p.CreateJournalStats(
			context.Background(),
			j.ID(),
			generateFakeStudentIDs(10),
		)
		require.NoError(t, err)
	})

	t.Run("delet journal", func(t *testing.T) {
		t.Cleanup(cleanup(p))
		j, err := journalFac.NewJournal(testScheduelID, testTeacherID, testDate)
		require.NoError(t, err)

		err = p.CreateJournal(context.Background(), j)
		require.NoError(t, err)

		err = p.DeleteJournal(context.Background(), j.ID())
		require.NoError(t, err)

		got, err := p.GetJournal(context.Background(), j.ID())
		require.EqualError(t, err, errs.ErrNotFound.Error())
		assert.Empty(t, got)
	})

	t.Run("update journal", func(t *testing.T) {
		t.Cleanup(cleanup(p))
		j, err := journalFac.NewJournal(testScheduelID, testTeacherID, testDate)
		require.NoError(t, err)

		err = p.CreateJournal(context.Background(), j)
		require.NoError(t, err)

		j.SetScheduleID(uuid.New())
		err = p.UpdateJournal(context.Background(), j)
		require.NoError(t, err)

		got, err := p.GetJournal(context.Background(), j.ID())
		require.NoError(t, err)
		assert.Equal(t, j, got)
	})

	t.Run("mark student", func(t *testing.T) {
		t.Cleanup(cleanup(p))
		j, err := journalFac.NewJournal(testScheduelID, testTeacherID, testDate)
		require.NoError(t, err)
		err = p.CreateJournal(context.Background(), j)
		require.NoError(t, err)

		js, err := journal.NewStatus(j.ID(), testStudentID, true, 5)
		require.NoError(t, err)

		err = p.MarkStudent(context.Background(), js)
		require.NoError(t, err)

	})

	t.Run("set student's attendance", func(t *testing.T) {
		t.Cleanup(cleanup(p))
		j, err := journalFac.NewJournal(testScheduelID, testTeacherID, testDate)
		require.NoError(t, err)

		err = p.CreateJournal(context.Background(), j)
		require.NoError(t, err)

		js, err := journal.NewStatus(j.ID(), testStudentID, true, 5)
		require.NoError(t, err)

		err = p.SetStudentAttendance(context.Background(), js)
		require.NoError(t, err)
	})

}

func generateFakeStudentIDs(n int) []uuid.UUID {
	ids := make([]uuid.UUID, 0)
	for i := 1; i <= n; i++ {
		ids = append(ids, uuid.New())
	}

	return ids
}

func cleanup(p *PostgresRepository) func() {
	return func() {
		if err := p.cleanup(context.Background()); err != nil {
			log.Println("failed to cleanup db, should be done manually", err)
		}
	}
}