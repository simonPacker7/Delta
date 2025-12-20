package main

import (
	"log"
	"net/http"
	"os"

	"github.com/redis/go-redis/v9"
)

func main() {
	// Get environment vars
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080"
	}
	redisURL := os.Getenv("REDIS_URL")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: redisPassword,
		DB:       0, // use default DB
	})

	hub := createHub(rdb)
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
