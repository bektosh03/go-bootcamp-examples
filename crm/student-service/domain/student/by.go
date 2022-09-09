package student

import "github.com/google/uuid"

type By interface {
	isStudentBy()
}

type ByID struct {
	ID uuid.UUID
}

func (b ByID) isStudentBy() {}

type ByEmail struct {
	Email string
}

func (b ByEmail) isStudentBy() {}

type ByPhoneNumber struct {
	PhoneNumber string
}

func (b ByPhoneNumber) isStudentBy() {}
