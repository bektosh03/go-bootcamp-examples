package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	h := NewHandler(NewServices("localhost:8001"))

	r := chi.NewRouter()

	r.Post("/teacher", h.RegisterTeacher)
	r.Get("/teacher/{id}", h.GetTeacher)
	r.Post("/subject", h.CreateSubject)
	r.Get("/subject/{id}", h.GetSubject)

	http.ListenAndServe(":8080", r)
}
