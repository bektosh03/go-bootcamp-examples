package server

import (
	"github.com/bektosh03/crmprotos/studentpb"
	"student-service/domain/group"
	"student-service/domain/student"
)

func toProtoGroups(grs []group.Group) *studentpb.GroupList {
	protoGrs := make([]*studentpb.Group, 0, len(grs))
	for _, item := range grs {
		protoGrs = append(protoGrs, toProtoGroup(item))
	}
	return &studentpb.GroupList{
		Groups: protoGrs,
	}
}

func toProtoStudents(ss []student.Student) *studentpb.StudentList {
	protoSs := make([]*studentpb.Student, 0, len(ss))
	for _, item := range ss {
		protoSs = append(protoSs, toProtoStudent(item))
	}
	return &studentpb.StudentList{
		Students: protoSs,
	}
}

func toProtoStudent(s student.Student) *studentpb.Student {
	return &studentpb.Student{
		Id:          s.ID().String(),
		FirstName:   s.FirstName(),
		LastName:    s.LastName(),
		Email:       s.Email(),
		PhoneNumber: s.PhoneNumber(),
		Level:       s.Level(),
		Password:    s.Password(),
		GroupId:     s.GroupID().String(),
	}
}

func toProtoGroup(g group.Group) *studentpb.Group {
	return &studentpb.Group{
		Id:            g.ID().String(),
		Name:          g.Name(),
		MainTeacherId: g.MainTeacherID().String(),
	}
}
