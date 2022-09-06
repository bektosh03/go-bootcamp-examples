package subject

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

var (
	// ErrInvalidSubjectData indicates that data passed to create Subject instance was invalid
	ErrInvalidSubjectData = errors.New("invalid subject data")
)

// Subject represents domain object that holds required info for a subject
// All core business logic relevant to subjects should be done through this struct
type Subject struct {
	id          uuid.UUID
	name        string
	description string
}

// Getters ...
func (s Subject) ID() uuid.UUID {
	return s.id
}

func (s Subject) Name() string {
	return s.name
}

func (s Subject) Description() string {
	return s.description
}

// Setters ...
func (s *Subject) SetName(name string) error {
	s.name = name
	return s.validate()
}

func (s *Subject) SetDescription(desc string) error {
	s.description = desc
	return s.validate()
}

func (s Subject) validate() error {
	if s.name == "" {
		return fmt.Errorf("%w: name is empty", ErrInvalidSubjectData)
	}
	if s.description == "" {
		return fmt.Errorf("%w: description is empty", ErrInvalidSubjectData)
	}

	return nil
}

type UnmarshalSubjectArgs struct {
	ID          uuid.UUID
	Name        string
	Description string
}

func UnmarshalSubject(args UnmarshalSubjectArgs) (Subject, error) {
	s := Subject{
		id:          args.ID,
		name:        args.Name,
		description: args.Description,
	}

	if err := s.validate(); err != nil {
		return Subject{}, err
	}

	return s, nil
}
