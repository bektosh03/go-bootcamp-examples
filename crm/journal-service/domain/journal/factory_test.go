package journal

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"time"
)

func TestFactory_NewJournal(t *testing.T) {
	type args struct {
		scheduleId uuid.UUID
		teacherID uuid.UUID
		date       time.Time
	}

	dateValue, err := time.Parse("2006-01-02", date)
	require.NoError(t, err)

	tests := []struct {
		name    string
		args    args
		want    Journal
		wantErr bool
	}{
		{
			name: "should pass",
			args: args{
				scheduleId: testScheduleID,
				teacherID: testTeacherID,
				date:       dateValue,
			},
			want: Journal{
				id:         testJournalID,
				scheduleID: testScheduleID,
				teacherID: testTeacherID,
				date:       dateValue,
			},
			wantErr: false,
		},
		{
			name: "should fail with date",
			args: args{
				scheduleId: testScheduleID,
				date:       time.Time{},
			},
			want:    Journal{},
			wantErr: true,
		},
	}

	f := NewFactory(testIDGenerator{})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := f.NewJournal(tt.args.scheduleId, tt.args.teacherID,tt.args.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewJournal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewJournal() got = %v, want %v", got, tt.want)
			}
		})
	}
}

var (
	testScheduleID = uuid.New()
	testJournalID  = uuid.New()
	testTeacherID = uuid.New()
	date           = "2022-09-14"
)

type testIDGenerator struct{}

func (t testIDGenerator) GenerateUUID() uuid.UUID {
	return testJournalID
}
