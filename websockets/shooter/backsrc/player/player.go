package player

import (
	"github.com/gorilla/websocket"
)

type Player struct {
	Name                       string `json:"name"`
	Health                     int    `json:"health"`
	Conn                       *websocket.Conn
	isPlayerWaitingForOpponent bool
}

func (p *Player) SetWaitingForOpponent(b bool) {
	p.isPlayerWaitingForOpponent = b
}

func (p *Player) IsPlayerWaitingForOpponent() bool {
	return p.isPlayerWaitingForOpponent
}
