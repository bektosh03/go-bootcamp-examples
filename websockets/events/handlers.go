package main

import (
	"encoding/json"
	"net/http"
)

func AddEventHandler(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var event Event
		if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
			panic(err)
		}

		if err := s.AddEvent(event); err != nil {
			panic(err)
		}

		w.WriteHeader(http.StatusOK)
	}
}

func WebSocketHandler(h *Hub) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := r.URL.Query().Get("email")
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			panic(err)
		}

		h.AddClient(email, conn)
	}
}
