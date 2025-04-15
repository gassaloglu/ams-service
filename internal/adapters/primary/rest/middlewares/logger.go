package middlewares

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func Logger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Start timer
		start := time.Now()

		// Process request
		err := c.Next()

		// Get request timing
		latency := time.Since(start)

		// Get status code and size
		statusCode := c.Response().StatusCode()
		size := len(c.Response().Body())

		// Get client IP
		clientIP := c.IP()

		// Get path
		path := c.Path()
		raw := string(c.Request().URI().QueryString())

		// Log the request
		logEvent := log.Info()
		if statusCode >= 400 {
			logEvent = log.Error().Int("status", statusCode)
		}

		// Append query string if exists
		if raw != "" {
			path = path + "?" + raw
		}

		// Log the request details
		logEvent.
			Str("method", c.Method()).
			Str("path", path).
			Dur("latency", latency).
			Str("ip", clientIP).
			Int("size", size).
			Msg("HTTP Request")

		return err
	}
}
