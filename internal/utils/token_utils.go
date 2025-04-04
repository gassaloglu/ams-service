package utils

import (
	"ams-service/internal/config"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func ExtractIDFromToken(ctx *fiber.Ctx, idKey string) (string, error) {
	authHeader := ctx.Get("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.JWTSecretKey), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		value, ok := claims[idKey]
		if !ok {
			return "", fmt.Errorf("%s not found in token", idKey)
		}

		// Handle different types of values
		switch v := value.(type) {
		case string:
			return v, nil
		case float64:
			return fmt.Sprintf("%.0f", v), nil
		default:
			return "", fmt.Errorf("invalid %s type: %T", idKey, v)
		}
	}

	return "", fmt.Errorf("invalid token")
}
func ExtractUserIDFromToken(ctx *fiber.Ctx) (string, error) {
	return ExtractIDFromToken(ctx, "user_id")
}

func ExtractEmployeeIDFromToken(ctx *fiber.Ctx) (string, error) {
	return ExtractIDFromToken(ctx, "employee_id")
}

func ExtractRoleFromToken(ctx *fiber.Ctx) (string, error) {
	return ExtractIDFromToken(ctx, "role")
}

// Helper function to extract user or employee ID from token
func ExtractUserOrEmployeeID(ctx *fiber.Ctx) (string, error) {
	userID, userErr := ExtractUserIDFromToken(ctx)
	if userErr == nil {
		return userID, nil
	}

	employeeID, employeeErr := ExtractEmployeeIDFromToken(ctx)
	if employeeErr == nil {
		return employeeID, nil
	}

	log.Error().Err(userErr).Err(employeeErr).Msg("Error extracting user or employee ID from token")
	return "", fiber.NewError(http.StatusUnauthorized, "Unauthorized")
}

// Helper function to check role authorization
func CheckRoleAuthorization(ctx *fiber.Ctx, allowedRoles []string) (string, error) {
	role, err := ExtractRoleFromToken(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Error extracting role from token")
		return "", fiber.NewError(http.StatusUnauthorized, "Unauthorized")
	}

	if role == "" {
		log.Warn().Msg("Role not found in token")
		return "", fiber.NewError(http.StatusForbidden, "Forbidden")
	}

	for _, allowedRole := range allowedRoles {
		if role == allowedRole {
			return role, nil
		}
	}

	log.Warn().Str("role", role).Msg("Unauthorized role attempting to access resource")
	return "", fiber.NewError(http.StatusForbidden, "Forbidden")
}
