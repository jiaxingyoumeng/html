package api

import (
  "net/http"

  "github.com/gin-gonic/gin"
  "github.com/xiaozhi-scientific/backend/internal/service"
)

type WritingRequest struct {
  Title          string `json:"title" binding:"required"`
  Topic          string `json:"topic" binding:"required"`
  Locale         string `json:"locale"`
  PubmedQuery    string `json:"pubmedQuery"`
  IncludeFilters bool   `json:"includeFilters"`
}

func StartWriting(reviewService *service.ReviewService) gin.HandlerFunc {
  return func(c *gin.Context) {
    var payload WritingRequest
    if err := c.ShouldBindJSON(&payload); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }

    result, err := reviewService.StartWritingWorkflow(c.Request.Context(), service.WritingRequest{
      Title:          payload.Title,
      Topic:          payload.Topic,
      Locale:         payload.Locale,
      PubmedQuery:    payload.PubmedQuery,
      IncludeFilters: payload.IncludeFilters,
    })
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
      return
    }

    c.JSON(http.StatusAccepted, result)
  }
}

func GetWritingStatus(reviewService *service.ReviewService) gin.HandlerFunc {
  return func(c *gin.Context) {
    result, err := reviewService.GetWritingStatus(c.Param("id"))
    if err != nil {
      c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
      return
    }
    c.JSON(http.StatusOK, result)
  }
}

func DownloadWritingDocument(reviewService *service.ReviewService) gin.HandlerFunc {
  return func(c *gin.Context) {
    filePath, err := reviewService.GetWritingDocumentPath(c.Param("id"))
    if err != nil {
      c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
      return
    }
    c.FileAttachment(filePath, "review.docx")
  }
}
