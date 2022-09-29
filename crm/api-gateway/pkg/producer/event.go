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

func (e RegisteredEvent) Encode() ([]byte, error) {
	return json.Marshal(e)
}

func (e RegisteredEvent) Length() int {
	data, _ := e.Encode()
	return len(data)
}
