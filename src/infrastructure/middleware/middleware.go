package middleware

import (
	"github.com/Nagoya-Caravan-Hackathon-PDD/game-back/pkg/paseto"
	"github.com/labstack/echo/v4"
)

type middleware struct {
	maker paseto.Maker
}

type Middleware interface {
	Auth(next echo.HandlerFunc) echo.HandlerFunc
}

func NewMiddleware(maker paseto.Maker) Middleware {
	return &middleware{maker: maker}
}

const (
	WS_REQUEST = "ws_request"
)

type authtoken struct {
	Token string `param:"token"`
}

func (m *middleware) Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		token := ctx.Param("token")
		if len(token) == 0 {
			return echo.ErrBadRequest
		}

		paylaod, err := m.velifyPaseto(token)
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
