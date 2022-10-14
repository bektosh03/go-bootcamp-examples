package repository

import (
	"shooter/backsrc/match"
	"shooter/backsrc/player"
	"sync"
)

func NewInMemory() *InMemory {
	return &InMemory{
		players: make(map[string]player.Player),
		matches: make(map[string]match.Match),
		mu:      &sync.Mutex{},
	}
}

type InMemory struct {
	players map[string]player.Player
	matches map[string]match.Match
	mu      *sync.Mutex
}

func (r *InMemory) RemoveMatch(id string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.matches, id)
}

func (m *InMemory) CreateMatch(match match.Match) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.matches[match.ID] = match
}

func (m *InMemory) GetMatch(id string) match.Match {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.matches[id]
}

func (m *InMemory) UpdateMatch(match match.Match) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.matches[match.ID] = match
}

func (m *InMemory) SavePlayer(p player.Player) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.players[p.Name] = p
}

func (m *InMemory) ListPlayers() []player.Player {
	m.mu.Lock()
	defer m.mu.Unlock()
	players := make([]player.Player, 0, len(m.players))
	for _, p := range m.players {
		players = append(players, p)
	}

	return players
}

func (m *InMemory) RemovePlayer(name string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.players, name)
}

func (m *InMemory) GetPlayer(name string) player.Player {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.players[name]
}
