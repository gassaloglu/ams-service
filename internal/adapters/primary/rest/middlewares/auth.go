package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func AuthMiddleware(jwtSecretKey string) fiber.Handler {
	if strings.EqualFold(os.Getenv("DISABLE_AUTH"), "true") {
		return func(c *fiber.Ctx) error {
			return c.Next()
		}
	}

	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			log.Error().Msg("Authorization header is required")
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "Authorization header is required",
			})
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			log.Error().Msg("Bearer token is required")
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "Bearer token is required",
			})
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(jwtSecretKey), nil
		})

		if err != nil || !token.Valid {
			log.Error().Err(err).Msg("Invalid token")
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token",
			})
		}

		log.Info().Msg("Token validated successfully")
		return c.Next()
	}
}
