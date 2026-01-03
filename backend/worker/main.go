package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/simonPacker7/Delta/backend/shared/postgresclient"
	"github.com/simonPacker7/Delta/backend/shared/redisclient"
	"github.com/simonPacker7/Delta/backend/worker/routes"
	authService "github.com/simonPacker7/Delta/backend/worker/services/auth"
	gameService "github.com/simonPacker7/Delta/backend/worker/services/game"
	sessionService "github.com/simonPacker7/Delta/backend/worker/services/session"
	userService "github.com/simonPacker7/Delta/backend/worker/services/user"
	wordService "github.com/simonPacker7/Delta/backend/worker/services/word"
)

func main() {
	// Get environment vars
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080"
	}

	redisURL := os.Getenv("REDIS_URL")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	dbURL := os.Getenv("DATABASE_URL")
	dbName := os.Getenv("DATABASE_NAME")
	dbUsername := os.Getenv("DATABASE_USERNAME")
	dbPassword := os.Getenv("DATABASE_PASSWORD")

	startWordsPath := os.Getenv("START_WORDS_PATH")
	if startWordsPath == "" {
		startWordsPath = "/app/assets/4-StartWords.json"
	}

	var redisConfig = redisclient.RedisConfig{
		Addr:     redisURL,
		Password: redisPassword,
		DB:       0,
	}

	var postgresConfig = postgresclient.PostgresConfig{
		Addr:     dbURL,
		DB:       dbName,
		Username: dbUsername,
		Password: dbPassword,
	}

	// Create repositories
	rClient := redisclient.NewRedisClient(redisConfig)
	pClient, err := postgresclient.NewPostgresClient(postgresConfig)
	if err != nil {
		log.Println(err)
	}

	app := fiber.New()

	// Create services
	auth := authService.NewService(pClient)
	session := sessionService.NewService(rClient, &redisConfig)
	users := userService.NewService(pClient)
	words := wordService.NewService(startWordsPath)
	game := gameService.NewService(rClient, words)

	// Create endpoints
	routes.AuthRouter(app.Group("/api/auth"), auth, session)
	routes.UserRouter(app.Group("/api/user"), users, session)
	routes.GameRouter(app.Group("/api/game"), game, session)

	app.Listen(":" + port)
}
