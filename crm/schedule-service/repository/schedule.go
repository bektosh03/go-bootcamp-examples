package repository

import (
	"time"

	"github.com/google/uuid"
)

type Schedule struct {
	ID           uuid.UUID    `db:"id"`
	GroupID      uuid.UUID    `db:"group_id"`
	SubjectID    uuid.UUID    `db:"subject_id"`
	TeacherID    uuid.UUID    `db:"teacher_id"`
	Weekday      time.Weekday `db:"weekday"`
	LessonNumber int32        `db:"lesson_number"`
}
