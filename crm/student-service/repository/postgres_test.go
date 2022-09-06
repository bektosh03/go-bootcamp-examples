package repository

import (
	"context"
	"github.com/bektosh03/crmcommon/id"
	"github.com/bektosh03/crmcommon/postgres"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"student-service/domain/group"
	"testing"
)

func TestPostgres(t *testing.T) {
	// INITIALIZATION
	cfg := postgres.Config{
		PostgresHost:           "localhost",
		PostgresPort:           "5432",
		PostgresUser:           "pulat",
		PostgresPassword:       "9",
		PostgresDB:             "studentdb",
		PostgresMigrationsPath: "migrations",
	}
	p, err := NewPostgres(cfg)
	require.NoError(t, err)

	t.Cleanup(p.cleanUp())

	groupFactory := group.NewFactory(id.Generator{})
	//studentFactory := student.NewFactory(id.Generator{})

	t.Run("test for create group", func(t *testing.T) {
		t.Cleanup(p.cleanUp())
		g, err := groupFactory.NewGroup(
			"Golang",
			testMainTeacherID,
		)
		require.NoError(t, err)

		err = p.CreateGroup(context.Background(), g)
		require.NoError(t, err)

		got, err := p.GetGroup(context.Background(), g.ID())
		require.NoError(t, err)
		assert.Equal(t, g, got)
	})
}

var (
	testMainTeacherID = uuid.New()
)


