package main

import (
	"context"
	"encoding/json"
	"log"
	"sync"

	"github.com/redis/go-redis/v9"
)

type GameMessage struct {
	SenderID string `json:"sender_id"`
	GameID   string `json:"game_id"`
	Type     string `json:"message_type"` // eg. "MOVE" | "RESIGN" | "START" etc..
	Payload  string `json:"payload"`
}

type Hub struct {

	// Active clients mapped by GameID
	games map[string]map[*Client]bool

	// Inbound messages from clients (to be published to Redis)
	broadcast chan *GameMessage

	register chan *Client

	unregister chan *Client

	rdb *redis.Client

	// Mutex to protect the `games` map (Go maps are not thread-safe)
	mu sync.RWMutex
}

func createHub(rdb *redis.Client) *Hub {
	return &Hub{
		broadcast:  make(chan *GameMessage),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		games:      make(map[string]map[*Client]bool),
		rdb:        rdb,
	}
}

func (h *Hub) Run() {
	// Start listening to Redis global events in a background routine
	go h.ListenToRedis()
}

func (h *Hub) ListenToRedis() {
	ctx := context.Background()
	// Subscribe to a pattern so we catch updates for ANY game
	pubsub := h.rdb.PSubscribe(ctx, "game:*")
	defer pubsub.Close()

	ch := pubsub.Channel()

	for msg := range ch {
		var gameMsg GameMessage
		if err := json.Unmarshal([]byte(msg.Payload), &gameMsg); err != nil {
			log.Printf("Error unmarshalling from Redis: %v", err)
			continue
		}

		// Find if we have any local players for this game
		h.mu.RLock()
		clients, ok := h.games[gameMsg.GameID]
		h.mu.RUnlock()

		if ok {
			for client := range clients {
				select {
				case client.Send <- []byte(gameMsg.Payload): // TODO change -> send the entire game message
				default:
					// If client buffer is full, close and remove to prevent blocking
					log.Printf("Client Send buffer is full, closing and removing client %v", client.UserID)
					close(client.Send)
					h.mu.Lock()
					delete(h.games[gameMsg.GameID], client)
					h.mu.Unlock()
				}
			}
		}
	}
}
