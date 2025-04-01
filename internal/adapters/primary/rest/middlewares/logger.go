package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process request
		c.Next()

		// Get request timing
		latency := time.Since(start)

		// Get status code and size
		statusCode := c.Writer.Status()
		size := c.Writer.Size()

		// Get client IP
		clientIP := c.ClientIP()

		// Get path
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

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
			Str("method", c.Request.Method).
			Str("path", path).
			Dur("latency", latency).
			Str("ip", clientIP).
			Int("size", size).
			Msg("HTTP Request")
	}
}
