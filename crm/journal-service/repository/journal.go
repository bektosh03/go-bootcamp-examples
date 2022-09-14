package repository

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Journal struct {
	ID         uuid.UUID `db:"id"`
	ScheduleID uuid.UUID `db:"schedule_id"`
	TeacherID  uuid.UUID `db:"teacher_id"`
	Date       time.Time `db:"date"`
}

type Stats struct {
	JournalID uuid.UUID     `db:"journal_id"`
	StudentID uuid.UUID     `db:"student_id"`
	Attended  bool          `db:"attended"`
	Mark      sql.NullInt32 `db:"mark"`
}

type Entry struct {
	JournalID  uuid.UUID     `db:"journal_id"`
	ScheduleID uuid.UUID     `db:"schedule_id"`
	Date       time.Time     `db:"date"`
	StudentID  uuid.UUID     `db:"student_id"`
	Attended   bool          `db:"attended"`
	Mark       sql.NullInt32 `db:"mark"`
}
