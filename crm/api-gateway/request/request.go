package request

import (
	"encoding/json"
	"time"
)

type AdminRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type MarkStudentRequest struct {
	JournalID string `json:"journal_id"`
	StudentID string `json:"student_id"`
	Mark      int32  `json:"mark"`
}

type SetStudentAttendanceRequest struct {
	StudentID string `json:"student_id"`
	Attended  bool   `json:"attended"`
	JournalID string `json:"journal_id"`
}

type GetSpecificDateScheduleForTeacherRequest struct {
	TeacherID string    `json:"teacher_id"`
	Date      time.Time `json:"date"` // date should be in format: YYYY-MM-DD
}

func (r *GetSpecificDateScheduleForTeacherRequest) UnmarshalJSON(data []byte) error {
	m := make(map[string]string)
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	date, err := time.Parse("2006-01-02", m["date"])
	if err != nil {
		return err
	}

	r.TeacherID = m["teacher_id"]
	r.Date = date

	return nil
}

type GetSpecificDateScheduleForGroupRequest struct {
	GroupID string    `json:"group_id"`
	Date    time.Time `json:"date"`
}

func (r *GetSpecificDateScheduleForGroupRequest) UnmarshalJSON(data []byte) error {
	m := make(map[string]string)
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	date, err := time.Parse("2006-01-02", m["date"])
	if err != nil {
		return err
	}

	r.GroupID = m["group_id"]
	r.Date = date

	return nil
}

type Journal struct {
	ID         string    `json:"id"`
	ScheduleID string    `json:"schedule_id"`
	Date       time.Time `json:"date"`
}

type CreateJournalRequest struct {
	ScheduleID string    `json:"schedule_id"`
	Date       time.Time `json:"date"`
}

func (r *CreateJournalRequest) UnmarshalJSON(data []byte) error {
	m := make(map[string]string)
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	date, err := time.Parse("2006-01-02", m["date"])
	if err != nil {
		return err
	}

	r.ScheduleID = m["schedule_id"]
	r.Date = date

	return nil
}

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
	Password    string `json:"password"`
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
	Password    string `json:"password"`
	GroupID     string `json:"group_id"`
}

type RegisterTeacherRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
	SubjectID   string `json:"subject_id"`
}

type CreateSubjectRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type GetTeacherRequest struct {
	TeacherID   string `json:"teacher_id,omitempty"`
	Email       string `json:"email,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
}

type GetStudentRequest struct {
	StudentID   string `json:"student_id,omitempty"`
	Email       string `json:"email,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
}
