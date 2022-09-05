package repository_test

import (
	"context"
	"log"
	"teacher-service/config"
	"teacher-service/domain/subject"
	"teacher-service/domain/teacher"
	"teacher-service/pkg/id"
	"teacher-service/repository"
	"testing"

	"github.com/stretchr/testify/require"
)

var testPostgresCfg = config.PostgresConfig{
	PostgresHost:           "localhost",
	PostgresPort:           "5432",
	PostgresUser:           "postgres",
	PostgresPassword:       "1234",
	PostgresDB:             "crm_test",
	PostgresMigrationsPath: "migrations",
}

func TestPostgres_CreateSubject(t *testing.T) {
	p, err := repository.NewPostgres(testPostgresCfg)
	require.NoError(t, err)

	t.Cleanup(cleanup(p))

	subjectFactory := subject.NewFactory(id.Generator{})
	teacherFactory := teacher.NewFactory(id.Generator{})

	t.Run("test for create subject", func(t *testing.T) {
		t.Cleanup(cleanup(p))
		s, err := subjectFactory.NewSubject(
			"Math",
			"description",
		)

		err = p.CreateSubject(context.Background(), s)
		require.NoError(t, err)
	})

	t.Run("test for create teacher", func(t *testing.T) {
		t.Cleanup(cleanup(p))
		s, err := subjectFactory.NewSubject(
			"Math",
			"description",
		)

		err = p.CreateSubject(context.Background(), s)
		require.NoError(t, err)

		tch, err := teacherFactory.NewTeacher(
			"First",
			"Last",
			"first@last.com",
			"+998991234567",
			s.ID(),
		)

		err = p.CreateTeacher(context.Background(), tch)
		require.NoError(t, err)
	})
}

func cleanup(p *repository.Postgres) func() {
	return func() {
		if err := p.Cleanup(context.Background()); err != nil {
			log.Panicln("failed to cleanup db, should be done manually", err)
		}
	}
}
