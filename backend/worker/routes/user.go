package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/simonPacker7/Delta/backend/worker/handlers"
	sessionService "github.com/simonPacker7/Delta/backend/worker/services/session"
	userService "github.com/simonPacker7/Delta/backend/worker/services/user"
)

func UserRouter(app fiber.Router, users *userService.Service, sess *sessionService.Service) {
	app.Use(handlers.AuthRoute(sess))
	app.Get("/profile", handlers.GetUserProfile(users))
}
