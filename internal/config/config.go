// internal/config/config.go
package config

import "os"

type Config struct {
    Address     string
    Environment string
}

func New() *Config {
    port := getEnv("PORT", "8000")
    return &Config{
        Address:     ":" + port,
        Environment: getEnv("ENV", "development"),
    }
}

func getEnv(key, fallback string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return fallback
}