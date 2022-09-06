package group

import (
	"github.com/bektosh03/crmcommon/id"
	"github.com/google/uuid"
)

// NewFactory initializes a new Factory
func NewFactory(idGenerator id.IGenerator) Factory {
	return Factory{
		idGenerator: idGenerator,
	}
}

// Factory is a struct that creates new instances of Group. Read more about factory pattern
type Factory struct {
	idGenerator id.IGenerator
}

// NewGroup creates a new instance of Group with given data, checking its validity
func (f Factory) NewGroup(name string, mainTeacherID uuid.UUID) (Group, error) {
	s := Group{
		id:            f.idGenerator.GenerateUUID(),
		name:          name,
		mainTeacherID: mainTeacherID,
	}

	if err := s.validate(); err != nil {
		return Group{}, err
	}

	return s, nil
}
