package main

import (
	"log"

	"github.com/Nagoya-Caravan-Hackathon-PDD/game-back/cmd/config"
	"github.com/Nagoya-Caravan-Hackathon-PDD/game-back/pkg/paseto"
	"github.com/Nagoya-Caravan-Hackathon-PDD/game-back/src/infrastructure/server"
)

func main() {
	config.LoadEnv()

	maker, err := paseto.NewPasetoMaker(config.Config.Paseto.SecretKey)
	if err != nil {
		log.Fatal("failed to create paseto maker :", err)
	}
	server.New().Run(maker)
}
