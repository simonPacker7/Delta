package entities

type GameStatus string

const (
	GameStatusWaiting GameStatus = "waiting"
	GameStatusReady   GameStatus = "ready"
	GameStatusActive  GameStatus = "active"
	GameStatusEnded   GameStatus = "completed"
)

type GameType string

const (
	GameTypeOnline  GameType = "online"
	GameTypePrivate GameType = "private"
)

type GameMove struct {
	PlayerID   string `json:"playerId"`
	PlayerName string `json:"playerName"`
	Word       string `json:"word"`
	Timestamp  int64  `json:"timestamp"`
}

type Game struct {
	ID             string     `json:"id" redis:"id"`
	Type           GameType   `json:"type" redis:"type"`
	Status         GameStatus `json:"status" redis:"status"`
	JoinCode       string     `json:"joinCode" redis:"join_code"`
	Player1ID      string     `json:"player1Id" redis:"player1_id"`
	Player1Name    string     `json:"player1Name" redis:"player1_name"`
	Player2ID      string     `json:"player2Id" redis:"player2_id"`
	Player2Name    string     `json:"player2Name" redis:"player2_name"`
	CurrentWord    string     `json:"currentWord" redis:"current_word"`
	CurrentTurnID  string     `json:"currentTurnId" redis:"current_turn_id"`
	WinnerID       string     `json:"winnerId" redis:"winner_id"`
	WinReason      string     `json:"winReason" redis:"win_reason"`
	ConnectedCount int        `json:"connectedCount" redis:"connected_count"`
	CreatedAt      int64      `json:"createdAt" redis:"created_at"`
}

type FindGameResponse struct {
	Status string `json:"status"` // "waiting" or "matched"
	GameID string `json:"gameId"`
}

type CreatePrivateGameResponse struct {
	GameID   string `json:"gameId"`
	JoinCode string `json:"joinCode"`
}

type JoinPrivateGameInput struct {
	JoinCode string `json:"joinCode"`
}

type JoinPrivateGameResponse struct {
	Status string `json:"status"` // "matched" or "error"
	GameID string `json:"gameId"`
}
