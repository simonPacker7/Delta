package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/simonPacker7/Delta/backend/api/routes"
	authService "github.com/simonPacker7/Delta/backend/api/services/auth"
	sessionService "github.com/simonPacker7/Delta/backend/api/services/session"
	userService "github.com/simonPacker7/Delta/backend/api/services/user"
	"github.com/simonPacker7/Delta/backend/shared/postgresclient"
	"github.com/simonPacker7/Delta/backend/shared/redisclient"
)

func main() {
	// Get environment vars
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080"
	}

	dbURL := os.Getenv("DATABASE_URL")
	dbName := os.Getenv("DATABASE_NAME")
	dbUsername := os.Getenv("DATABASE_USERNAME")
	dbPassword := os.Getenv("DATABASE_PASSWORD")

	redisURL := os.Getenv("REDIS_URL")
	redisPassword := os.Getenv("REDIS_PASSWORD")

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

	// Create services
	auth := authService.NewService(pClient)
	session := sessionService.NewService(rClient, &redisConfig)
	users := userService.NewService(pClient)

	app := fiber.New()
	// Create endpoints
	routes.AuthRouter(app.Group("/api/auth"), auth, session)
	routes.UserRouter(app.Group("/api/user"), users, session)

	app.Listen(":" + port)
}
