package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/simonPacker7/Delta/backend/shared/entities"
	userService "github.com/simonPacker7/Delta/backend/worker/services/user"
)

func GetUserProfile(users *userService.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		session := c.Locals("sessionContext").(entities.SessionContext)

		profile, err := users.GetUserProfile(session.Email)

		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(ErrorResponse(err))
		}

		return c.JSON(profile)
	}
}
