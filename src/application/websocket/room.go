package websocket

type Rooms struct {
	Subscriptions []*Subscription
}

type Subscription struct {
	AccountID string
	Client    *Client
}

func (room *Rooms) AddSubscription(subscription *Subscription) {
	room.Subscriptions = append(room.Subscriptions, subscription)
}

func (room *Rooms) GetSubscription(accountID string) []Subscription {
	var subs []Subscription

	for _, sub := range room.Subscriptions {
		if sub.AccountID == accountID {
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

func (room *Rooms) Publish(msg []byte) {

	subs := room.All()

	for _, sub := range subs {
		err := sub.Client.Send(msg)
		if err != nil {
			return
		}
	}
}
