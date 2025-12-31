package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/simonPacker7/Delta/backend/shared/redisclient"
)

const (
	pollInterval     = 2 * time.Second
	batchSize        = 10
)

func main() {
	log.Println("Starting Arbiter service...")

	// Get environment vars
	redisURL := os.Getenv("REDIS_URL")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	if redisURL == "" {
		redisURL = "localhost:6379"
	}

	redisConfig := redisclient.RedisConfig{
		Addr:     redisURL,
		Password: redisPassword,
		DB:       0,
	}

	rClient := redisclient.NewRedisClient(redisConfig)
	log.Println("Connected to Redis")

	// Create arbiter
	arbiter := NewArbiter(rClient)

	// Handle graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Start polling loop
	go arbiter.Run()

	<-stop
	log.Println("Shutting down Arbiter...")
	arbiter.Stop()
}

type Arbiter struct {
	redisClient *redisclient.RedisClient
	stopChan    chan struct{}
	running     bool
}

func NewArbiter(r *redisclient.RedisClient) *Arbiter {
	return &Arbiter{
		redisClient: r,
		stopChan:    make(chan struct{}),
		running:     false,
	}
}

func (a *Arbiter) Run() {
	a.running = true
	ticker := time.NewTicker(pollInterval)
	defer ticker.Stop()

	log.Printf("Arbiter polling every %v for expired games", pollInterval)

	for {
		select {
		case <-ticker.C:
			a.processExpiredGames()
		case <-a.stopChan:
			log.Println("Arbiter stopped")
			return
		}
	}
}

func (a *Arbiter) Stop() {
	if a.running {
		close(a.stopChan)
		a.running = false
	}
}

func (a *Arbiter) processExpiredGames() {
	// Atomically claim AND end expired games in one operation
	// This prevents race conditions where a player moves between claim and end
	endedGames, err := a.redisClient.AtomicClaimAndEndExpiredGames(batchSize)
	if err != nil {
		log.Printf("Error processing expired games: %v", err)
		return
	}

	if len(endedGames) == 0 {
		return
	}

	for _, result := range endedGames {
		log.Printf("Game %s ended - winner: %s (opponent timed out)", result.GameID, result.WinnerID)
	}
}
}
