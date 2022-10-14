package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"shooter/backsrc/command"
	"shooter/backsrc/hub"
	"shooter/backsrc/match"
	"shooter/backsrc/player"
	"shooter/backsrc/service"

	"github.com/google/uuid"
)

const (
	EventPlayerJoined       = "player_joined"
	EventNewAvailablePlayer = "new_available_player"
	EventMatchCreated       = "match_created"
	EventMatchStarted       = "match_started"
	EventYouWon             = "you_won"
	EventYouLost            = "you_lost"
	EventGameOver           = "game_over"
	EventPlayerUpdate       = "player_update"
)

func NewWebsocketHandler(s service.Service, h *hub.Hub, commands <-chan command.Command) *WebsocketHandler {
	return &WebsocketHandler{
		s:        s,
		hub:      h,
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
			fmt.Printf("received command %s: %s\n", cmd.Name(), cmd.Payload())
			switch cmd.Name() {
			case command.WaitForOpponent:
				var payload WaitForOpponentPayload
				if err := json.Unmarshal(cmd.Payload(), &payload); err != nil {
					log.Printf("failed to unmarshal %s payload: %v\n", cmd.Name(), err)
					continue
				}
				h.s.WaitForSomeone(payload.Player.Name)
				event := Event{
					Name:   EventNewAvailablePlayer,
					Player: payload.Player,
				}

				h.notifyOthers(payload.Player.Name, event.Marshal())

			case command.Play:
				var payload PlayPayload
				if err := json.Unmarshal(cmd.Payload(), &payload); err != nil {
					log.Printf("failed to unmarshal %s payload: %v\n", cmd.Name(), err)
					continue
				}
				player := h.s.GetPlayer(payload.Player.Name)
				rival := h.s.GetPlayer(payload.Rival.Name)
				event := Event{
					Name:   EventPlayerJoined,
					Player: payload.Player,
				}

				if err := h.hub.Write(rival, event.Marshal()); err != nil {
					log.Printf("failed to write event to %s: %v\n", rival.Name, err)
					continue
				}

				m := match.Match{
					ID:           uuid.NewString(),
					Player1:      player,
					Player2:      rival,
					Player1Ready: false,
					Player2Ready: false,
				}
				h.s.CreateMatch(m)

				event = Event{
					Name: EventMatchCreated,
					Metadata: map[string]interface{}{
						"match_id": m.ID,
					},
				}

				h.hub.Write(player, event.Marshal())
				h.hub.Write(rival, event.Marshal())

			case command.Start:
				var payload StartPayload
				if err := json.Unmarshal(cmd.Payload(), &payload); err != nil {
					log.Printf("failed to unmarshal %s payload: %v\n", cmd.Name(), err)
					continue
				}

				m := h.s.GetMatch(payload.MatchID)
				if m.Player1.Name == payload.Player.Name {
					m.Player1Ready = true
				}
				if m.Player2.Name == payload.Player.Name {
					m.Player2Ready = true
				}

				h.s.UpdateMatch(m)

				if m.Player1Ready && m.Player2Ready {
					event := Event{
						Name: EventMatchStarted,
						Metadata: map[string]interface{}{
							"match_id": m.ID,
						},
					}
					h.hub.Write(m.Player1, event.Marshal())
					h.hub.Write(m.Player2, event.Marshal())
				}
			case command.Shoot:
				var payload ShootPayload
				err := json.Unmarshal(cmd.Payload(), &payload)
				if err != nil {
					log.Printf("failed to unmarshal %s payload: %v\n", cmd.Name(), err)
					continue
				}
				m := h.s.GetMatch(payload.MatchID)
				var shooter, shotPlayer player.Player
				if payload.Player.Name == m.Player1.Name {
					shooter = m.Player1
					shotPlayer = m.Player2
				} else {
					shooter = m.Player2
					shotPlayer = m.Player1
				}
				shotPlayer.Health -= 10
				h.s.SavePlayer(shotPlayer)
				m.Player1 = shooter
				m.Player2 = shotPlayer
				h.s.UpdateMatch(m)

				if shotPlayer.Health <= 0 {
					event := Event{
						Name:   EventYouWon,
						Player: shooter,
						Metadata: map[string]interface{}{
							"match_id": m.ID,
						},
					}
					h.hub.Write(shooter, event.Marshal())
					// TODO send player lost event

					event = Event{
						Name: EventGameOver,
						Metadata: map[string]interface{}{
							"match_id": payload.MatchID,
						},
					}
					h.hub.Write(shooter, event.Marshal())
					h.hub.Write(shotPlayer, event.Marshal())
					h.s.RemoveMatch(m.ID)
					h.s.RemovePlayer(shotPlayer.Name)
				}

				h.hub.Write(shotPlayer, Event{
					Name:     EventPlayerUpdate,
					Player:   shotPlayer,
					Metadata: nil,
				}.Marshal())

			default:
				panic("no such command")
			}
		}
	}()
}

func (h *WebsocketHandler) notifyOthers(self string, event []byte) {
	players := h.s.AllPlayers()
	for _, p := range players {
		if p.Name == self {
			continue
		}

		if err := h.hub.Write(p, event); err != nil {
			log.Printf("failed to write in notifyAll: %v\n", err)
			continue
		}
	}
}

type ShootPayload struct {
	MatchID string        `json:"match_id"`
	Player  player.Player `json:"player"`
}

type WaitForOpponentPayload struct {
	Player player.Player `json:"player"`
}

type PlayPayload struct {
	Player player.Player `json:"player"`
	Rival  player.Player `json:"rival"`
}

type StartPayload struct {
	MatchID string        `json:"match_id"`
	Player  player.Player `json:"player"`
}

type Event struct {
	Name     string                 `json:"name"`
	Player   player.Player          `json:"player"`
	Metadata map[string]interface{} `json:"metadata"`
}

func (e Event) Marshal() []byte {
	data, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}

	return data
}
