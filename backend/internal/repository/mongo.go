package repository

import "errors"

type MongoRepository struct{}

func NewMongoRepository(uri string, name string) (*MongoRepository, error) {
  if uri == "" || name == "" {
    return nil, errors.New("missing mongo configuration")
  }
  return &MongoRepository{}, nil
}

func (repo *MongoRepository) Close() error {
  return nil
}
