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
		httperr.InvalidJSON(w, r)
		return
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
		httperr.InvalidJSON(w, r)
		return
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
