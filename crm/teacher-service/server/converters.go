package server

import (
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
		SubjectId:   t.SubjectID().String(),
	}
}
