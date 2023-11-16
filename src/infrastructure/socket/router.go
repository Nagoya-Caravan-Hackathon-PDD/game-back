package socket

import (
	"github.com/Nagoya-Caravan-Hackathon-PDD/game-back/src/application/websocket"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func NewWsRouter() *echo.Echo {
	e := echo.New()
	s := &websocket.WsServer{}

	e.Use(NewMiddleware(e).Auth())
	e.GET("/game/:token", s.ServeWs)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	return e
}
