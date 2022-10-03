package producer

import "encoding/json"

type EventFor string

const (
	EventForTeacher = "teacher"
	EventForStudent = "student"
)

type RegisteredEvent struct {
	Email    string   `json:"email"`
	FullName string   `json:"full_name"`
	For      EventFor `json:"type"` // teacher or student
}

type StudentMarkedEvent struct {
	Mark      int32  `json:"mark"`
	StudentID string `json:"student_id"`
	JournalID string `json:"journal_id"`
}

func (e RegisteredEvent) Encode() ([]byte, error) {
	return json.Marshal(e)
}

func (e RegisteredEvent) Length() int {
	data, _ := e.Encode()
	return len(data)
}

func (e StudentMarkedEvent) Encode() ([]byte, error) {
	return json.Marshal(e)
}

func (e StudentMarkedEvent) Length() int {
	data, _ := e.Encode()
	return len(data)
}
