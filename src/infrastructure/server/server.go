package server

import "github.com/Nagoya-Caravan-Hackathon-PDD/game-back/src/infrastructure/socket"

type server struct {
}

func New() *server {
	return &server{}
}

func (s *server) Run() {
	e := socket.NewWsRouter()
	runWithGracefulShutdown(e)
}
