package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/simonPacker7/Delta/backend/shared/entities"
	gameService "github.com/simonPacker7/Delta/backend/worker/services/game"
)

func FindGame(game *gameService.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get user from session context (set by AuthRoute middleware)
		sessionCtx, ok := c.Locals("sessionContext").(entities.SessionContext)
		if !ok {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		response, err := game.FindGame(sessionCtx.ID, sessionCtx.Name)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(ErrorResponse(err))
		}

		return c.JSON(response)
	}
}

func GetGame(game *gameService.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		gameID := c.Params("id")
		if gameID == "" {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(ErrorResponse(fiber.NewError(fiber.StatusBadRequest, "game ID is required")))
		}

		gameData, err := game.GetGame(gameID)
		if err != nil {
			c.Status(fiber.StatusNotFound)
			return c.JSON(ErrorResponse(err))
		}

		return c.JSON(gameData)
	}
}

func CreatePrivateGame(game *gameService.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sessionCtx, ok := c.Locals("sessionContext").(entities.SessionContext)
		if !ok {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		response, err := game.CreatePrivateGame(sessionCtx.ID, sessionCtx.Name)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(ErrorResponse(err))
		}

		return c.JSON(response)
	}
}

func JoinPrivateGame(game *gameService.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sessionCtx, ok := c.Locals("sessionContext").(entities.SessionContext)
		if !ok {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		var requestBody entities.JoinPrivateGameInput
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(ErrorResponse(err))
		}

		if requestBody.JoinCode == "" {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(ErrorResponse(fiber.NewError(fiber.StatusBadRequest, "join code is required")))
		}

		response, err := game.JoinPrivateGame(sessionCtx.ID, sessionCtx.Name, requestBody.JoinCode)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(ErrorResponse(err))
		}

		return c.JSON(response)
	}
}
