package service

import (
  "context"
  "errors"
)

type AuthService struct {
  repo          interface{}
  jwtSecret     string
  jwtExpiration string
}

type TokenClaims map[string]any

type AuthUser struct {
  ID       string `json:"id"`
  Username string `json:"username"`
  Email    string `json:"email"`
}

func NewAuthService(repo interface{}, jwtSecret string, jwtExpiration string) *AuthService {
  return &AuthService{repo: repo, jwtSecret: jwtSecret, jwtExpiration: jwtExpiration}
}

func (service *AuthService) Login(ctx context.Context, username string, password string) (string, *AuthUser, error) {
  return "", nil, errors.New("not implemented")
}

func (service *AuthService) Register(ctx context.Context, user any) (string, *AuthUser, error) {
  return "", nil, errors.New("not implemented")
}

func (service *AuthService) ValidateToken(token string) (TokenClaims, error) {
  return nil, errors.New("not implemented")
}

func (service *AuthService) SendPasswordResetEmail(ctx context.Context, email string) error {
  return errors.New("not implemented")
}
