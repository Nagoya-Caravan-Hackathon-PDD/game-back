package domain

type WsAuthRequest struct {
	AuthToken string `json:"auth_token"`
	RoomID    string `json:"room_id"`
}
