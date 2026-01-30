package api

import (
  "encoding/json"
  "net/http"

  "github.com/xiaozhi-scientific/backend/internal/service"
)

func GetProfile(userService *service.UserService) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    userID := r.URL.Query().Get("userId")
    profile, err := userService.GetProfile(r.Context(), userID)
    if err != nil {
      writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
      return
    }
    writeJSON(w, http.StatusOK, profile)
  }
}

func UpdateProfile(userService *service.UserService) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    userID := r.URL.Query().Get("userId")
    var payload map[string]any
    if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
      writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
      return
    }
    profile, err := userService.UpdateProfile(r.Context(), userID, payload)
    if err != nil {
      writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
      return
    }
    writeJSON(w, http.StatusOK, profile)
  }
}

func GetSubscription(userService *service.UserService) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    userID := r.URL.Query().Get("userId")
    subscription, err := userService.GetSubscription(r.Context(), userID)
    if err != nil {
      writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
      return
    }
    writeJSON(w, http.StatusOK, subscription)
  }
}

func UpdateSubscription(userService *service.UserService) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    userID := r.URL.Query().Get("userId")
    var payload struct {
      PlanID string `json:"planId"`
    }
    if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
      writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
      return
    }
    subscription, err := userService.UpdateSubscription(r.Context(), userID, payload.PlanID)
    if err != nil {
      writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
      return
    }
    writeJSON(w, http.StatusOK, subscription)
  }
}
