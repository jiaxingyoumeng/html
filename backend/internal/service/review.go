package service

import "context"

type ReviewService struct {
  repo          interface{}
  pubmedBaseURL string
  pubmedAPIKey  string
  storageDir    string
  tasks         *writingTaskStore
}

func NewReviewService(repo interface{}, pubmedBaseURL string, pubmedAPIKey string, storageDir string) *ReviewService {
  return &ReviewService{
    repo:          repo,
    pubmedBaseURL: pubmedBaseURL,
    pubmedAPIKey:  pubmedAPIKey,
    storageDir:    storageDir,
    tasks:         newWritingTaskStore(),
  }
}

func (service *ReviewService) CreateReview(ctx context.Context, payload map[string]any) (map[string]any, error) {
  return payload, nil
}

func (service *ReviewService) GetReview(ctx context.Context, id string) (map[string]any, error) {
  return map[string]any{"id": id}, nil
}

func (service *ReviewService) UpdateReview(ctx context.Context, id string, payload map[string]any) (map[string]any, error) {
  payload["id"] = id
  return payload, nil
}

func (service *ReviewService) DeleteReview(ctx context.Context, id string) error {
  return nil
}

func (service *ReviewService) GetUserReviews(ctx context.Context, page int, limit int) ([]map[string]any, error) {
  return []map[string]any{}, nil
}

func (service *ReviewService) GenerateReviewContent(ctx context.Context, reviewID string) (map[string]any, error) {
  return map[string]any{"reviewId": reviewID, "status": "processing"}, nil
}

func (service *ReviewService) CheckGenerationStatus(ctx context.Context, taskID string) (map[string]any, error) {
  return map[string]any{"taskId": taskID, "status": "processing"}, nil
}

 
