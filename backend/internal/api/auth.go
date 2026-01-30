package api

import (
  "encoding/json"
  "net/http"

  "github.com/xiaozhi-scientific/backend/internal/model"
  "github.com/xiaozhi-scientific/backend/internal/service"
)

func Login(authService *service.AuthService) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    var credentials struct {
      Username string `json:"username"`
      Password string `json:"password"`
    }

    if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
      writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
      return
    }

    token, user, err := authService.Login(r.Context(), credentials.Username, credentials.Password)
    if err != nil {
      writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "Invalid username or password"})
      return
    }

    writeJSON(w, http.StatusOK, map[string]any{
      "token": token,
      "user":  user,
    })
  }
}

func Register(authService *service.AuthService) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    var user model.User

    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
      writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
      return
    }

    token, createdUser, err := authService.Register(r.Context(), &user)
    if err != nil {
      writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
      return
    }

    writeJSON(w, http.StatusCreated, map[string]any{
      "token": token,
      "user":  createdUser,
    })
  }
}

func Logout() http.HandlerFunc {
  return func(w http.ResponseWriter, _ *http.Request) {
    writeJSON(w, http.StatusOK, map[string]string{"message": "Successfully logged out"})
  }
}

func ResetPassword(authService *service.AuthService) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    var request struct {
      Email string `json:"email"`
    }

    if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
      writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
      return
    }

    if err := authService.SendPasswordResetEmail(r.Context(), request.Email); err != nil {
      writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to send reset email"})
      return
    }

    writeJSON(w, http.StatusOK, map[string]string{"message": "Password reset email sent"})
  }
}
