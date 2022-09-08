package service

import (
	"teacher-service/domain/subject"
	"teacher-service/domain/teacher"
)

// Repository defines methods that should be present in our storage provider
type Repository interface {
	teacher.Repository
	subject.Repository
}
