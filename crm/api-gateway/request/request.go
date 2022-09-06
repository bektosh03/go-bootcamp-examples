package request

import "time"


type CreateScheduleRequest struct {
	GroupID      string       `json:"group_id,omitempty"`
	SubjectID    string       `json:"subject_id,omitempty"`
	TeacherID    string       `json:"teacher_id,omitempty"`
	WeekDay      time.Weekday `json:"week_day,omitempty"`
	LessonNumber int32        `json:"lesson_number,omitempty"`
}

type CreateGroupRequest struct {
	Name          string `json:"name"`
	MainTeacherID string `json:"main_teacher_id"`
}

type RegisterStudentRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Level       int32  `json:"level"`
	GroupID     string `json:"group_id"`
}

type RegisterTeacherRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	SubjectID   string `json:"subject_id"`
}

type CreateSubjectRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}
