package gameService

import (
	"errors"
	"time"

	"github.com/simonPacker7/Delta/backend/shared/entities"
	"github.com/simonPacker7/Delta/backend/shared/redisclient"
	wordService "github.com/simonPacker7/Delta/backend/worker/services/word"
)

type Service struct {
	redisClient *redisclient.RedisClient
	wordService *wordService.Service
}

func NewService(r *redisclient.RedisClient, w *wordService.Service) *Service {
	return &Service{
		redisClient: r,
		wordService: w,
	}
}

// FindGame implements the matchmaking flow using atomic Redis operations:
// 1. Atomically try to pop and join a waiting game from the queue
// 2. If no valid game found, create a new game and add to queue
func (s *Service) FindGame(playerID string, playerName string) (entities.FindGameResponse, error) {
	startWord := s.wordService.GetRandomStartWord()

	// Atomically try to pop and join a game
	gameID, _, matched, err := s.redisClient.AtomicPopAndJoinGame(playerID, playerName, startWord)
	if err != nil {
		return entities.FindGameResponse{}, err
	}

	if matched {
		return entities.FindGameResponse{
			Status: "matched",
			GameID: gameID,
		}, nil
	}

	// No valid game found - create a new one
	return s.createNewGame(playerID, playerName)
}

func (s *Service) createNewGame(playerID string, playerName string) (entities.FindGameResponse, error) {
	gameID := redisclient.GenerateId()

	game := entities.Game{
		ID:             gameID,
		Type:           entities.GameTypeOnline,
		Status:         entities.GameStatusWaiting,
		JoinCode:       "",
		Player1ID:      playerID,
		Player1Name:    playerName,
		Player2ID:      "",
		Player2Name:    "",
		CurrentWord:    "",
		CurrentTurnID:  "",
		ConnectedCount: 0,
		CreatedAt:      time.Now().UnixMilli(),
	}

	err := s.redisClient.CreateGame(game)
	if err != nil {
		return entities.FindGameResponse{}, err
	}

	// Add to matchmaking queue
	err = s.redisClient.PushToMatchmakingQueue(gameID)
	if err != nil {
		return entities.FindGameResponse{}, err
	}

	return entities.FindGameResponse{
		Status: "waiting",
		GameID: gameID,
	}, nil
}

// GetGame retrieves the current game state
func (s *Service) GetGame(gameID string) (entities.Game, error) {
	game, err := s.redisClient.GetGame(gameID)
	if err != nil {
		return entities.Game{}, err
	}
	if game.ID == "" {
		return entities.Game{}, errors.New("game not found")
	}
	return game, nil
}

// CreatePrivateGame creates a new private game with a join code
func (s *Service) CreatePrivateGame(playerID string, playerName string) (entities.CreatePrivateGameResponse, error) {
	gameID := redisclient.GenerateId()
	joinCode := redisclient.GenerateGameCode()

	game := entities.Game{
		ID:             gameID,
		Type:           entities.GameTypePrivate,
		Status:         entities.GameStatusWaiting,
		JoinCode:       joinCode,
		Player1ID:      playerID,
		Player1Name:    playerName,
		Player2ID:      "",
		Player2Name:    "",
		CurrentWord:    "",
		CurrentTurnID:  "",
		ConnectedCount: 0,
		CreatedAt:      time.Now().UnixMilli(),
	}

	err := s.redisClient.CreateGame(game)
	if err != nil {
		return entities.CreatePrivateGameResponse{}, err
	}

	// Map join code to game ID for lookup
	err = s.redisClient.SetPrivateGameCode(joinCode, gameID)
	if err != nil {
		return entities.CreatePrivateGameResponse{}, err
	}

	return entities.CreatePrivateGameResponse{
		GameID:   gameID,
		JoinCode: joinCode,
	}, nil
}

// JoinPrivateGame atomically joins a private game using a join code
// Uses Lua script to prevent race conditions when two players try to join simultaneously
func (s *Service) JoinPrivateGame(playerID string, playerName string, joinCode string) (entities.JoinPrivateGameResponse, error) {
	startWord := s.wordService.GetRandomStartWord()

	// Atomically validate and join the game
	gameID, _, err := s.redisClient.AtomicJoinPrivateGame(joinCode, playerID, playerName, startWord)
	if err != nil {
		// Convert atomic operation errors to user-friendly messages
		if atomicErr, ok := err.(*redisclient.AtomicOperationError); ok {
			switch atomicErr.Message {
			case "invalid_code":
				return entities.JoinPrivateGameResponse{}, errors.New("invalid join code")
			case "game_not_available":
				return entities.JoinPrivateGameResponse{}, errors.New("game is no longer available")
			case "cannot_join_own_game":
				return entities.JoinPrivateGameResponse{}, errors.New("cannot join your own game")
			}
		}
		return entities.JoinPrivateGameResponse{}, err
	}

	return entities.JoinPrivateGameResponse{
		Status: "matched",
		GameID: gameID,
	}, nil
}
