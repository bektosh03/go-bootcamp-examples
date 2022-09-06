package group

import (
	"github.com/google/uuid"
	"reflect"
	"testing"
)

func TestFactory_NewGroup(t *testing.T) {
	type args struct {
		name          string
		mainTeacherID uuid.UUID
	}
	tests := []struct {
		name    string
		args    args
		want    Group
		wantErr bool
	}{
		{
			name: "should pass",
			args: args{
				name:          "EK-32",
				mainTeacherID: testMainTeacherID,
			},
			want: Group{
				id:            testGroupID,
				name:          "EK-32",
				mainTeacherID: testMainTeacherID,
			},
			wantErr: false,
		},
		{
			name: "with empty group name",
			args: args{
				name:          "",
				mainTeacherID: testMainTeacherID,
			},
			want:    Group{},
			wantErr: true,
		},
	}

	f := NewFactory(testIDGenerator{})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := f.NewGroup(tt.args.name, tt.args.mainTeacherID)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSubject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSubject() got = %v, want %v", got, tt.want)
			}
		})
	}
}

type testIDGenerator struct{}

func (g testIDGenerator) GenerateUUID() uuid.UUID {
	return testGroupID
}

var (
	testMainTeacherID = uuid.New()
	testGroupID       = uuid.New()
)
