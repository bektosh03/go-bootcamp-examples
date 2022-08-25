package teacher

import "github.com/google/uuid"

// IDGenerator is interface that abstracts away actual way of generating new uuids
type IDGenerator interface {
	GenerateID() uuid.UUID
}

type uuidGenerator struct{}

func (g uuidGenerator) GenerateID() uuid.UUID {
	return uuid.New()
}
