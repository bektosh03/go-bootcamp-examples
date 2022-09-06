package student

import (
	"github.com/google/uuid"
	"reflect"
	"testing"
)

func TestFactory_NewStudent(t *testing.T) {
	type args struct {
		firstName   string
		lastName    string
		email       string
		phoneNumber string
		level       int32
		groupID     uuid.UUID
	}
	tests := []struct {
		name    string
		args    args
		want    Student
		wantErr bool
	}{
		{
			name: "should pass",
			args: args{
				firstName:   "Pulat",
				lastName:    "Nazarov",
				email:       "nazarovpulat3855@gmail.com",
				phoneNumber: "+998971234567",
				level:       3,
				groupID:     testGroupID,
			},
			want: Student{
				id:          testStudentID,
				firstName:   "Pulat",
				lastName:    "Nazarov",
				email:       "nazarovpulat3855@gmail.com",
				phoneNumber: "+998971234567",
				level:       3,
				groupID:     testGroupID,
			},
			wantErr: false,
		},
		{
			name: "with empty name",
			args: args{
				firstName:   "",
				lastName:    "Nazarov",
				email:       "nazarovpulat3855@gmail.com",
				phoneNumber: "+998971234567",
				level:       3,
				groupID:     testGroupID,
			},
			want:    Student{},
			wantErr: true,
		},
		{
			name: "with empty surname",
			args: args{
				firstName:   "Pulat",
				lastName:    "",
				email:       "nazarovpulat3855@gmail.com",
				phoneNumber: "+998971234567",
				level:       3,
				groupID:     testGroupID,
			},
			want:    Student{},
			wantErr: true,
		},
		{
			name: "with invalid email",
			args: args{
				firstName:   "Pulat",
				lastName:    "Nazarov",
				email:       "nn@gail.com",
				phoneNumber: "+998971234567",
				level:       3,
				groupID:     testGroupID,
			},
			want:    Student{},
			wantErr: true,
		},
		{
			name: "with invalid phone number",
			args: args{
				firstName:   "Pulat",
				lastName:    "Nazarov",
				email:       "nazarovpulat3855@gmail.com",
				phoneNumber: "+99971234567",
				level:       3,
				groupID:     testGroupID,
			},
			want:    Student{},
			wantErr: true,
		},
		{
			name: "with invalid level",
			args: args{
				firstName:   "Pulat",
				lastName:    "Nazarov",
				email:       "nazarovpulat3855@gmail.com",
				phoneNumber: "+998971234567",
				level:       9,
				groupID:     testGroupID,
			},
			want:    Student{},
			wantErr: true,
		},
	}

	f := NewFactory(testIDGenerator{})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := f.NewStudent(tt.args.firstName, tt.args.lastName, tt.args.email, tt.args.phoneNumber, tt.args.level, tt.args.groupID)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewStudent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStudent() got = %v, want %v", got, tt.want)
			}
		})
	}
}

var (
	testGroupID   = uuid.New()
	testStudentID = uuid.New()
)

type testIDGenerator struct{}

func (g testIDGenerator) GenerateUUID() uuid.UUID {
	return testStudentID
}
