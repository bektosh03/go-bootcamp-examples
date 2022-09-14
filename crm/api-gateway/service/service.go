package service

import (
	"api-gateway/request"
	"api-gateway/response"
	"context"
	"time"
)

func New(teacherService TeacherServiceClient, studentService StudentServiceClient, scheduleService ScheduleServiceClient, journalService JournalServiceClient) Service {
	return Service{
		Teacher:  teacherService,
		Student:  studentService,
		Schedule: scheduleService,
		Journal:  journalService,
	}
}

type Service struct {
	Teacher  TeacherServiceClient
	Student  StudentServiceClient
	Schedule ScheduleServiceClient
	Journal  JournalServiceClient
}

type JournalServiceClient interface {
	RegisterJournal(ctx context.Context, scheduleID, teacherID string, date time.Time, studentIDs []string) (response.Journal, error)
	GetJournal(ctx context.Context, journalId string) (response.Journal, error)
	UpdateJournal(ctx context.Context, journal request.Journal) (response.Journal, error)
	DeleteJournal(ctx context.Context, id string) error
	SetStudentAttendance(ctx context.Context, req request.SetStudentAttendanceRequest) error
	MarkStudent(ctx context.Context, req request.MarkStudentRequest) error
	GetStudentJournal(ctx context.Context, studentID string, start, end time.Time) ([]response.JournalEntry, error)
	GetTeacherJournal(ctx context.Context, teacherID string, start, end time.Time) ([]response.JournalEntry, error)
}

type ScheduleServiceClient interface {
	RegisterSchedule(context.Context, request.CreateScheduleRequest) (response.Schedule, error)
	GetSchedule(context.Context, string) (response.Schedule, error)
	UpdateSchedule(ctx context.Context, req request.Schedule) (response.Schedule, error)
	DeleteSchedule(ctx context.Context, id string) error
	GetFullScheduleForTeacher(ctx context.Context, teacherID string) ([]response.Schedule, error)
	GetFullScheduleForGroup(ctx context.Context, groupID string) ([]response.Schedule, error)
	GetSpecificDateScheduleForTeacher(ctx context.Context, teacherID string, date time.Time) ([]response.Schedule, error)
	GetSpecificDateScheduleForGroup(ctx context.Context, groupID string, date time.Time) ([]response.Schedule, error)
}

type StudentServiceClient interface {
	RegisterStudent(context.Context, request.RegisterStudentRequest) (response.Student, error)
	GetStudent(context.Context, request.GetStudentRequest) (response.Student, error)
	UpdateStudent(context.Context, request.Student) (response.Student, error)
	DeleteStudent(context.Context, string) error
	ListStudents(context.Context, int32, int32) ([]response.Student, error)
	CreateGroup(context.Context, request.CreateGroupRequest) (response.Group, error)
	GetGroup(context.Context, string) (response.Group, error)
	UpdateGroup(context.Context, request.Group) (response.Group, error)
	DeleteGroup(context.Context, string) error
	ListGroups(context.Context, int32, int32) ([]response.Group, error)
	GetGroupStudents(ctx context.Context, groupID string) ([]response.Student, error)
	GetGroupStudentIDs(ctx context.Context, groupID string) ([]string, error)
}

type TeacherServiceClient interface {
	RegisterTeacher(context.Context, request.RegisterTeacherRequest) (response.Teacher, error)
	GetTeacher(context.Context, request.GetTeacherRequest) (response.Teacher, error)
	DeleteTeacher(context.Context, string) error
	ListTeachers(context.Context, int32, int32) ([]response.Teacher, error)

	CreateSubject(context.Context, request.CreateSubjectRequest) (response.Subject, error)
	GetSubject(context.Context, string) (response.Subject, error)
	DeleteSubject(context.Context, string) error
	ListSubjects(context.Context, int32, int32) ([]response.Subject, error)
}
