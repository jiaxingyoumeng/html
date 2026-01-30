package model

import "time"

type Review struct {
  ID                string             `bson:"_id,omitempty"`
  UserID            string             `bson:"user_id"`
  Title             string             `bson:"title"`
  Content           string             `bson:"content"`
  Topic             string             `bson:"topic"`
  ImpactFactorRange *ImpactFactorRange `bson:"impact_factor_range"`
  YearRange         *YearRange         `bson:"year_range"`
  Status            string             `bson:"status"`
  Progress          int                `bson:"progress"`
  References        []Reference        `bson:"references"`
  CreatedAt         time.Time          `bson:"created_at"`
  UpdatedAt         time.Time          `bson:"updated_at"`
}

type ImpactFactorRange struct {
  Min int `bson:"min"`
  Max int `bson:"max"`
}

type YearRange struct {
  StartYear int `bson:"start_year"`
  EndYear   int `bson:"end_year"`
}

type Reference struct {
  ID       string `bson:"id"`
  Title    string `bson:"title"`
  Authors  string `bson:"authors"`
  Journal  string `bson:"journal"`
  Year     int    `bson:"year"`
  Volume   string `bson:"volume"`
  Issue    string `bson:"issue"`
  Pages    string `bson:"pages"`
  Doi      string `bson:"doi"`
  Citation string `bson:"citation"`
}
