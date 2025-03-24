package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const AUTH_LOG_PREFIX = "auth.go"

func AuthMiddleware(jwtSecretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			LogError(fmt.Sprintf("%s - Authorization header is required", AUTH_LOG_PREFIX))
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			LogError(fmt.Sprintf("%s - Bearer token is required", AUTH_LOG_PREFIX))
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Bearer token is required"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(jwtSecretKey), nil
		})

		if err != nil || !token.Valid {
			LogError(fmt.Sprintf("%s - Invalid token: %v", AUTH_LOG_PREFIX, err))
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		LogInfo(fmt.Sprintf("%s - Token validated successfully", AUTH_LOG_PREFIX))
		c.Next()
	}
}
