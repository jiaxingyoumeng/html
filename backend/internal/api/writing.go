package api

import (
  "encoding/json"
  "net/http"
  "strings"

  "github.com/xiaozhi-scientific/backend/internal/service"
)

type WritingRequest struct {
  Title          string `json:"title"`
  Topic          string `json:"topic"`
  Locale         string `json:"locale"`
  PubmedQuery    string `json:"pubmedQuery"`
  IncludeFilters bool   `json:"includeFilters"`
}

func StartWriting(reviewService *service.ReviewService) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    var payload WritingRequest
    if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
      writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
      return
    }

    result, err := reviewService.StartWritingWorkflow(r.Context(), service.WritingRequest{
      Title:          payload.Title,
      Topic:          payload.Topic,
      Locale:         payload.Locale,
      PubmedQuery:    payload.PubmedQuery,
      IncludeFilters: payload.IncludeFilters,
    })
    if err != nil {
      writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
      return
    }

    writeJSON(w, http.StatusAccepted, result)
  }
}

func GetWritingStatus(reviewService *service.ReviewService) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    id := strings.TrimPrefix(r.URL.Path, "/api/writing/")
    result, err := reviewService.GetWritingStatus(strings.TrimSuffix(id, "/download"))
    if err != nil {
      writeJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
      return
    }
    writeJSON(w, http.StatusOK, result)
  }
}

func DownloadWritingDocument(reviewService *service.ReviewService) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    id := strings.TrimPrefix(r.URL.Path, "/api/writing/")
    id = strings.TrimSuffix(id, "/download")
    filePath, err := reviewService.GetWritingDocumentPath(id)
    if err != nil {
      writeJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
      return
    }
    http.ServeFile(w, r, filePath)
  }
}
