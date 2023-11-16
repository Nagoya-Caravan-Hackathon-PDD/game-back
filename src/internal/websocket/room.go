package websocket

type Rooms struct {
	Subscriptions []*Subscription
}

type Subscription struct {
	GameID string
	Client *Client
}

func (room *Rooms) AddSubscription(subscription *Subscription) {
	room.Subscriptions = append(room.Subscriptions, subscription)
}

func (room *Rooms) GetSubscription(gameID string) []Subscription {
	var subs []Subscription

	for _, sub := range room.Subscriptions {
		if sub.GameID == gameID {
			subs = append(subs, *sub)
		}
	}

	return subs
}

func (room *Rooms) All() []Subscription {
	var subs []Subscription

	for _, sub := range room.Subscriptions {
		subs = append(subs, *sub)
	}
	return subs
}

func (room *Rooms) Publish(gameID string, msg []byte) {

	subs := room.GetSubscription(gameID)

	for _, sub := range subs {
		err := sub.Client.Send(msg)
		if err != nil {
			return
		}
	}
}
