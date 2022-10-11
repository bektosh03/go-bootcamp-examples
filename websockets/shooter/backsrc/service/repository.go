package service

import (
	"shooter/backsrc/match"
	"shooter/backsrc/player"
)

type Repository interface {
	CreateMatch(match match.Match)
	GetMatch(id string) match.Match
	ShootPlayer() error
	UpdateMatch(match.Match)
	GetPlayer(name string) player.Player
	SavePlayer(p player.Player)
	ListPlayers() []player.Player
	RemovePlayer(name string)
}
