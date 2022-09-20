package server

import (
	"teacher-service/domain/subject"
	"teacher-service/domain/teacher"

	"github.com/bektosh03/crmprotos/teacherpb"
)

func toProtoTeacher(t teacher.Teacher) *teacherpb.Teacher {
	return &teacherpb.Teacher{
		Id:          t.ID().String(),
		FirstName:   t.FirstName(),
		LastName:    t.LastName(),
		Email:       t.Email(),
		PhoneNumber: t.PhoneNumber(),
		Password:    t.Password(),
		SubjectId:   t.SubjectID().String(),
	}
}

func toProtoSubject(s subject.Subject) *teacherpb.Subject {
	return &teacherpb.Subject{
		Id:          s.ID().String(),
		Name:        s.Name(),
		Description: s.Description(),
	}
}
