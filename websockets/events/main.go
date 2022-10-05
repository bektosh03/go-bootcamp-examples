package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
	"net/http"
)

var (
	upgrader = websocket.Upgrader{}
)

func main() {
	repo := NewInMemoryRepository()
	conflicts := make(chan ConflictMessage)
	s := NewService(repo, conflicts)
	h := NewHub(conflicts)

	r := chi.NewRouter()
	setupRouter(r, s, h)

	h.Run()
	http.ListenAndServe("localhost:8080", r)
}
