package main

import (
	"context"
	"encoding/json"
	"log"
	"sync"

	"github.com/redis/go-redis/v9"
	"github.com/simonPacker7/Delta/backend/game-service/word"
	"github.com/simonPacker7/Delta/backend/shared/redisclient"
)

// ClientAction represents an incoming message from a client
type ClientAction struct {
	Action string `json:"action"` // "join_game", "leave_game", "submit_word"
	GameID string `json:"gameId"`
	Word   string `json:"word,omitempty"`
}

// GameMessage represents a message to broadcast to game clients
type GameMessage struct {
	Type    string      `json:"type"` // "game_ready", "game_start", "move", "game_ended", etc.
	GameID  string      `json:"gameId"`
	Payload interface{} `json:"payload,omitempty"`
}

type Hub struct {
	// Active clients mapped by GameID
	games map[string]map[*Client]bool

	// All connected clients mapped by UserID
	clients map[string]*Client

	// Channel for client actions that need processing
	actions chan *ClientActionRequest

	register chan *Client

	unregister chan *Client

	rdb *redis.Client

	redisClient *redisclient.RedisClient

	wordService *word.Service

	// Mutex to protect the maps (Go maps are not thread-safe)
	mu sync.RWMutex
}

// ClientActionRequest bundles a client with their action
type ClientActionRequest struct {
	Client *Client
	Action *ClientAction
}

func createHub(rdb *redis.Client, redisClient *redisclient.RedisClient, wordService *word.Service) *Hub {
	return &Hub{
		actions:     make(chan *ClientActionRequest, 256),
		register:    make(chan *Client, 256),
		unregister:  make(chan *Client, 256),
		games:       make(map[string]map[*Client]bool),
		clients:     make(map[string]*Client),
		rdb:         rdb,
		redisClient: redisClient,
		wordService: wordService,
	}
}

func (h *Hub) Run() {
	// Start listening to Redis Pub/Sub in background
	go h.ListenToRedis()

	// Main event loop
	for {
		select {
		case client := <-h.register:
			h.handleRegister(client)

		case client := <-h.unregister:
			h.handleUnregister(client)

		case actionReq := <-h.actions:
			h.handleAction(actionReq)
		}
	}
}

func (h *Hub) handleRegister(client *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.clients[client.UserID] = client
	log.Printf("Client registered: %s", client.UserID)
}

func (h *Hub) handleUnregister(client *Client) {
	// First, leave any game (updates Redis connected_count)
	if client.GameID != "" {
		h.leaveGameInternal(client)
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	// Remove from clients map and close send channel
	if _, ok := h.clients[client.UserID]; ok {
		delete(h.clients, client.UserID)
		close(client.Send)
	}

	log.Printf("Client unregistered: %s", client.UserID)
}

func (h *Hub) handleAction(actionReq *ClientActionRequest) {
	client := actionReq.Client
	action := actionReq.Action

	switch action.Action {
	case "join_game":
		h.handleJoinGame(client, action.GameID)
	case "leave_game":
		h.handleLeaveGame(client)
	case "submit_word":
		h.handleSubmitWord(client, action.GameID, action.Word)
	default:
		log.Printf("Unknown action: %s", action.Action)
	}
}

func (h *Hub) handleJoinGame(client *Client, gameID string) {
	// Leave previous game if in one (with Redis update)
	if client.GameID != "" && client.GameID != gameID {
		h.leaveGameInternal(client)
	}

	// Join the game session in Redis (atomic increment of connected_count)
	connectedCount, gameStarted, err := h.redisClient.AtomicJoinGameSession(gameID)
	if err != nil {
		log.Printf("Error joining game session: %v", err)
		h.sendErrorToClient(client, "join_failed", err.Error())
		return
	}

	// Add client to local game map
	h.mu.Lock()
	client.GameID = gameID
	if h.games[gameID] == nil {
		h.games[gameID] = make(map[*Client]bool)
	}
	h.games[gameID][client] = true
	h.mu.Unlock()

	log.Printf("Client %s joined game %s (connected: %d, started: %v)", client.UserID, gameID, connectedCount, gameStarted)

	// Send join confirmation to client
	h.sendToClient(client, GameMessage{
		Type:   "joined_game",
		GameID: gameID,
	})

	// If game just started (both players connected), fetch game state and broadcast
	if gameStarted {
		game, err := h.redisClient.GetGame(gameID)
		if err != nil {
			log.Printf("Error getting game state: %v", err)
			return
		}

		// Broadcast game start to all clients in this game
		h.broadcastGameMessage(gameID, GameMessage{
			Type:   "game_started",
			GameID: gameID,
			Payload: map[string]interface{}{
				"currentWord":   game.CurrentWord,
				"currentTurnId": game.CurrentTurnID,
				"player1Id":     game.Player1ID,
				"player1Name":   game.Player1Name,
				"player2Id":     game.Player2ID,
				"player2Name":   game.Player2Name,
				"startWord":     game.CurrentWord,
			},
		})
	}
}

func (h *Hub) leaveGameInternal(client *Client) {
	if client.GameID == "" {
		return
	}

	gameID := client.GameID

	// Decrement connected_count in Redis
	if err := h.redisClient.AtomicLeaveGameSession(gameID); err != nil {
		log.Printf("Error leaving game session: %v", err)
	}

	// Remove from local game map
	h.mu.Lock()
	if gameClients, ok := h.games[gameID]; ok {
		delete(gameClients, client)
		if len(gameClients) == 0 {
			delete(h.games, gameID)
		}
	}
	client.GameID = ""
	h.mu.Unlock()
}

func (h *Hub) handleLeaveGame(client *Client) {
	if client.GameID == "" {
		return
	}

	log.Printf("Client %s leaving game %s", client.UserID, client.GameID)
	h.leaveGameInternal(client)
}

func (h *Hub) handleSubmitWord(client *Client, gameID string, word string) {
	if client.GameID != gameID {
		h.sendErrorToClient(client, "submit_failed", "not_in_game")
		return
	}

	// First, get current game state to validate the move
	game, err := h.redisClient.GetGame(gameID)
	if err != nil {
		log.Printf("Error getting game: %v", err)
		h.sendErrorToClient(client, "submit_failed", "game_not_found")
		return
	}

	// Validate it's this player's turn
	if game.CurrentTurnID != client.UserID {
		h.sendErrorToClient(client, "submit_failed", "not_your_turn")
		return
	}

	// Check if word has already been played (before validating move)
	alreadyPlayed, err := h.redisClient.IsWordPlayed(gameID, word)
	if err != nil {
		log.Printf("Error checking if word played: %v", err)
		h.sendErrorToClient(client, "submit_failed", "server_error")
		return
	}
	if alreadyPlayed {
		h.sendErrorToClient(client, "submit_failed", "word_already_played")
		return
	}

	// Validate the word is a valid move using word service
	if !h.wordService.IsValidMove(game.CurrentWord, word) {
		h.sendErrorToClient(client, "submit_failed", "invalid_move")
		return
	}

	// Atomically submit the word (updates game state and resets timer)
	success, newWord, nextTurnID, err := h.redisClient.AtomicSubmitWord(gameID, client.UserID, word)
	if err != nil {
		log.Printf("Error submitting word: %v", err)
		// Check for specific error types
		if err.Error() == "word_already_played" {
			h.sendErrorToClient(client, "submit_failed", "word_already_played")
		} else {
			h.sendErrorToClient(client, "submit_failed", err.Error())
		}
		return
	}

	if !success {
		h.sendErrorToClient(client, "submit_failed", "atomic_failed")
		return
	}

	// Get the player name from game state
	playerName := game.Player1Name
	if client.UserID == game.Player2ID {
		playerName = game.Player2Name
	}

	log.Printf("Client %s submitted word '%s' for game %s, next turn: %s", client.UserID, word, gameID, nextTurnID)

	// Broadcast the move to all clients in this game
	h.broadcastGameMessage(gameID, GameMessage{
		Type:   "word_submitted",
		GameID: gameID,
		Payload: map[string]interface{}{
			"playerId":      client.UserID,
			"playerName":    playerName,
			"word":          newWord,
			"currentTurnId": nextTurnID,
		},
	})
}

// sendToClient sends a message to a specific client
func (h *Hub) sendToClient(client *Client, msg GameMessage) {
	data, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Error marshaling message: %v", err)
		return
	}
	select {
	case client.Send <- data:
	default:
		log.Printf("Client %s send buffer full", client.UserID)
	}
}

// sendErrorToClient sends an error message to a client
func (h *Hub) sendErrorToClient(client *Client, errorType string, message string) {
	h.sendToClient(client, GameMessage{
		Type:   "error",
		GameID: client.GameID,
		Payload: map[string]interface{}{
			"errorType": errorType,
			"message":   message,
		},
	})
}

// broadcastGameMessage sends a message to all clients in a game
func (h *Hub) broadcastGameMessage(gameID string, msg GameMessage) {
	data, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Error marshaling broadcast message: %v", err)
		return
	}
	h.BroadcastToGame(gameID, data)
}

// BroadcastToGame sends a message to all clients in a game
func (h *Hub) BroadcastToGame(gameID string, message []byte) {
	h.mu.RLock()
	gameClients, ok := h.games[gameID]
	if !ok {
		h.mu.RUnlock()
		return
	}

	// Copy client list to avoid holding lock during send
	clientList := make([]*Client, 0, len(gameClients))
	for client := range gameClients {
		clientList = append(clientList, client)
	}
	h.mu.RUnlock()

	// Send to each client, collect failures
	var failedClients []*Client
	for _, client := range clientList {
		select {
		case client.Send <- message:
		default:
			// Buffer full, mark for removal
			failedClients = append(failedClients, client)
		}
	}

	// Clean up failed clients (send to unregister channel to handle properly)
	for _, client := range failedClients {
		log.Printf("Client %s buffer full, scheduling disconnect", client.UserID)
		// Don't block if unregister channel is full
		select {
		case h.unregister <- client:
		default:
			log.Printf("Unregister channel full for client %s", client.UserID)
		}
	}
}

func (h *Hub) ListenToRedis() {
	ctx := context.Background()
	// Subscribe to a pattern so we catch updates for ANY game
	pubsub := h.rdb.PSubscribe(ctx, "game:*")
	defer pubsub.Close()

	ch := pubsub.Channel()

	for msg := range ch {
		// Extract gameID from channel name (format: "game:{gameId}")
		gameID := msg.Channel[5:] // Remove "game:" prefix

		// Broadcast the raw message to all clients in this game
		h.BroadcastToGame(gameID, []byte(msg.Payload))
	}
}
