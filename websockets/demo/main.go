package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

type Event struct {
	ID       int                    `json:"id"`
	Metadata map[string]interface{} `json:"metadata"`
}

type MockEventsGenerator struct {
	counter      int
	tickDuration time.Duration
}

func (g *MockEventsGenerator) generate() Event {
	g.counter++
	return Event{
		ID: g.counter,
		Metadata: map[string]interface{}{
			"type":    "event",
			"message": "hello",
		},
	}
}

func (g *MockEventsGenerator) Events(ctx context.Context) <-chan Event {
	ch := make(chan Event)

	go func() {
		defer close(ch)

		for {
			select {
			case <-time.After(g.tickDuration):
				ch <- g.generate()
			case <-ctx.Done():
				return
			}
		}
	}()

	return ch
}

func main() {
	eventGenerator := MockEventsGenerator{
		counter:      0,
		tickDuration: time.Second * 2,
	}
	upgrader := websocket.Upgrader{}
	r := chi.NewRouter()

	r.Get("/ws", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*11)
		defer cancel()

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			panic(err)
		}
		defer func() {
			conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "the end"))
			conn.Close()
		}()

		for event := range eventGenerator.Events(ctx) {
			if err = conn.WriteJSON(event); err != nil {
				fmt.Println("failed to write json:", err)
				return
			}
		}
	})

	http.ListenAndServe("localhost:8080", r)
}
