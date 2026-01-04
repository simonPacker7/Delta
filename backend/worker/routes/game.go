package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/simonPacker7/Delta/backend/worker/handlers"
	gameService "github.com/simonPacker7/Delta/backend/worker/services/game"
	sessionService "github.com/simonPacker7/Delta/backend/worker/services/session"
)

func GameRouter(app fiber.Router, game *gameService.Service, sess *sessionService.Service) {
	// All game routes require authentication
	app.Use(handlers.AuthRoute(sess))

	app.Get("/find", handlers.FindGame(game))
	app.Post("/private/create", handlers.CreatePrivateGame(game))
	app.Post("/private/join", handlers.JoinPrivateGame(game))
	app.Delete("/matchmaking/:id", handlers.CancelMatchmaking(game))
	app.Get("/:id", handlers.GetGame(game))
}
