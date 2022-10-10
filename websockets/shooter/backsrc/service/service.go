package service

import "shooter/backsrc/player"

func New(repo Repository) Service {
	return Service{repo: repo}
}

type Service struct {
	repo Repository
}

func (s Service) WaitForSomeone(p player.Player) {
	p.SetWaitingForOpponent(true)
	s.repo.SavePlayer(p)
}

func (s Service) CreatePlayer(p player.Player) {
	p.Health = 100
	s.repo.SavePlayer(p)
}

func (s Service) AvailablePlayers() []player.Player {
	players := s.repo.ListPlayers()
	for i := 0; i < len(players); i++ {
		if !players[i].IsPlayerWaitingForOpponent() {
			players = append(players[:i], players[i+1:]...)
		}
	}

	return players
}
