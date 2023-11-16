package socket

import (
	"github.com/Nagoya-Caravan-Hackathon-PDD/game-back/cmd/config"
	"github.com/Nagoya-Caravan-Hackathon-PDD/game-back/pkg/paseto"
	"github.com/Nagoya-Caravan-Hackathon-PDD/game-back/src/application/websocket"
	"github.com/Nagoya-Caravan-Hackathon-PDD/game-back/src/infrastructure/middleware"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func NewWsRouter() *echo.Echo {
	e := echo.New()
	s := &websocket.WsServer{}

	maker, err := paseto.NewPasetoMaker(config.Config.Paseto.SecretKey)
	if err != nil {
		panic(err)
	}

	e.Use(middleware.NewMiddleware(e, maker).Auth)
	e.GET("/game/:token", s.ServeWs)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	return e
}
