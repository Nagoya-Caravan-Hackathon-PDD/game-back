package websocket

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Nagoya-Caravan-Hackathon-PDD/game-back/src/domain"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type WsServer struct {
}

var upgrader = websocket.Upgrader{}
var rooms = Rooms{}

func (s *WsServer) ServeWs(c echo.Context) error {

	accountID := c.Param("account_id")
	log.Println(accountID)
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)

	if err != nil {
		c.Logger().Error(err)
	}

	defer ws.Close()

	client := &Client{Ws: ws}

	rooms.AddSubscription(&Subscription{AccountID: accountID, Client: client})

	data, err := json.Marshal(domain.Response{Message: fmt.Sprintf("Welcome %s", accountID)})
	if err != nil {
		c.Logger().Error(err)
	}

	rooms.Publish(data)
	for {
		_, msg, err := ws.ReadMessage()

		if err != nil {
			c.Logger().Error(err)
			break
		}

		rooms.Publish(msg)
	}
	return nil
}
