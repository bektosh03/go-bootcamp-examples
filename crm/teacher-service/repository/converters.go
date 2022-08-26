package repository

import (
	"teacher-service/domain/subject"
	"teacher-service/domain/teacher"
)

func toRepositoryTeacher(t teacher.Teacher) Teacher {
	return Teacher{
		ID:          t.ID(),
		FirstName:   t.FirstName(),
		LastName:    t.LastName(),
		Email:       t.Email(),
		PhoneNumber: t.PhoneNumber(),
		SubjectID:   t.SubjectID(),
	}
}

func toRepositorySubject(s subject.Subject) Subject {
	return Subject{
		ID:          s.ID(),
		Name:        s.Name(),
		Description: s.Description(),
	}
}
