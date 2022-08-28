package subject

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
)

func TestFactory_NewSubject(t *testing.T) {
	type args struct {
		name        string
		description string
	}
	tests := []struct {
		name    string
		args    args
		want    Subject
		wantErr bool
	}{
		{
			name: "should pass",
			args: args{
				name:        "Math",
				description: "study math",
			},
			want: Subject{
				id:          testSubjectID,
				name:        "Math",
				description: "study math",
			},
			wantErr: false,
		},
		{
			name: "with empty name",
			args: args{
				name:        "",
				description: "study math",
			},
			want:    Subject{},
			wantErr: true,
		},
		{
			name: "with empty description",
			args: args{
				name:        "Math",
				description: "",
			},
			want:    Subject{},
			wantErr: true,
		},
	}

	f := NewFactory(testIDGenerator{})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := f.NewSubject(tt.args.name, tt.args.description)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

type testIDGenerator struct{}

func (g testIDGenerator) GenerateUUID() uuid.UUID {
	return testSubjectID
}

var testSubjectID = uuid.New()
