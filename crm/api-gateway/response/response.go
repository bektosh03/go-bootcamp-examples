package response

import (
	"encoding/json"
	"time"
)

type PopulatedSchedule struct {
	ID           string       `json:"id"`
	Group        Group        `json:"group"`
	Subject      Subject      `json:"subject"`
	Teacher      Teacher      `json:"teacher"`
	Weekday      time.Weekday `json:"week_day"`
	LessonNumber int32        `json:"lesson_number"`
}

func (s PopulatedSchedule) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"id":            s.ID,
		"group":         s.Group,
		"subject":       s.Subject,
		"teacher":       s.Teacher,
		"week_day":      s.Weekday.String(),
		"lesson_number": s.LessonNumber,
	}

	return json.Marshal(m)
}

type Schedule struct {
	ID           string       `json:"id"`
	GroupId      string       `json:"group_id"`
	SubjectID    string       `json:"subject_id"`
	TeacherID    string       `json:"teacher_id"`
	WeekDay      time.Weekday `json:"week_day"`
	LessonNumber int32        `json:"lesson_number"`
}
type Group struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	MainTeacherID string `json:"main_teacher_id"`
}

type Student struct {
	ID          string `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Level       int32  `json:"level"`
	GroupID     string `json:"subject_id"`
}

type Teacher struct {
	ID          string `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	SubjectID   string `json:"subject_id"`
}

type Subject struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
