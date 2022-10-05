package main

import (
	"fmt"
	"github.com/gorilla/websocket"
)

func NewHub(conflicts <-chan ConflictMessage) *Hub {
	return &Hub{
		clients:          make(map[string]*websocket.Conn),
		conflictMessages: conflicts,
	}
}

type Hub struct {
	clients          map[string]*websocket.Conn
	conflictMessages <-chan ConflictMessage
}

func (h *Hub) Run() {
	go func() {
		for conflictMsg := range h.conflictMessages {
			conn, ok := h.clients[conflictMsg.To]
			if !ok {
				fmt.Println("no such client:", conflictMsg.To)
				continue
			}

			if err := conn.WriteJSON(conflictMsg.Conflicts); err != nil {
				fmt.Println("failed to write json to websocket:", err)
				continue
			}
		}
	}()
}

func (h *Hub) AddClient(email string, conn *websocket.Conn) {
	h.clients[email] = conn
}
