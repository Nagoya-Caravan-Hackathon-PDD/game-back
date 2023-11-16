package domain

type GameStatus struct {
	GameID string            `json:"game_id"`
	Status map[string]Stauts `json:"status"`
}

type Stauts struct {
	HP int `json:"hp"`
}
