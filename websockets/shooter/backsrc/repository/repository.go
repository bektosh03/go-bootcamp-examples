package repository

import "shooter/backsrc/player"

func NewInMemory() *InMemory {
	return &InMemory{players: make(map[string]player.Player)}
}

type InMemory struct {
	players map[string]player.Player
}

func (m *InMemory) SavePlayer(p player.Player) {
	m.players[p.Name] = p
}

func (m *InMemory) ListPlayers() []player.Player {
	players := make([]player.Player, 0, len(m.players))
	for _, p := range m.players {
		players = append(players, p)
	}

	return players
}

func (m *InMemory) RemovePlayer(name string) {
	delete(m.players, name)
}
