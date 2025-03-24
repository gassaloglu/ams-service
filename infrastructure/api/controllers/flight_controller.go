package controllers

import (
	"ams-service/config"
	"ams-service/core/entities"
	"ams-service/core/services"
	"ams-service/middlewares"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var FLIGHT_LOG_PREFIX string = "flight_controller.go"

type FlightController struct {
	service *services.FlightService
}

func NewFlightController(service *services.FlightService) *FlightController {
	return &FlightController{service: service}
}

func (c *FlightController) GetSpecificFlight(ctx *gin.Context) {
	var request entities.GetSpecificFlightRequest
	if err := ctx.ShouldBindQuery(&request); err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error binding query: %v", FLIGHT_LOG_PREFIX, err))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error binding query"})
		return
	}

	userID, err := extractUserIDFromToken(ctx)
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error extracting user ID from token: %v", FLIGHT_LOG_PREFIX, err))
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	resultChan := make(chan entities.Flight)
	errorChan := make(chan error)

	c.service.GetSpecificFlight(request, userID, resultChan, errorChan)

	select {
	case flight := <-resultChan:
		ctx.JSON(http.StatusOK, flight)
	case err := <-errorChan:
		middlewares.LogError(fmt.Sprintf("%s - Error getting specific flight: %v", FLIGHT_LOG_PREFIX, err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting specific flight"})
	case <-time.After(10 * time.Second):
		ctx.JSON(http.StatusGatewayTimeout, gin.H{"error": "Request timed out"})
	}
}

func (c *FlightController) GetAllFlights(ctx *gin.Context) {
	flights, err := c.service.GetAllFlights()
	if err != nil {
		middlewares.LogError(fmt.Sprintf("%s - Error getting all flights: %v", FLIGHT_LOG_PREFIX, err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "TODO: Error getting all flights"})
		return
	}
	ctx.JSON(http.StatusOK, flights)
}

func extractUserIDFromToken(ctx *gin.Context) (string, error) {
	authHeader := ctx.GetHeader("Authorization")
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
		userID, ok := claims["user_id"].(string)
		if !ok {
			// Handle case where user_id is a float64
			if userIDFloat, ok := claims["user_id"].(float64); ok {
				userID = fmt.Sprintf("%.0f", userIDFloat)
			} else {
				return "", fmt.Errorf("invalid user_id type")
			}
		}
		return userID, nil
	}

	return "", fmt.Errorf("invalid token")
}
