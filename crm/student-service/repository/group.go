package repository

import "github.com/google/uuid"

type Group struct {
	ID            uuid.UUID `db:"id"`
	Name          string    `db:"name"`
	MainTeacherID uuid.UUID `db:"main_teacher_id"`
}
