package middlewares

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	infoLogger  = log.New(os.Stdout, "[INFO] - ", log.Ldate|log.Ltime)
	errorLogger = log.New(os.Stderr, "[ERROR] - ", log.Ldate|log.Ltime)
)

func LogInfo(message string) {
	infoLogger.Println(message)
}

func LogError(message string) {
	errorLogger.Println(message)
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)

		LogInfo(fmt.Sprintf("%s %s %d %s", c.Request.Method, c.Request.URL.Path, c.Writer.Status(), duration))
	}
}
