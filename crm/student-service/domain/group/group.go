package group

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

var (
	// ErrInvalidGroupData indicates that data passed to create Group instance was invalid
	ErrInvalidGroupData = errors.New("invalid group data")
)

// Group represents domain object that holds required info for a group
// All core business logic relevant to groups should be done through this struct
type Group struct {
	id            uuid.UUID
	name          string
	mainTeacherID uuid.UUID
}

func (s Group) ID() uuid.UUID {
	return s.id
}

func (s Group) Name() string {
	return s.name
}

func (s Group) MainTeacherID() uuid.UUID {
	return s.mainTeacherID
}

func (s Group) validate() error {
	if s.name == "" {
		return fmt.Errorf("%w: name is empty", ErrInvalidGroupData)
	}

	return nil
}

type UnmarshalGroupArgs struct {
	ID            uuid.UUID
	Name          string
	MainTeacherID uuid.UUID
}

func UnmarshalGroup(args UnmarshalGroupArgs) (Group, error) {
	s := Group{
		id:            args.ID,
		name:          args.Name,
		mainTeacherID: args.MainTeacherID,
	}

	if err := s.validate(); err != nil {
		return Group{}, err
	}

	return s, nil
}
