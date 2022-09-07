package adapter

import (
	"api-gateway/response"
	"github.com/bektosh03/crmprotos/studentpb"
)

func fromProtoToResponseGroups(groups *studentpb.GroupList) ([]response.Group, error) {
	grs := make([]response.Group, 0, len(groups.Groups))
	for _, item := range groups.GetGroups() {
		grs = append(grs, fromProtoToResponseGroup(item))
	}
	return grs, nil
}

func fromProtoToResponseGroup(group *studentpb.Group) response.Group {
	return response.Group{
		ID:            group.Id,
		Name:          group.Name,
		MainTeacherID: group.MainTeacherId,
	}
}

func fromProtoToResponseStudents(students *studentpb.StudentList) ([]response.Student, error) {
	sts := make([]response.Student, 0, len(students.Students))
	for _, item := range students.Students {
		sts = append(sts, fromProtoToResponseStudent(item))
	}
	return sts, nil
}

func fromProtoToResponseStudent(student *studentpb.Student) response.Student {
	return response.Student{
		ID:          student.Id,
		FirstName:   student.FirstName,
		LastName:    student.LastName,
		Email:       student.Email,
		PhoneNumber: student.PhoneNumber,
		Level:       student.Level,
		GroupID:     student.GroupId,
	}
}
