package repository

import "github.com/google/uuid"

// Teacher represents teacher.Teacher struct for repository usage
type Teacher struct {
	ID          uuid.UUID `db:"id"`
	FirstName   string    `db:"first_name"`
	LastName    string    `db:"last_name"`
	Email       string    `db:"email"`
	PhoneNumber string    `db:"phone_number"`
	Password    string    `db:"password"`
	SubjectID   uuid.UUID `db:"subject_id"`
}
