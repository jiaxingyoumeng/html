package service

import "context"

type UserService struct {
  repo interface{}
}

func NewUserService(repo interface{}) *UserService {
  return &UserService{repo: repo}
}

func (service *UserService) GetProfile(ctx context.Context, userID string) (map[string]any, error) {
  return map[string]any{}, nil
}

func (service *UserService) UpdateProfile(ctx context.Context, userID string, data map[string]any) (map[string]any, error) {
  return data, nil
}

func (service *UserService) GetSubscription(ctx context.Context, userID string) (map[string]any, error) {
  return map[string]any{}, nil
}

func (service *UserService) UpdateSubscription(ctx context.Context, userID string, planID string) (map[string]any, error) {
  return map[string]any{"planId": planID}, nil
}
