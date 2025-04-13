package middlewares

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	var fiberError *fiber.Error

	if errors.As(err, &fiberError) {
		log.Error().Int("status", fiberError.Code).Err(err).Msg("Request error")
		return c.Status(fiberError.Code).JSON(fiberError.Message)
	}

	return nil
}
