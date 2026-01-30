package api

import (
  "net/http"
  "strconv"

  "github.com/gin-gonic/gin"
  "github.com/xiaozhi-scientific/backend/internal/service"
)

func CreateReview(reviewService *service.ReviewService) gin.HandlerFunc {
  return func(c *gin.Context) {
    var payload map[string]any
    if err := c.ShouldBindJSON(&payload); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }
    review, err := reviewService.CreateReview(c.Request.Context(), payload)
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
    }
    c.JSON(http.StatusCreated, review)
  }
}

func GetReview(reviewService *service.ReviewService) gin.HandlerFunc {
  return func(c *gin.Context) {
    review, err := reviewService.GetReview(c.Request.Context(), c.Param("id"))
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
    }
    c.JSON(http.StatusOK, review)
  }
}

func UpdateReview(reviewService *service.ReviewService) gin.HandlerFunc {
  return func(c *gin.Context) {
    var payload map[string]any
    if err := c.ShouldBindJSON(&payload); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }
    review, err := reviewService.UpdateReview(c.Request.Context(), c.Param("id"), payload)
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
    }
    c.JSON(http.StatusOK, review)
  }
}

func DeleteReview(reviewService *service.ReviewService) gin.HandlerFunc {
  return func(c *gin.Context) {
    if err := reviewService.DeleteReview(c.Request.Context(), c.Param("id")); err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
    }
    c.Status(http.StatusNoContent)
  }
}

func GetUserReviews(reviewService *service.ReviewService) gin.HandlerFunc {
  return func(c *gin.Context) {
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
    reviews, err := reviewService.GetUserReviews(c.Request.Context(), page, limit)
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
    }
    c.JSON(http.StatusOK, reviews)
  }
}

func GenerateReviewContent(reviewService *service.ReviewService) gin.HandlerFunc {
  return func(c *gin.Context) {
    result, err := reviewService.GenerateReviewContent(c.Request.Context(), c.Param("id"))
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
    }
    c.JSON(http.StatusAccepted, result)
  }
}

func CheckGenerationStatus(reviewService *service.ReviewService) gin.HandlerFunc {
  return func(c *gin.Context) {
    result, err := reviewService.CheckGenerationStatus(c.Request.Context(), c.Param("taskId"))
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
    }
    c.JSON(http.StatusOK, result)
  }
}
