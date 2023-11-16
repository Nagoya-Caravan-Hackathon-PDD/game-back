package main

import (
	"github.com/Nagoya-Caravan-Hackathon-PDD/game-back/cmd/config"
	"github.com/Nagoya-Caravan-Hackathon-PDD/game-back/src/infrastructure/server"
)

func main() {
	config.LoadEnv()
	server.New().Run()
}
