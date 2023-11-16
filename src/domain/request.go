package domain

type NotificateType int

const (
	OrderNotiifcae NotificateType = iota
	ChatNotificate
	SystemNotificate
)

type NotificateRequest struct {
	Sender         string         `json:"sender"`
	ToAccountID    string         `json:"to_account_id"`
	NotificateType NotificateType `json:"notificate_type"`
	Detail         string         `json:"detail"`
}

type CommandSelection struct {
}
