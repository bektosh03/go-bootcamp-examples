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

// CreateSubject creates a new subject
func (h Handler) CreateSubject(w http.ResponseWriter, r *http.Request) {
	var req request.CreateSubjectRequest
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		httperr.InvalidJSON(w, r)
		return
	}

	subject, err := h.service.Teacher.CreateSubject(context.Background(), req)
	if err != nil {
		httperr.Handle(w, r, err)
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, subject)
}

// GetSubject fetches subject data from database by subjectID
func (h Handler) GetSubject(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	subject, err := h.service.Teacher.GetSubject(context.Background(), id)
	if err != nil {
		httperr.Handle(w, r, err)
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, subject)
}

// DeleteSubject deletes subject by subjectID
func (h Handler) DeleteSubject(w http.ResponseWriter, r *http.Request) {
	subjectID := chi.URLParam(r, "id")

	err := h.service.Teacher.DeleteSubject(context.Background(), subjectID)
	if err != nil {
		httperr.Handle(w, r, err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, render.M{
		"ok":      true,
		"message": "subject deleted successfully",
	})
}

// ListSubjects fetches list of subjects from database
func (h Handler) ListSubjects(w http.ResponseWriter, r *http.Request) {
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

	res, err := h.service.Teacher.ListSubjects(context.Background(), int32(page), int32(limit))
	if err != nil {
		httperr.InternalError(w, r, err.Error())
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, res)
}
