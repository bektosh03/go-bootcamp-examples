package handler

import (
	"api-gateway/pkg/httperr"
	"api-gateway/request"
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

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
		httperr.InvalidJSON(w, r)
		return
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
		httperr.InvalidJSON(w, r)
		return
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
