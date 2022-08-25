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

// Factory constructs new Teacher domain objects
type Factory struct {
	idGenerator IDGenerator
}

func NewFactory(idGenerator IDGenerator) Factory {
	return Factory{
		idGenerator: idGenerator,
	}
}

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

// NewTeacher is a constructor that checks if the provided data for Teacher is valid or not
// new Teacher objects can only be created through this constuctor which ensures everything is valid
func (f Factory) NewTeacher(
	firstName, lastName, email, phoneNumber string,
	subjectID uuid.UUID,
) (Teacher, error) {
	t := Teacher{
		id:          f.idGenerator.GenerateID(),
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		phoneNumber: phoneNumber,
		subjectID:   subjectID,
	}

	if err := t.validate(); err != nil {
		return Teacher{}, err
	}

	return t, nil
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
