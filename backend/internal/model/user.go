package model

import (
  "time"

  "go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
  ID           primitive.ObjectID `bson:"_id,omitempty"`
  Username     string             `bson:"username" binding:"required,min=3,max=20"`
  Email        string             `bson:"email" binding:"required,email"`
  Phone        string             `bson:"phone" binding:"required,len=11"`
  PasswordHash string             `bson:"password_hash" binding:"required,min=6"`
  Avatar       string             `bson:"avatar"`
  Role         string             `bson:"role"`
  Subscription *Subscription      `bson:"subscription"`
  CreatedAt    time.Time          `bson:"created_at"`
  UpdatedAt    time.Time          `bson:"updated_at"`
}

type Subscription struct {
  Level        string    `bson:"level"`
  ExpiryDate   time.Time `bson:"expiry_date"`
  TotalReviews int       `bson:"total_reviews"`
  UsedReviews  int       `bson:"used_reviews"`
}
