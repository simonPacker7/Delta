package handlers

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	authService "github.com/simonPacker7/Delta/backend/api/services/auth"
	sessionService "github.com/simonPacker7/Delta/backend/api/services/session"
	"github.com/simonPacker7/Delta/backend/shared/entities"
)

func RegisterUser(auth *authService.Service, sess *sessionService.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.RegisterUserInput
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(ErrorResponse(err))
		}
		if requestBody.Email == "" || requestBody.Name == "" || requestBody.Password == "" {
			c.Status(http.StatusInternalServerError)
			return c.JSON(ErrorResponse(errors.New(
				"please specify email, name and password")))
		}
		userId, err := auth.RegisterUser(requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(ErrorResponse(err))
		}

		err = sess.SetSession(c, userId, requestBody.Name, requestBody.Email)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(ErrorResponse(err))
		}

		return c.SendStatus(200)
	}
}

func LoginUser(auth *authService.Service, sess *sessionService.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.LoginUserInput
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(ErrorResponse(err))
		}
		if requestBody.Email == "" || requestBody.Password == "" {
			c.Status(http.StatusInternalServerError)
			return c.JSON(ErrorResponse(errors.New(
				"please specify email and password")))
		}
		userProfile, err := auth.LoginUser(requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(ErrorResponse(err))
		}
		err = sess.SetSession(c, userProfile.ID, userProfile.Name, userProfile.Email)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(ErrorResponse(err))
		}

		return c.SendStatus(200)
	}
}

func LogoutUser(sess *sessionService.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess.DestorySession(c)

		return c.SendStatus(200)
	}
}

// Guards endpoints in the API
func AuthRoute(sess *sessionService.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sessionCtx, err := sess.GetSession(c)

		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		c.Locals("user", sessionCtx.Email)
		c.Locals("sessionContext", sessionCtx)
		return c.Next()
	}
}
