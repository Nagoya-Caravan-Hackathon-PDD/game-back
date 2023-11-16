package middleware

import (
	"log"

	"github.com/Nagoya-Caravan-Hackathon-PDD/game-back/pkg/paseto"
	"github.com/labstack/echo/v4"
)

type middleware struct {
	e     *echo.Echo
	maker paseto.Maker
}

type Middleware interface {
	Auth(next echo.HandlerFunc) echo.HandlerFunc
}

func NewMiddleware(e *echo.Echo, maker paseto.Maker) Middleware {
	return &middleware{e: e, maker: maker}
}

const (
	WS_REQUEST = "ws_request"
)

type authtoken struct {
	Token string `param:"token"`
}

func (m *middleware) Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var token authtoken
		if err := ctx.Bind(&token); err != nil {
			log.Println(err)
			return echo.ErrBadRequest
		}

		if len(token.Token) == 0 {
			return echo.ErrBadRequest
		}

		paylaod, err := m.velifyPaseto(token.Token)
		if err != nil {
			return err
		}

		ctx.Set(WS_REQUEST, paylaod)
		return next(ctx)
	}
}

// // URIデコードを一旦関数化
func (m *middleware) velifyPaseto(token string) (*paseto.Payload, error) {
	return m.maker.VerifyToken(token)
}
