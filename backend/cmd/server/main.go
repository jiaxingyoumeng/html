package main

import (
  "log"

  "github.com/gin-gonic/gin"
  "github.com/xiaozhi-scientific/backend/internal/api"
  "github.com/xiaozhi-scientific/backend/internal/config"
  "github.com/xiaozhi-scientific/backend/internal/middleware"
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

  router := gin.Default()

  apiGroup := router.Group("/api")
  {
    auth := apiGroup.Group("/auth")
    {
      auth.POST("/login", api.Login(authService))
      auth.POST("/register", api.Register(authService))
      auth.POST("/logout", api.Logout())
      auth.POST("/reset-password", api.ResetPassword(authService))
    }

    protected := apiGroup.Group("/")
    protected.Use(middleware.AuthMiddleware(authService))
    {
      writing := protected.Group("/writing")
      {
        writing.POST("/", api.StartWriting(reviewService))
        writing.GET("/:id", api.GetWritingStatus(reviewService))
        writing.GET("/:id/download", api.DownloadWritingDocument(reviewService))
      }

      user := protected.Group("/user")
      {
        user.GET("/profile", api.GetProfile(userService))
        user.PUT("/profile", api.UpdateProfile(userService))
        user.GET("/subscription", api.GetSubscription(userService))
        user.PUT("/subscription", api.UpdateSubscription(userService))
      }

      reviews := protected.Group("/reviews")
      {
        reviews.POST("/", api.CreateReview(reviewService))
        reviews.GET("/:id", api.GetReview(reviewService))
        reviews.PUT("/:id", api.UpdateReview(reviewService))
        reviews.DELETE("/:id", api.DeleteReview(reviewService))
        reviews.GET("/", api.GetUserReviews(reviewService))
        reviews.POST("/:id/generate", api.GenerateReviewContent(reviewService))
        reviews.GET("/generation/:taskId", api.CheckGenerationStatus(reviewService))
      }
    }
  }

  log.Printf("Starting server on %s", cfg.Server.Address)
  if err := router.Run(cfg.Server.Address); err != nil {
    log.Fatalf("Failed to start server: %v", err)
  }
}
