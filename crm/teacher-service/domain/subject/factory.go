package subject

import "teacher-service/pkg/id"

// NewFactory initializes a new Factory
func NewFactory(idGenerator id.IGenerator) Factory {
	return Factory{
		idGenerator: idGenerator,
	}
}

// Factory is a struct that creates new instances of Subject. Read more about factory pattern
type Factory struct {
	idGenerator id.IGenerator
}

// NewSubject creates a new instance of Subject with given data, checking its validity
func (f Factory) NewSubject(name, description string) (Subject, error) {
	s := Subject{
		id:          f.idGenerator.GenerateUUID(),
		name:        name,
		description: description,
	}

	if err := s.validate(); err != nil {
		return Subject{}, err
	}

	return s, nil
}
