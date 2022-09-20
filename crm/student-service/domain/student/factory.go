package student

import (
	"github.com/bektosh03/crmcommon/id"
	"github.com/google/uuid"
)

// Factory constructs new Student domain objects
type Factory struct {
	idGenerator id.IGenerator
}

func NewFactory(idGenerator id.IGenerator) Factory {
	return Factory{
		idGenerator: idGenerator,
	}
}

// NewStudent is a constructor that checks if the provided data for Student is valid or not
// new Student objects can only be created through this constructor which ensures everything is valid
func (f Factory) NewStudent(
	firstName, lastName, email, phoneNumber, password string,
	level int32,
	groupID uuid.UUID,
) (Student, error) {
	t := Student{
		id:          f.idGenerator.GenerateUUID(),
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		phoneNumber: phoneNumber,
		level:       level,
		password:    password,
		groupID:     groupID,
	}

	if err := t.validate(); err != nil {
		return Student{}, err
	}

	return t, nil
}
