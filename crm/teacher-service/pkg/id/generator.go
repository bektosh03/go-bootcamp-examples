package id

import "github.com/google/uuid"

type IGenerator interface {
	GenerateUUID() uuid.UUID
}

type Generator struct{}

func (g Generator) GenerateUUID() uuid.UUID {
	return uuid.New()
}
