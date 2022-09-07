package request

import "time"

type CreateScheduleRequest struct {
	GroupID      string       `json:"group_id"`
	SubjectID    string       `json:"subject_id"`
	TeacherID    string       `json:"teacher_id"`
	WeekDay      time.Weekday `json:"week_day"`
	LessonNumber int32        `json:"lesson_number"`
}

type Schedule struct {
	ID           string       `json:"id"`
	GroupID      string       `json:"group_id"`
	SubjectID    string       `json:"subject_id"`
	TeacherID    string       `json:"teacher_id"`
	WeekDay      time.Weekday `json:"week_day"`
	LessonNumber int32        `json:"lesson_number"`
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

type Group struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	MainTeacherID string `json:"main_teacher_id"`
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
	Name        string `json:"name"`
	Description string `json:"description"`
}
