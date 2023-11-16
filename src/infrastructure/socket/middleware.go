package socket

import (
	"encoding/base64"
	"encoding/json"

	"github.com/Nagoya-Caravan-Hackathon-PDD/game-back/src/domain"
	"github.com/labstack/echo/v4"
)

type middleware struct {
	e *echo.Echo
}

type Middleware interface {
	Auth() echo.MiddlewareFunc
}

func NewMiddleware(e *echo.Echo) Middleware {
	return &middleware{e: e}
}

const (
	WS_REQUEST = "ws_request"
)

func (m *middleware) Auth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			wsauth, err := m.decodeURI(ctx.Param("token"))
			if err != nil {
				return err
			}

			ctx.Set(WS_REQUEST, wsauth)
			return next(ctx)
		}
	}
}

// // URIデコードを一旦関数化
func (m *middleware) decodeURI(uri string) (wsauth *domain.WsAuthRequest, err error) {
	decodedUri, err := base64.URLEncoding.DecodeString(uri)
	if err != nil {
		return nil, echo.ErrBadRequest
	}
	if err = json.Unmarshal(decodedUri, &wsauth); err != nil {
		return nil, echo.ErrInternalServerError
	}
	return
}
