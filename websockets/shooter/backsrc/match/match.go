package match

import "shooter/backsrc/player"

type Match struct {
	ID           string
	Player1      player.Player
	Player2      player.Player
	Player1Ready bool
	Player2Ready bool
}
