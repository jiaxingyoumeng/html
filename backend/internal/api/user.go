package api

import (
  "net/http"

  "github.com/gin-gonic/gin"
  "github.com/xiaozhi-scientific/backend/internal/service"
)

func GetProfile(userService *service.UserService) gin.HandlerFunc {
  return func(c *gin.Context) {
    userID := c.GetString("userId")
    profile, err := userService.GetProfile(c.Request.Context(), userID)
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
    }
    c.JSON(http.StatusOK, profile)
  }
}

func UpdateProfile(userService *service.UserService) gin.HandlerFunc {
  return func(c *gin.Context) {
    userID := c.GetString("userId")
    var payload map[string]any
    if err := c.ShouldBindJSON(&payload); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }
    profile, err := userService.UpdateProfile(c.Request.Context(), userID, payload)
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
    }
    c.JSON(http.StatusOK, profile)
  }
}

func GetSubscription(userService *service.UserService) gin.HandlerFunc {
  return func(c *gin.Context) {
    userID := c.GetString("userId")
    subscription, err := userService.GetSubscription(c.Request.Context(), userID)
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
    }
    c.JSON(http.StatusOK, subscription)
  }
}

func UpdateSubscription(userService *service.UserService) gin.HandlerFunc {
  return func(c *gin.Context) {
    userID := c.GetString("userId")
    var payload struct {
      PlanID string `json:"planId"`
    }
    if err := c.ShouldBindJSON(&payload); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }
    subscription, err := userService.UpdateSubscription(c.Request.Context(), userID, payload.PlanID)
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
    }
    c.JSON(http.StatusOK, subscription)
  }
}
