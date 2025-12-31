package main

import (
	"log"
	"net/http"
	"os"

	"github.com/redis/go-redis/v9"
	"github.com/simonPacker7/Delta/backend/game-service/word"
	"github.com/simonPacker7/Delta/backend/shared/redisclient"
)

func main() {
	// Get environment vars
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080"
	}
	redisURL := os.Getenv("REDIS_URL")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	wordMapPath := os.Getenv("WORD_MAP_PATH")
	if wordMapPath == "" {
		wordMapPath = "/app/assets/4-WordMap.json"
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: redisPassword,
		DB:       0, // use default DB
	})

	// Initialize redis client wrapper for game operations
	redisClient := redisclient.NewRedisClient(redisclient.RedisConfig{
		Addr:     redisURL,
		Password: redisPassword,
		DB:       0,
	})

	// Initialize word service
	wordService := word.NewService(wordMapPath)

	hub := createHub(rdb, redisClient, wordService)
	go hub.Run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		userId, err := authenticateRequest(rdb, r)
		if err != nil {
			w.WriteHeader(401)
			return
		}

		ServeWs(hub, userId, w, r)
	})

	log.Println("Listening on port ", port)
	http.ListenAndServe(":"+port, nil)
}
