package middleware

import (
  "net/http"
  "strings"

  "github.com/gin-gonic/gin"
  "github.com/xiaozhi-scientific/backend/internal/service"
)

func AuthMiddleware(authService *service.AuthService) gin.HandlerFunc {
  return func(c *gin.Context) {
    authHeader := c.GetHeader("Authorization")
    if authHeader == "" {
      c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
      c.Abort()
      return
    }

    parts := strings.SplitN(authHeader, " ", 2)
    if len(parts) != 2 || parts[0] != "Bearer" {
      c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}"})
      c.Abort()
      return
    }

    claims, err := authService.ValidateToken(parts[1])
    if err != nil {
      c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
      c.Abort()
      return
    }

    if userID, ok := claims["userId"].(string); ok {
      c.Set("userId", userID)
    }
    if username, ok := claims["username"].(string); ok {
      c.Set("username", username)
    }

    c.Next()
  }
}
