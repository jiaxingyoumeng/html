package config

import "github.com/spf13/viper"

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
  viper.SetDefault("server.address", ":8080")
  viper.SetDefault("database.uri", "mongodb://localhost:27017")
  viper.SetDefault("database.name", "xiaozhi_scientific")
  viper.SetDefault("pubmed.base_url", "https://eutils.ncbi.nlm.nih.gov/entrez/eutils")
  viper.SetDefault("pubmed.api_key", "")
  viper.SetDefault("jwt.secret", "change-me")
  viper.SetDefault("jwt.expiration", "24h")

  viper.AutomaticEnv()

  cfg := &Config{}
  cfg.Server.Address = viper.GetString("server.address")
  cfg.Database.URI = viper.GetString("database.uri")
  cfg.Database.Name = viper.GetString("database.name")
  cfg.PubMed.BaseURL = viper.GetString("pubmed.base_url")
  cfg.PubMed.APIKey = viper.GetString("pubmed.api_key")
  cfg.JWT.Secret = viper.GetString("jwt.secret")
  cfg.JWT.Expiration = viper.GetString("jwt.expiration")

  return cfg, nil
}
