package student

import (
	"errors"
	"fmt"
	"github.com/bektosh03/crmcommon/validate"
	"github.com/google/uuid"
)

var (
	//ErrInvalidStudentData means that data passed for constructing Student structure was bad
	ErrInvalidStudentData = errors.New("invalid student data")
)

// Student represents domain object that holds required info for a student]
// All core business logic relevant to students should be done through this struct
type Student struct {
	id          uuid.UUID
	firstName   string
	lastName    string
	email       string
	phoneNumber string
	level       int32
	password    string
	groupID     uuid.UUID
}

func (t Student) ID() uuid.UUID {
	return t.id
}

func (t Student) FirstName() string {
	return t.firstName
}

func (t Student) LastName() string {
	return t.lastName
}

func (t Student) Email() string {
	return t.email
}

func (t Student) PhoneNumber() string {
	return t.phoneNumber
}

func (t Student) Level() int32 {
	return t.level
}

func (t Student) Password() string {
	return t.password
}

func (t Student) GroupID() uuid.UUID {
	return t.groupID
}

func (t Student) validate() error {
	if t.firstName == "" {
		return fmt.Errorf("%w: empty first name", ErrInvalidStudentData)
	}
	if t.lastName == "" {
		return fmt.Errorf("%w: empty last name", ErrInvalidStudentData)
	}
	//if err := validate.Email(t.email); err != nil {
	//	return fmt.Errorf("%w: %v", ErrInvalidStudentData, err)
	//}
	if t.password == "" {
		fmt.Errorf("%w: empty password", ErrInvalidStudentData)
	}
	if err := validate.PhoneNumber(t.phoneNumber); err != nil {
		return fmt.Errorf("%w: %v", ErrInvalidStudentData, err)
	}
	if t.level <= 0 || t.level >= 5 {
		return fmt.Errorf("%w: invalid student level data", ErrInvalidStudentData)
	}

	return nil
}

type UnmarshalStudentArgs struct {
	ID          uuid.UUID
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	Level       int32
	Password    string
	GroupID     uuid.UUID
}

func UnmarshalStudent(args UnmarshalStudentArgs) (Student, error) {
	t := Student{
		id:          args.ID,
		firstName:   args.FirstName,
		lastName:    args.LastName,
		email:       args.Email,
		phoneNumber: args.PhoneNumber,
		level:       args.Level,
		password:    args.Password,
		groupID:     args.GroupID,
	}
	if err := t.validate(); err != nil {
		return Student{}, err
	}

	return t, nil
}
