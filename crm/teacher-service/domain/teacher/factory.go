package teacher

import (
	"teacher-service/pkg/id"

	"github.com/google/uuid"
)

// Factory constructs new Teacher domain objects
type Factory struct {
	idGenerator id.IGenerator
}

func NewFactory(idGenerator id.IGenerator) Factory {
	return Factory{
		idGenerator: idGenerator,
	}
}

// NewTeacher is a constructor that checks if the provided data for Teacher is valid or not
// new Teacher objects can only be created through this constuctor which ensures everything is valid
func (f Factory) NewTeacher(
	firstName, lastName, email, phoneNumber, password string,
	subjectID uuid.UUID,
) (Teacher, error) {
	t := Teacher{
		id:          f.idGenerator.GenerateUUID(),
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		phoneNumber: phoneNumber,
		password: password,
		subjectID:   subjectID,
	}

	if err := t.validate(); err != nil {
		return Teacher{}, err
	}

	return t, nil
}
