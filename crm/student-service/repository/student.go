package repository

import "github.com/google/uuid"

// Student represents student.Student struct for repository usage
type Student struct {
	ID          uuid.UUID `db:"id"`
	FirstName   string    `db:"first_name"`
	LastName    string    `db:"last_name"`
	Email       string    `db:"email"`
	PhoneNumber string    `db:"phone_number"`
	Level       int32     `db:"level"`
	Password    string    `db:"password"`
	GroupID     uuid.UUID `db:"group_id"`
}
