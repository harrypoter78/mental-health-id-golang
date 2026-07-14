package config

import (
    "fmt"
    "os"

    "github.com/joho/godotenv"
)

type Config struct {
    DBHost        string
    DBPort        string
    DBUser        string
    DBPassword    string
    DBName        string
    SessionSecret string
}

func LoadConfig() (*Config, error) {
    _ = godotenv.Load()

    cfg := &Config{
        DBHost:        getEnv("DB_HOST", "127.0.0.1"),
        DBPort:        getEnv("DB_PORT", "3306"),
        DBUser:        getEnv("DB_USER", "root"),
        DBPassword:    getEnv("DB_PASSWORD", ""),
        DBName:        getEnv("DB_NAME", "gangguanmental"),
        SessionSecret: getEnv("SESSION_SECRET", ""),
    }

    if cfg.SessionSecret == "" {
        return nil, fmt.Errorf("SESSION_SECRET must be set")
    }

    return cfg, nil
}

func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}
