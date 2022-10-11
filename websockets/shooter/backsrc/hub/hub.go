package hub

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"shooter/backsrc/command"
	"shooter/backsrc/player"
)

func New(commands chan<- command.Command) *Hub {
	return &Hub{
		commands: commands,
	}
}

type Hub struct {
	commands chan<- command.Command
}

func (h *Hub) Read(p player.Player) {
	go func() {
		for {
			_, msg, err := p.Conn.ReadMessage()
			if err != nil {
				log.Printf("failed to read message from %s: %v\n", p.Name, err)
				continue
			}
			var cmd Command
			if err := json.Unmarshal(msg, &cmd); err != nil {
				log.Printf("failed to unmarshal message from %s: %v\n", p.Name, err)
				continue
			}
			cmd.payload = msg

			h.commands <- cmd
		}
	}()
}

func (h *Hub) Write(p player.Player, event []byte) error {
	return p.Conn.WriteMessage(websocket.TextMessage, event)
}

type Command struct {
	Command string `json:"command"`
	payload []byte
}

func (c Command) Name() string {
	return c.Command
}

func (c Command) Payload() []byte {
	return c.payload
}
