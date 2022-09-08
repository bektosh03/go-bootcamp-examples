package teacher

import "github.com/google/uuid"

type By interface {
	isTeacherBy()
}

type ByID struct {
	ID uuid.UUID
}

func (b ByID) isTeacherBy() {}

type ByEmail struct {
	Email string
}

func (b ByEmail) isTeacherBy() {}

type ByPhoneNumber struct {
	PhoneNumber string
}

func (b ByPhoneNumber) isTeacherBy() {}
