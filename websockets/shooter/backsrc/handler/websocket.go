package handler

import (
	"encoding/json"
	"log"
	"shooter/backsrc/command"
	"shooter/backsrc/hub"
	"shooter/backsrc/player"
	"shooter/backsrc/service"
)

func NewWebsocketHandler(s service.Service, commands <-chan command.Command) *WebsocketHandler {
	return &WebsocketHandler{
		s:        s,
		commands: commands,
	}
}

type WebsocketHandler struct {
	s        service.Service
	hub      *hub.Hub
	commands <-chan command.Command
}

func (h *WebsocketHandler) Run() {
	go func() {
		for cmd := range h.commands {
			switch cmd.Name() {
			case command.WaitForOpponent:
				var payload WaitForOpponentPayload
				if err := json.Unmarshal(cmd.Payload(), &payload); err != nil {
					log.Printf("failed to unmarshal %s payload: %v\n", cmd.Name(), err)
					continue
				}
				h.s.WaitForSomeone(payload.Player)
			}
		}
	}()
}

type WaitForOpponentPayload struct {
	Player player.Player `json:"player"`
}
