package server

import (
	"github.com/Nagoya-Caravan-Hackathon-PDD/game-back/pkg/paseto"
	"github.com/Nagoya-Caravan-Hackathon-PDD/game-back/src/infrastructure/socket"
)

type server struct {
}

func New() *server {
	return &server{}
}

func (s *server) Run(maker paseto.Maker) {
	e := socket.NewWsRouter(maker)
	runWithGracefulShutdown(e)
}
