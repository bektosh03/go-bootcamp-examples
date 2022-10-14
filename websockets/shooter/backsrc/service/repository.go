package service

import (
	"shooter/backsrc/match"
	"shooter/backsrc/player"
)

type Repository interface {
	RemoveMatch(id string)
	CreateMatch(match match.Match)
	GetMatch(id string) match.Match
	UpdateMatch(match.Match)
	GetPlayer(name string) player.Player
	SavePlayer(p player.Player)
	ListPlayers() []player.Player
	RemovePlayer(name string)
}