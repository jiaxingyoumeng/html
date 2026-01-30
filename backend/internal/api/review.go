package api

import (
  "encoding/json"
  "net/http"
  "strconv"
  "strings"

  "github.com/xiaozhi-scientific/backend/internal/service"
)

func CreateReview(reviewService *service.ReviewService) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    var payload map[string]any
    if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
      writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
      return
    }
    review, err := reviewService.CreateReview(r.Context(), payload)
    if err != nil {
      writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
      return
    }
    writeJSON(w, http.StatusCreated, review)
  }
}

func GetReview(reviewService *service.ReviewService) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    id := strings.TrimPrefix(r.URL.Path, "/api/reviews/")
    review, err := reviewService.GetReview(r.Context(), id)
    if err != nil {
      writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
      return
    }
    writeJSON(w, http.StatusOK, review)
  }
}

func UpdateReview(reviewService *service.ReviewService) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    id := strings.TrimPrefix(r.URL.Path, "/api/reviews/")
    var payload map[string]any
    if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
      writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
      return
    }
    review, err := reviewService.UpdateReview(r.Context(), id, payload)
    if err != nil {
      writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
      return
    }
    writeJSON(w, http.StatusOK, review)
  }
}

func DeleteReview(reviewService *service.ReviewService) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    id := strings.TrimPrefix(r.URL.Path, "/api/reviews/")
    if err := reviewService.DeleteReview(r.Context(), id); err != nil {
      writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
      return
    }
    w.WriteHeader(http.StatusNoContent)
  }
}

func GetUserReviews(reviewService *service.ReviewService) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    page, _ := strconv.Atoi(r.URL.Query().Get("page"))
    if page == 0 {
      page = 1
    }
    limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
    if limit == 0 {
      limit = 10
    }
    reviews, err := reviewService.GetUserReviews(r.Context(), page, limit)
    if err != nil {
      writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
      return
    }
    writeJSON(w, http.StatusOK, reviews)
  }
}

func GenerateReviewContent(reviewService *service.ReviewService) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    id := strings.TrimPrefix(r.URL.Path, "/api/reviews/")
    result, err := reviewService.GenerateReviewContent(r.Context(), id)
    if err != nil {
      writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
      return
    }
    writeJSON(w, http.StatusAccepted, result)
  }
}

func CheckGenerationStatus(reviewService *service.ReviewService) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    taskID := strings.TrimPrefix(r.URL.Path, "/api/reviews/generation/")
    result, err := reviewService.CheckGenerationStatus(r.Context(), taskID)
    if err != nil {
      writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
      return
    }
    writeJSON(w, http.StatusOK, result)
  }
}
