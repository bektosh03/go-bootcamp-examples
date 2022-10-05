package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func setupRouter(r chi.Router, s Service, h *Hub) {
	r.Use(middleware.Logger)
	r.Post("/events", AddEventHandler(s))
	r.Get("/ws", WebSocketHandler(h))
}
