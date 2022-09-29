package handler

import (
	"api-gateway/pkg/auth"
	"api-gateway/pkg/httperr"
	"api-gateway/pkg/producer"
	"api-gateway/request"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// RegisterTeacher creates a new teacher
func (h Handler) RegisterTeacher(w http.ResponseWriter, r *http.Request) {
	var req request.RegisterTeacherRequest
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		httperr.InvalidJSON(w, r)
		return
	}

	teacher, err := h.service.Teacher.RegisterTeacher(context.Background(), req)
	if err != nil {
		httperr.Handle(w, r, err)
		return
	}

	if err = h.producer.Produce(producer.RegisteredEvent{
		Email:    teacher.Email,
		FullName: fmt.Sprintf("%s %s", teacher.FirstName, teacher.LastName),
		For:      producer.EventForTeacher,
	}); err != nil {
		httperr.Handle(w, r, err)
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, render.M{
		"teacher": teacher,
		"token":   auth.NewJWT(teacher.ID),
	})
}

// GetTeacher fetches teacher's data from database by teacherID
func (h Handler) GetTeacher(w http.ResponseWriter, r *http.Request) {
	var req request.GetTeacherRequest
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		httperr.InvalidJSON(w, r)
		return
	}

	teacher, err := h.service.Teacher.GetTeacher(context.Background(), req)
	if err != nil {
		httperr.Handle(w, r, err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, teacher)
}

// DeleteTeacher deletes teacher by ID
func (h Handler) DeleteTeacher(w http.ResponseWriter, r *http.Request) {
	teacherID := chi.URLParam(r, "id")

	err := h.service.Teacher.DeleteTeacher(context.Background(), teacherID)
	if err != nil {
		httperr.Handle(w, r, err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, render.M{
		"ok":      true,
		"message": "teacher deleted successfully",
	})
}

// ListTeachers fetches list of teachers
func (h Handler) ListTeachers(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		httperr.BadRequest(w, r, err.Error())
		return
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		httperr.BadRequest(w, r, err.Error())
		return
	}

	res, err := h.service.Teacher.ListTeachers(context.Background(), int32(page), int32(limit))
	if err != nil {
		httperr.InternalError(w, r, err.Error())
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, res)
}
