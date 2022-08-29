package teacher

import (
	"errors"
	"fmt"
	"teacher-service/pkg/validate"

	"github.com/google/uuid"
)

var (
	// ErrInvalidTeacherData means that data passed for constructing Teacher structure was bad
	ErrInvalidTeacherData = errors.New("invalid teacher data")
)

// Teacher represents domain object that holds required info for a teacher
// All core business logic relevant to teachers should be done through this struct
type Teacher struct {
	id          uuid.UUID
	firstName   string
	lastName    string
	email       string
	phoneNumber string
	subjectID   uuid.UUID
}

func (t Teacher) ID() uuid.UUID {
	return t.id
}

func (t Teacher) FirstName() string {
	return t.firstName
}

func (t Teacher) LastName() string {
	return t.lastName
}

func (t Teacher) Email() string {
	return t.email
}

func (t Teacher) PhoneNumber() string {
	return t.phoneNumber
}

func (t Teacher) SubjectID() uuid.UUID {
	return t.subjectID
}

func (t Teacher) validate() error {
	if t.firstName == "" {
		return fmt.Errorf("%w: empty first name", ErrInvalidTeacherData)
	}
	if t.lastName == "" {
		return fmt.Errorf("%w: empty last name", ErrInvalidTeacherData)
	}
	if err := validate.Email(t.email); err != nil {
		return fmt.Errorf("%w: %v", ErrInvalidTeacherData, err)
	}
	if err := validate.PhoneNumber(t.phoneNumber); err != nil {
		return fmt.Errorf("%w: %v", ErrInvalidTeacherData, err)
	}

	return nil
}

type UnmarshalTeacherArgs struct {
	ID          uuid.UUID
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	SubjectID   uuid.UUID
}

func UnmarshalTeacher(args UnmarshalTeacherArgs) (Teacher, error) {
	t := Teacher{
		id:          args.ID,
		firstName:   args.FirstName,
		lastName:    args.LastName,
		email:       args.Email,
		phoneNumber: args.PhoneNumber,
		subjectID:   args.SubjectID,
	}
	if err := t.validate(); err != nil {
		return Teacher{}, err
	}

	return t, nil
}
