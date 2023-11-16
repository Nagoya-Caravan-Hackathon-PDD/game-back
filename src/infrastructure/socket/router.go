package socket

import (
	"github.com/Nagoya-Caravan-Hackathon-PDD/game-back/pkg/paseto"
	"github.com/Nagoya-Caravan-Hackathon-PDD/game-back/src/infrastructure/middleware"
	"github.com/Nagoya-Caravan-Hackathon-PDD/game-back/src/internal/websocket"
	"github.com/labstack/echo/v4"
)

func NewWsRouter(maker paseto.Maker) *echo.Echo {
	e := echo.New()
	s := &websocket.WsServer{}

	e.Use(middleware.NewMiddleware(maker).Auth)
	e.POST("/game_timer/:token", s.NotificateController)
	e.GET("/game_timer/:token", s.ServeWs)
	return e
}
