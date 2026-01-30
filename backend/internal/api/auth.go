package api

import (
  "net/http"

  "github.com/gin-gonic/gin"
  "github.com/xiaozhi-scientific/backend/internal/model"
  "github.com/xiaozhi-scientific/backend/internal/service"
)

func Login(authService *service.AuthService) gin.HandlerFunc {
  return func(c *gin.Context) {
    var credentials struct {
      Username string `json:"username" binding:"required"`
      Password string `json:"password" binding:"required"`
    }

    if err := c.ShouldBindJSON(&credentials); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }

    token, user, err := authService.Login(c.Request.Context(), credentials.Username, credentials.Password)
    if err != nil {
      c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
      return
    }

    c.JSON(http.StatusOK, gin.H{
      "token": token,
      "user":  user,
    })
  }
}

func Register(authService *service.AuthService) gin.HandlerFunc {
  return func(c *gin.Context) {
    var user model.User

    if err := c.ShouldBindJSON(&user); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }

    token, createdUser, err := authService.Register(c.Request.Context(), &user)
    if err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }

    c.JSON(http.StatusCreated, gin.H{
      "token": token,
      "user":  createdUser,
    })
  }
}

func Logout() gin.HandlerFunc {
  return func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
  }
}

func ResetPassword(authService *service.AuthService) gin.HandlerFunc {
  return func(c *gin.Context) {
    var request struct {
      Email string `json:"email" binding:"required,email"`
    }

    if err := c.ShouldBindJSON(&request); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }

    if err := authService.SendPasswordResetEmail(c.Request.Context(), request.Email); err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send reset email"})
      return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Password reset email sent"})
  }
}
