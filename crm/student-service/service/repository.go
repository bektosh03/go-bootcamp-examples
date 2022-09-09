package service

import (
	"student-service/domain/group"
	"student-service/domain/student"
)

// Repository defines methods that should be present in our storage provider
type Repository interface {
	student.Repository
	group.Repository
}
