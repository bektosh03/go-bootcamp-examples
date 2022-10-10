package service

import "shooter/backsrc/player"

type Repository interface {
	SavePlayer(p player.Player)
	ListPlayers() []player.Player
	RemovePlayer(name string)
}
