package utils

import (
	"ams-service/internal/config"
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func ExtractIDFromToken(ctx *gin.Context, idKey string) (string, error) {
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
		id, ok := claims[idKey].(string)
		if !ok {
			// Handle case where ID is a float64
			if idFloat, ok := claims[idKey].(float64); ok {
				id = fmt.Sprintf("%.0f", idFloat)
			} else {
				return "", fmt.Errorf("invalid %s type", idKey)
			}
		}
		return id, nil
	}

	return "", fmt.Errorf("invalid token")
}

func ExtractUserIDFromToken(ctx *gin.Context) (string, error) {
	return ExtractIDFromToken(ctx, "user_id")
}

func ExtractEmployeeIDFromToken(ctx *gin.Context) (string, error) {
	return ExtractIDFromToken(ctx, "employee_id")
}
