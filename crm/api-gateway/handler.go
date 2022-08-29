package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func NewHandler(services Services) Handler {
	return Handler{
		services: services,
	}
}

type Handler struct {
	services Services
}

func (h Handler) CreateSubject(w http.ResponseWriter, r *http.Request) {
	var request CreateSubjectRequest
	if err := render.DecodeJSON(r.Body, &request); err != nil {
		panic(err)
	}

	subject, err := h.services.TeacherService.CreateSubject(context.Background(), request)
	if err != nil {
		panic(err)
	}

	render.JSON(w, r, subject)
}

func (h Handler) GetSubject(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	subject, err := h.services.TeacherService.GetSubject(context.Background(), id)
	if err != nil {
		panic(err)
	}

	render.JSON(w, r, subject)
}

func (h Handler) RegisterTeacher(w http.ResponseWriter, r *http.Request) {
	var request RegisterTeacherRequest
	if err := render.DecodeJSON(r.Body, &request); err != nil {
		panic(err)
	}

	teacher, err := h.services.TeacherService.RegisterTeacher(context.Background(), request)
	if err != nil {
		panic(err)
	}

	// STEP 3 - compose (add) data from services to single response
	render.JSON(w, r, teacher)
}

func (h Handler) GetTeacher(w http.ResponseWriter, r *http.Request) {
	teacherID := chi.URLParam(r, "id")

	teacher, err := h.services.TeacherService.GetTeacher(context.Background(), teacherID)
	if err != nil {
		panic(err)
	}

	// STEP 3 - compose (add) data from services to single response
	render.JSON(w, r, teacher)
}
