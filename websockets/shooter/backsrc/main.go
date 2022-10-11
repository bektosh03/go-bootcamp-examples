package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"net/http"
	"shooter/backsrc/command"
	"shooter/backsrc/handler"
	"shooter/backsrc/hub"
	"shooter/backsrc/player"
	"shooter/backsrc/repository"
	"shooter/backsrc/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

var (
	upgrader = websocket.Upgrader{}
)

func main() {
	r := chi.NewRouter()

	commands := make(chan command.Command)
	repo := repository.NewInMemory()
	s := service.New(repo)
	h := hub.New(commands)
	websocketHandler := handler.NewWebsocketHandler(s, h, commands)
	websocketHandler.Run()

	r.Use(cors.AllowAll().Handler)

	r.Get("/register", func(w http.ResponseWriter, r *http.Request) {
		var p player.Player
		p.Name = r.URL.Query().Get("name")

		serverWS(w, r, s, h, p)
	})

	http.ListenAndServe(":8000", r)
}

func serverWS(w http.ResponseWriter, r *http.Request, s service.Service, h *hub.Hub, p player.Player) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}

	p.Conn = conn

	s.CreatePlayer(p)

	h.Read(p)

	availablePlayers, err := json.Marshal(s.AvailablePlayers())
	if err != nil {
		panic(err)
	}
	if err = h.Write(p, availablePlayers); err != nil {
		panic(err)
	}
}

type Request struct {
	Name string `json:"name"`
}
