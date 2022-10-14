package service

import (
	"fmt"
	"shooter/backsrc/match"
	"shooter/backsrc/player"
)

func New(repo Repository) Service {
	return Service{repo: repo}
}

type Service struct {
	repo Repository
}

func (s Service) RemovePlayer(name string) {
	s.repo.RemovePlayer(name)
}

func (s Service) RemoveMatch(id string) {
	s.repo.RemoveMatch(id)
}

func (s Service) SavePlayer(p player.Player) {
	s.repo.SavePlayer(p)
}

func (s Service) CreateMatch(match match.Match) {
	s.repo.CreateMatch(match)
}

func (s Service) GetMatch(id string) match.Match {
	return s.repo.GetMatch(id)
}

func (s Service) UpdateMatch(m match.Match) {
	s.repo.UpdateMatch(m)
}

func (s Service) GetPlayer(name string) player.Player {
	return s.repo.GetPlayer(name)
}
func (s Service) WaitForSomeone(name string) {
	p := s.GetPlayer(name)
	p.SetWaitingForOpponent(true)
	s.repo.SavePlayer(p)
}

func (s Service) CreatePlayer(p player.Player) {
	p.Health = 100
	s.repo.SavePlayer(p)
}

func (s Service) AvailablePlayers() []player.Player {
	availablePlayers := make([]player.Player, 0)
	players := s.repo.ListPlayers()
	fmt.Println("available players:", players)
	for _, p := range players {
		if p.IsWaitingForOpponent() {
			availablePlayers = append(availablePlayers, p)
		}
	}

	return availablePlayers
}

func (s Service) AllPlayers() []player.Player {
	return s.repo.ListPlayers()
}
