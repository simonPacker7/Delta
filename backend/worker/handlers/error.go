package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type GlobalErrorHandlerResp struct {
	Success   bool   `json:"success"`
	Message   string `json:"message"`
	ErrorType string `json:"errorType"`
}

type GlobalResponse struct {
	Success bool   `json:"success"`
	Data    string `json:"data"`
	Error   string `json:"error"`
}

func SendError(c *fiber.Ctx, errorType string, message string) error {
	return c.Status(200).JSON(GlobalErrorHandlerResp{
		Success:   false,
		ErrorType: errorType,
		Message:   message,
	})
}

func ErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}

func ValidationError(errorType string, message string) GlobalErrorHandlerResp {
	return GlobalErrorHandlerResp{
		Success:   false,
		ErrorType: errorType,
		Message:   message,
	}
}
