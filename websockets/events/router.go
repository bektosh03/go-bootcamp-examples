package main

import (
	"github.com/go-chi/chi/v5"
)

func setupRouter(r chi.Router, s Service, h *Hub) {
	r.Post("/events", AddEventHandler(s))
	r.Get("/ws", WebSocketHandler(h))
}
