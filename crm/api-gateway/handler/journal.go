package handler

import (
	"api-gateway/pkg/httperr"
	"api-gateway/request"
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func (h Handler) RegisterJournal(w http.ResponseWriter, r *http.Request) {
	var req request.CreateJournalRequest
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		httperr.InvalidJSON(w, r)
		return
	}

	journal, err := h.service.Journal.RegisterJournal(context.Background(), req)
	if err != nil {
		httperr.Handle(w, r, err)
	}
	render.JSON(w, r, journal)
}

func (h Handler) GetJournal(w http.ResponseWriter, r *http.Request) {
	journalId := chi.URLParam(r, "id")
	journal, err := h.service.Journal.GetJournal(context.Background(), journalId)
	if err != nil {
		panic(err)
	}
	render.JSON(w, r, journal)
}

func (h Handler) DeleteJournal(w http.ResponseWriter, r *http.Request) {
	journalId := chi.URLParam(r, "id")
	err := h.service.Journal.DeleteJournal(context.Background(), journalId)
	if err != nil {
		panic(err)
	}
	render.JSON(w, r, render.M{
		"ok": "deleted",
	})
}

func (h Handler) UpdateJournal(w http.ResponseWriter, r *http.Request) {
	var req request.Journal
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		httperr.InvalidJSON(w, r)
		return
	}
	res, err := h.service.Journal.UpdateJournal(context.Background(), req)
	if err != nil {
		panic(err)
	}
	render.JSON(w, r, res)
}
