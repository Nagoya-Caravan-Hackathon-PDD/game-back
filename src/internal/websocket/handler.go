package websocket

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Nagoya-Caravan-Hackathon-PDD/game-back/cmd/config"
	"github.com/Nagoya-Caravan-Hackathon-PDD/game-back/pkg/paseto"
	"github.com/Nagoya-Caravan-Hackathon-PDD/game-back/src/domain"
	"github.com/Nagoya-Caravan-Hackathon-PDD/game-back/src/infrastructure/middleware"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type WsServer struct {
	maker paseto.Maker
}

var upgrader = websocket.Upgrader{}
var rooms = Rooms{}

func (s *WsServer) ServeWs(c echo.Context) error {

	paylaod := c.Get(middleware.WS_REQUEST).(*paseto.Payload)

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)

	if err != nil {
		c.Logger().Error(err)
	}

	defer ws.Close()

	client := &Client{Ws: ws}

	rooms.AddSubscription(&Subscription{GameID: paylaod.GameID, Client: client})

	msg, err := json.Marshal(fmt.Sprintf("%s joined room", paylaod.OwnerID))
	if err != nil {
		return echo.ErrInternalServerError
	}

	rooms.Publish(paylaod.GameID, msg)

	for {
		if err != nil {
			c.Logger().Error(err)
			break
		}
	}
	return nil
}

func (s *WsServer) NotificateController(ctx echo.Context) error {
	var notificate domain.NotificateRequest
	if err := ctx.Bind(&notificate); err != nil {
		return echo.ErrBadRequest
	}

	payload := ctx.Get(middleware.WS_REQUEST).(*paseto.Payload)
	if payload.OwnerID != config.Config.Server.AdminID {
		return echo.ErrForbidden
	}

	switch notificate.MessageType {
	case domain.FlagStart:
		msg, err := json.Marshal(domain.FlagStart)
		if err != nil {
			return echo.ErrInternalServerError
		}

		rooms.Publish(payload.GameID, msg)

		for i := notificate.Time; i > 0; i-- {
			msg, err := json.Marshal(fmt.Sprintf("%d", i))
			if err != nil {
				return echo.ErrInternalServerError
			}

			rooms.Publish(payload.GameID, msg)
			time.Sleep(1 * time.Second)
		}

	case domain.FlagTurn:
		msg, err := json.Marshal(domain.FlagTurn)
		if err != nil {
			return echo.ErrInternalServerError
		}

		rooms.Publish(payload.GameID, msg)

		for i := notificate.Time; i > 0; i-- {
			msg, err := json.Marshal(fmt.Sprintf("%d", i))
			if err != nil {
				return echo.ErrInternalServerError
			}

			rooms.Publish(payload.GameID, msg)
			time.Sleep(1 * time.Second)
		}
	case domain.FlagResult:
		msg, err := json.Marshal(domain.FlagResult)
		if err != nil {
			return echo.ErrInternalServerError
		}
		rooms.Publish(payload.GameID, msg)

	case domain.FlagEnd:
		msg, err := json.Marshal(domain.FlagEnd)
		if err != nil {
			return echo.ErrInternalServerError
		}
		rooms.Publish(payload.GameID, msg)
	}

	return ctx.JSON(http.StatusOK, nil)
}
