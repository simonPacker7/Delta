package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/simonPacker7/Delta/backend/worker/handlers"
	authService "github.com/simonPacker7/Delta/backend/worker/services/auth"
	sessionService "github.com/simonPacker7/Delta/backend/worker/services/session"
)

func AuthRouter(app fiber.Router, auth *authService.Service, sess *sessionService.Service) {
	app.Post("/register", handlers.RegisterUser(auth, sess))
	app.Post("/login", handlers.LoginUser(auth, sess))
	app.Post("/logout", handlers.LogoutUser(sess))
}
