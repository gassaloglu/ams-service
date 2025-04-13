package middlewares

import (
	"ams-service/internal/ports/primary"
	"errors"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

const (
	HrRole                string = "hr"
	AdminRole             string = "admin"
	FlightPlannerRole     string = "flight_planner"
	PassengerServicesRole string = "passenger_services"
	GroundServicesRole    string = "ground_services"
)

func TokenServiceInjector(service primary.TokenService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		ctx.Locals("tokenService", service)
		return ctx.Next()
	}
}

func Protection() fiber.Handler {
	return createProtection(func(token string, tokenService primary.TokenService) error {
		return tokenService.ValidateToken(token)
	})
}

func ProtectionForUsers() fiber.Handler {
	return createProtection(func(token string, tokenService primary.TokenService) error {
		return tokenService.ValidateUserToken(token)
	})
}

func ProtectionForEmployees() fiber.Handler {
	return createProtection(func(token string, tokenService primary.TokenService) error {
		return tokenService.ValidateEmployeeToken(token)
	})
}

func ProtectionForRoles(allowedRoles []string) fiber.Handler {
	return createProtection(func(token string, tokenService primary.TokenService) error {
		return tokenService.ValidateRole(token, allowedRoles)
	})
}

func isAuthDisabled() bool {
	return strings.EqualFold(os.Getenv("DISABLE_AUTH"), "true")
}

func getTokenService(ctx *fiber.Ctx) primary.TokenService {
	service := ctx.Locals("tokenService")
	if service == nil {
		log.Fatal("Token service not found in fiber context.")
	}

	return service.(primary.TokenService)
}

func extractBearerToken(ctx *fiber.Ctx) (string, error) {
	authHeader := ctx.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization header is required")
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		return "", errors.New("bearer token is required")
	}

	return tokenString, nil
}

func createProtection(rule func(string, primary.TokenService) error) fiber.Handler {
	if isAuthDisabled() {
		return nil
	}

	return func(ctx *fiber.Ctx) error {
		tokenService := getTokenService(ctx)
		token, err := extractBearerToken(ctx)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		err = rule(token, tokenService)
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, err.Error())
		}

		return ctx.Next()
	}
}
