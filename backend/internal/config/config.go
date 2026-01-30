package config

import "os"

type Config struct {
  Server struct {
    Address string
  }
  Database struct {
    URI  string
    Name string
  }
  PubMed struct {
    BaseURL string
    APIKey  string
  }
  JWT struct {
    Secret     string
    Expiration string
  }
}

func Load() (*Config, error) {
  cfg := &Config{}
  cfg.Server.Address = getEnv("SERVER_ADDRESS", ":8080")
  cfg.Database.URI = getEnv("DATABASE_URI", "mongodb://localhost:27017")
  cfg.Database.Name = getEnv("DATABASE_NAME", "xiaozhi_scientific")
  cfg.PubMed.BaseURL = getEnv("PUBMED_BASE_URL", "https://eutils.ncbi.nlm.nih.gov/entrez/eutils")
  cfg.PubMed.APIKey = getEnv("PUBMED_API_KEY", "")
  cfg.JWT.Secret = getEnv("JWT_SECRET", "change-me")
  cfg.JWT.Expiration = getEnv("JWT_EXPIRATION", "24h")

  return cfg, nil
}

func getEnv(key string, fallback string) string {
  value := os.Getenv(key)
  if value == "" {
    return fallback
  }
  return value
}
