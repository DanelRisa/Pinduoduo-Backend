package middleware

import (
    "github.com/gin-gonic/gin"
    "log"
    "time"
    "github.com/google/uuid"
)

// LoggingMiddleware
func LoggingMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        requestID := uuid.New().String()

        start := time.Now()

        // log.Printf("[%s] %s %s", requestID, c.Request.Method, c.Request.URL.Path)
        c.Next()
        duration := time.Since(start).Milliseconds()

        log.Printf("[%s] %s %s - %d - Duration: %dms", requestID, c.Request.Method, c.Request.URL.Path, c.Writer.Status(), duration)
    }
}
