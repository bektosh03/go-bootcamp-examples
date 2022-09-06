package handler

import (
	"api-gateway/request"
	"api-gateway/service"
	"context"
	"net/http"

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
