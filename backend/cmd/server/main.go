package main

import (
  "log"
  "net/http"
  "strings"

  "github.com/xiaozhi-scientific/backend/internal/api"
  "github.com/xiaozhi-scientific/backend/internal/config"
  "github.com/xiaozhi-scientific/backend/internal/repository"
  "github.com/xiaozhi-scientific/backend/internal/service"
)

func main() {
  cfg, err := config.Load()
  if err != nil {
    log.Fatalf("Failed to load config: %v", err)
  }

  repo, err := repository.NewMongoRepository(cfg.Database.URI, cfg.Database.Name)
  if err != nil {
    log.Fatalf("Failed to connect to database: %v", err)
  }
  defer func() {
    if closeErr := repo.Close(); closeErr != nil {
      log.Printf("Failed to close repository: %v", closeErr)
    }
  }()

  authService := service.NewAuthService(repo, cfg.JWT.Secret, cfg.JWT.Expiration)
  userService := service.NewUserService(repo)
  reviewService := service.NewReviewService(repo, cfg.PubMed.BaseURL, cfg.PubMed.APIKey, "storage")

  mux := http.NewServeMux()

  mux.Handle("/api/auth/login", api.Login(authService))
  mux.Handle("/api/auth/register", api.Register(authService))
  mux.Handle("/api/auth/logout", api.Logout())
  mux.Handle("/api/auth/reset-password", api.ResetPassword(authService))

  mux.Handle("/api/user/profile", api.GetProfile(userService))
  mux.Handle("/api/user/subscription", api.GetSubscription(userService))
  mux.Handle("/api/user/profile/update", api.UpdateProfile(userService))
  mux.Handle("/api/user/subscription/update", api.UpdateSubscription(userService))

  mux.Handle("/api/writing", api.StartWriting(reviewService))
  mux.Handle("/api/writing/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    if strings.HasSuffix(r.URL.Path, "/download") {
      api.DownloadWritingDocument(reviewService)(w, r)
      return
    }
    api.GetWritingStatus(reviewService)(w, r)
  }))

  mux.Handle("/api/reviews", api.CreateReview(reviewService))
  mux.Handle("/api/reviews/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    if strings.Contains(r.URL.Path, "/generation/") {
      api.CheckGenerationStatus(reviewService)(w, r)
      return
    }
    switch r.Method {
    case http.MethodGet:
      api.GetReview(reviewService)(w, r)
    case http.MethodPut:
      api.UpdateReview(reviewService)(w, r)
    case http.MethodDelete:
      api.DeleteReview(reviewService)(w, r)
    case http.MethodPost:
      api.GenerateReviewContent(reviewService)(w, r)
    default:
      w.WriteHeader(http.StatusMethodNotAllowed)
    }
  }))

  server := &http.Server{
    Addr:    cfg.Server.Address,
    Handler: mux,
  }

  log.Printf("Starting server on %s", cfg.Server.Address)
  if err := server.ListenAndServe(); err != nil {
    log.Fatalf("Failed to start server: %v", err)
  }
}
