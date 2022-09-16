package repository

import (
	"student-service/domain/group"
	"student-service/domain/student"
)

func toRepositoryStudent(s student.Student) Student {
	return Student{
		ID:          s.ID(),
		FirstName:   s.FirstName(),
		LastName:    s.LastName(),
		Email:       s.Email(),
		PhoneNumber: s.PhoneNumber(),
		Level:       s.Level(),
		Password:    s.Password(),
		GroupID:     s.GroupID(),
	}
}

func toRepositoryGroup(s group.Group) Group {
	return Group{
		ID:            s.ID(),
		Name:          s.Name(),
		MainTeacherID: s.MainTeacherID(),
	}
}

func toDomainStudents(repoStudents []Student) ([]student.Student, error) {
	students := make([]student.Student, 0, len(repoStudents))
	for _, repoStudent := range repoStudents {
		t, err := student.UnmarshalStudent(student.UnmarshalStudentArgs(repoStudent))
		if err != nil {
			return nil, err
		}
		students = append(students, t)
	}

	return students, nil
}

func toDomainGroups(repoGroups []Group) ([]group.Group, error) {
	groups := make([]group.Group, 0, len(repoGroups))
	for _, repoGroup := range repoGroups {
		s, err := group.UnmarshalGroup(group.UnmarshalGroupArgs(repoGroup))
		if err != nil {
			return nil, err
		}
		groups = append(groups, s)
	}

	return groups, nil
}
