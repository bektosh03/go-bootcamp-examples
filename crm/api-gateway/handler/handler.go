package handler

import (
	"api-gateway/request"
	"api-gateway/response"
	"api-gateway/service"
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func New(svc service.Service) Handler {
	return Handler{
		service: svc,
	}
}

type Handler struct {
	service service.Service
}

func (h Handler) RegisterSchedule(w http.ResponseWriter, r *http.Request) {
	var req request.CreateScheduleRequest
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		panic(err)
	}
	schedule, err := h.service.Schedule.RegisterSchedule(context.Background(), req)
	if err != nil {
		panic(err)
	}
	render.JSON(w, r, schedule)
}

func (h Handler) UpdateSchedule(w http.ResponseWriter, r *http.Request) {
	var req request.Schedule
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		panic(err)
	}

	res, err := h.service.Schedule.UpdateSchedule(context.Background(), req)
	if err != nil {
		panic(err)
	}

	render.JSON(w, r, res)
}

func (h Handler) DeleteSchedule(w http.ResponseWriter, r *http.Request) {
	scheduleID := chi.URLParam(r, "id")

	if err := h.service.Schedule.DeleteSchedule(context.Background(), scheduleID); err != nil {
		panic(err)
	}

	render.JSON(w, r, render.M{
		"ok": true,
	})
}

func (h Handler) GetFullScheduleForTeacher(w http.ResponseWriter, r *http.Request) {
	teacherID := chi.URLParam(r, "id")

	schedules, err := h.service.Schedule.GetFullScheduleForTeacher(context.Background(), teacherID)
	if err != nil {
		panic(err)
	}

	populatedSchedules := make([]response.PopulatedSchedule, 0, len(schedules))
	for _, sch := range schedules {
		group, err := h.service.Student.GetGroup(context.Background(), sch.GroupId)
		if err != nil {
			panic(err)
		}

		subject, err := h.service.Teacher.GetSubject(context.Background(), sch.SubjectID)
		if err != nil {
			panic(err)
		}

		teacher, err := h.service.Teacher.GetTeacher(context.Background(), sch.TeacherID)
		if err != nil {
			panic(err)
		}

		populatedSchedules = append(populatedSchedules, response.PopulatedSchedule{
			ID:           sch.ID,
			Group:        group,
			Subject:      subject,
			Teacher:      teacher,
			Weekday:      sch.WeekDay,
			LessonNumber: sch.LessonNumber,
		})
	}

	render.JSON(w, r, populatedSchedules)
}

func (h Handler) GetScheduleById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	schedule, err := h.service.Schedule.GetSchedule(context.Background(), id)
	if err != nil {
		panic(err)
	}
	render.JSON(w, r, schedule)
}

func (h Handler) ListGroups(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		panic(err)
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		panic(err)
	}

	groups, err := h.service.Student.ListGroups(context.Background(), int32(page), int32(limit))
	if err != nil {
		panic(err)
	}

	render.JSON(w, r, groups)
}

func (h Handler) DeleteGroup(w http.ResponseWriter, r *http.Request) {
	groupID := chi.URLParam(r, "id")

	if err := h.service.Student.DeleteGroup(context.Background(), groupID); err != nil {
		panic(err)
	}

	render.JSON(w, r, render.M{
		"ok": "deleted",
	})
}

func (h Handler) UpdateGroup(w http.ResponseWriter, r *http.Request) {
	var req request.Group
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		panic(err)
	}

	updatedGroup, err := h.service.Student.UpdateGroup(context.Background(), req)
	if err != nil {
		panic(err)
	}

	render.JSON(w, r, updatedGroup)
}

func (h Handler) CreateGroup(w http.ResponseWriter, r *http.Request) {
	var req request.CreateGroupRequest
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		panic(err)
	}

	group, err := h.service.Student.CreateGroup(context.Background(), req)
	if err != nil {
		panic(err)
	}

	render.JSON(w, r, group)
}

func (h Handler) GetGroup(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	group, err := h.service.Student.GetGroup(context.Background(), id)
	if err != nil {
		panic(err)
	}

	render.JSON(w, r, group)
}

func (h Handler) ListStudents(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		panic(err)
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		panic(err)
	}
	students, err := h.service.Student.ListStudents(context.Background(), int32(page), int32(limit))
	if err != nil {
		panic(err)
	}

	render.JSON(w, r, students)
}

func (h Handler) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	studentID := chi.URLParam(r, "id")

	if err := h.service.Student.DeleteStudent(context.Background(), studentID); err != nil {
		panic(err)
	}

	render.JSON(w, r, render.M{
		"ok": "deleted",
	})
}

func (h Handler) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	var req request.Student
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		panic(err)
	}

	updatedStudent, err := h.service.Student.UpdateStudent(context.Background(), req)
	if err != nil {
		panic(err)
	}

	render.JSON(w, r, updatedStudent)
}

func (h Handler) RegisterStudent(w http.ResponseWriter, r *http.Request) {
	var req request.RegisterStudentRequest
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		panic(err)
	}

	student, err := h.service.Student.RegisterStudent(context.Background(), req)
	if err != nil {
		panic(err)
	}

	render.JSON(w, r, student)
}

func (h Handler) GetStudent(w http.ResponseWriter, r *http.Request) {
	studentID := chi.URLParam(r, "id")

	student, err := h.service.Student.GetStudent(context.Background(), studentID)
	if err != nil {
		panic(err)
	}

	render.JSON(w, r, student)
}

func (h Handler) CreateSubject(w http.ResponseWriter, r *http.Request) {
	var req request.CreateSubjectRequest
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		panic(err)
	}

	subject, err := h.service.Teacher.CreateSubject(context.Background(), req)
	if err != nil {
		panic(err)
	}

	render.JSON(w, r, subject)
}

func (h Handler) GetSubject(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	subject, err := h.service.Teacher.GetSubject(context.Background(), id)
	if err != nil {
		panic(err)
	}

	render.JSON(w, r, subject)
}

func (h Handler) RegisterTeacher(w http.ResponseWriter, r *http.Request) {
	var req request.RegisterTeacherRequest
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		panic(err)
	}

	teacher, err := h.service.Teacher.RegisterTeacher(context.Background(), req)
	if err != nil {
		panic(err)
	}

	render.JSON(w, r, teacher)
}

func (h Handler) GetTeacher(w http.ResponseWriter, r *http.Request) {
	teacherID := chi.URLParam(r, "id")

	teacher, err := h.service.Teacher.GetTeacher(context.Background(), teacherID)
	if err != nil {
		panic(err)
	}

	render.JSON(w, r, teacher)
}
