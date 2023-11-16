package domain

type MessageType string

const (
	FlagStart MessageType = "start"
	FlagEnd   MessageType = "end"
)

type NotificateRequest struct {
	Token       string      `param:"token"`
	MessageType MessageType `json:"message_type"`
	Time        int         `json:"time"`
}
