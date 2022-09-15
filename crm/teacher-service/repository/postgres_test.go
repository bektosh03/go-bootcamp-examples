package repository_test

import (
	"context"
	"log"
	"teacher-service/domain/subject"
	"teacher-service/domain/teacher"
	"teacher-service/pkg/id"
	"teacher-service/repository"
	"testing"

	"github.com/bektosh03/crmcommon/errs"
	"github.com/bektosh03/crmcommon/postgres"
	"github.com/stretchr/testify/assert"
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

func TestPostgres_CreateAndGet(t *testing.T) {
	p, err := repository.NewPostgres(testPostgresCfg)
	require.NoError(t, err)

	t.Cleanup(cleanup(p))

	subjectFactory := subject.NewFactory(id.Generator{})
	teacherFactory := teacher.NewFactory(id.Generator{})

	t.Run("create subject and get subject", func(t *testing.T) {
		t.Cleanup(cleanup(p))
		s, err := subjectFactory.NewSubject(
			"Math",
			"description",
		)
		require.NoError(t, err)

		err = p.CreateSubject(context.Background(), s)
		require.NoError(t, err)

		got, err := p.GetSubject(context.Background(), s.ID())
		require.NoError(t, err)
		assert.Equal(t, s, got)
	})

	t.Run("create teacher and get teacher", func(t *testing.T) {
		t.Cleanup(cleanup(p))
		s, err := subjectFactory.NewSubject(
			"Math",
			"description",
		)
		require.NoError(t, err)

		err = p.CreateSubject(context.Background(), s)
		require.NoError(t, err)

		tch, err := teacherFactory.NewTeacher(
			"First",
			"Last",
			"khasanovasumbula@gmail.com",
			"+998991234567",
			"123",
			s.ID(),
		)
		require.NoError(t, err)

		err = p.CreateTeacher(context.Background(), tch)
		require.NoError(t, err)

		got, err := p.GetTeacher(context.Background(), teacher.ByID{ID: tch.ID()})
		tch.SetPassword(got.Password()) // because password comes in hashed from database
		require.NoError(t, err)
		assert.Equal(t, tch, got)
	})
}

func TestPostgres_Update(t *testing.T) {
	p, err := repository.NewPostgres(testPostgresCfg)
	require.NoError(t, err)

	t.Cleanup(cleanup(p))

	subjectFactory := subject.NewFactory(id.Generator{})
	teacherFactory := teacher.NewFactory(id.Generator{})

	t.Run("update teacher", func(t *testing.T) {
		t.Cleanup(cleanup(p))

		s, err := subjectFactory.NewSubject(
			"Math",
			"description",
		)
		require.NoError(t, err)

		err = p.CreateSubject(context.Background(), s)
		require.NoError(t, err)

		tch, err := teacherFactory.NewTeacher(
			"First",
			"Last",
			"khasanovasumbula@gmail.com",
			"+998991234567",
			"123",
			s.ID(),
		)
		require.NoError(t, err)

		err = p.CreateTeacher(context.Background(), tch)
		require.NoError(t, err)

		err = tch.SetFirstName("Khasanova")
		require.NoError(t, err)

		err = p.UpdateTeacher(context.Background(), tch)
		require.NoError(t, err)

		got, err := p.GetTeacher(context.Background(), teacher.ByID{ID: tch.ID()})
		require.NoError(t, err)
		tch.SetPassword(got.Password()) // because password comes in hashed from database
		assert.Equal(t, tch, got)
	})

	t.Run("update subject", func(t *testing.T) {
		t.Cleanup(cleanup(p))
		s, err := subjectFactory.NewSubject(
			"Math",
			"description",
		)
		require.NoError(t, err)

		err = p.CreateSubject(context.Background(), s)
		require.NoError(t, err)

		err = s.SetName("English")
		require.NoError(t, err)

		err = p.UpdateSubject(context.Background(), s)
		require.NoError(t, err)

		err = p.UpdateSubject(context.Background(), s)
		require.NoError(t, err)

		got, err := p.GetSubject(context.Background(), s.ID())
		require.NoError(t, err)
		assert.Equal(t, s, got)
	})
}

func TestPostgres_Delete(t *testing.T) {
	p, err := repository.NewPostgres(testPostgresCfg)
	require.NoError(t, err)

	t.Cleanup(cleanup(p))

	subjectFactory := subject.NewFactory(id.Generator{})
	teacherFactory := teacher.NewFactory(id.Generator{})

	t.Run("delete teacher", func(t *testing.T) {
		t.Cleanup(cleanup(p))
		s, err := subjectFactory.NewSubject(
			"Math",
			"description",
		)
		require.NoError(t, err)

		err = p.CreateSubject(context.Background(), s)
		require.NoError(t, err)

		tch, err := teacherFactory.NewTeacher(
			"First",
			"Last",
			"khasanovasumbula@gmail.com",
			"+998991234567",
			"123",
			s.ID(),
		)
		require.NoError(t, err)

		err = p.CreateTeacher(context.Background(), tch)
		require.NoError(t, err)

		err = p.DeleteTeacher(context.Background(), tch.ID())
		require.NoError(t, err)

		got, err := p.GetTeacher(context.Background(), teacher.ByID{ID: tch.ID()})
		require.ErrorIs(t, err, errs.ErrNotFound)
		assert.Equal(t, teacher.Teacher{}, got)
	})

	t.Run("delete subject", func(t *testing.T) {
		t.Cleanup(cleanup(p))
		s, err := subjectFactory.NewSubject(
			"Math",
			"description",
		)
		require.NoError(t, err)

		err = p.CreateSubject(context.Background(), s)
		require.NoError(t, err)

		err = p.DeleteSubject(context.Background(), s.ID())
		require.NoError(t, err)

		got, err := p.GetSubject(context.Background(), s.ID())
		require.ErrorIs(t, err, errs.ErrNotFound)
		assert.Equal(t, subject.Subject{}, got)
	})
}

func TestPostgres_List(t *testing.T) {
	p, err := repository.NewPostgres(testPostgresCfg)
	require.NoError(t, err)

	t.Cleanup(cleanup(p))

	subjectFactory := subject.NewFactory(id.Generator{})
	teacherFactory := teacher.NewFactory(id.Generator{})

	t.Run("list teachers", func(t *testing.T) {
		t.Cleanup(cleanup(p))
		s, err := subjectFactory.NewSubject(
			"Math",
			"description",
		)
		require.NoError(t, err)
		err = p.CreateSubject(context.Background(), s)
		require.NoError(t, err)

		want := []teacher.Teacher{}

		tch1, err := teacherFactory.NewTeacher(
			"First",
			"Last",
			"sumbulahasanova@gmail.com",
			"+998777777777",
			"123",
			s.ID(),
		)
		require.NoError(t, err)
		err = p.CreateTeacher(context.Background(), tch1)
		require.NoError(t, err)
		want = append(want, tch1)

		tch2, err := teacherFactory.NewTeacher(
			"First",
			"Last",
			"khasanovasumbula@gmail.com",
			"+998991234567",
			"123",
			s.ID(),
		)
		require.NoError(t, err)
		err = p.CreateTeacher(context.Background(), tch2)
		require.NoError(t, err)
		want = append(want, tch2)

		tch3, err := teacherFactory.NewTeacher(
			"First",
			"Last",
			"lolopepedamnn@gmail.com",
			"+998933332333",
			"123",
			s.ID(),
		)
		require.NoError(t, err)
		err = p.CreateTeacher(context.Background(), tch3)
		require.NoError(t, err)
		want = append(want, tch3)

		got, n, err := p.ListTeachers(context.Background(), 1, 3)
		require.NoError(t, err)
		assert.Equal(t, want, got)
		assert.Equal(t, 3, n)

	})

	t.Run("list subject", func(t *testing.T) {
		t.Cleanup(cleanup(p))

		want := []subject.Subject{}
		s1, err := subjectFactory.NewSubject(
			"Math",
			"description",
		)
		require.NoError(t, err)
		err = p.CreateSubject(context.Background(), s1)
		require.NoError(t, err)
		want = append(want, s1)

		s2, err := subjectFactory.NewSubject(
			"English",
			"description",
		)
		require.NoError(t, err)
		err = p.CreateSubject(context.Background(), s2)
		require.NoError(t, err)
		want = append(want, s2)

		s3, err := subjectFactory.NewSubject(
			"Biology",
			"description",
		)
		require.NoError(t, err)
		err = p.CreateSubject(context.Background(), s3)
		require.NoError(t, err)
		want = append(want, s3)

		got, n, err := p.ListSubjects(context.Background(), 1, 3)
		require.NoError(t, err)
		assert.Equal(t, want, got)
		assert.Equal(t, 3, n)

	})
}

func cleanup(p *repository.Postgres) func() {
	return func() {
		if err := p.Cleanup(context.Background()); err != nil {
			log.Panicln("failed to cleanup db, should be done manually", err)
		}
	}
}
