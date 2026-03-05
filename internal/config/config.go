package config

import (
	"errors"
	"os"
)

const (
	EnvDevelopment = "development"
	EnvProduction  = "production"
)

type Config struct {
	Env             string
	ServerAddr      string
	DatabaseURLDev  string
	DatabaseURLProd string
}

func Load() Config {
	return Config{
		Env:             getEnv("APP_ENV", EnvDevelopment),
		ServerAddr:      getEnv("SERVER_ADDR", ":8080"),
		DatabaseURLDev:  getEnv("DATABASE_URL_DEV", "postgres://ownp:your_secure_password_here@localhost:5432/ownp_db?sslmode=disable"),
		DatabaseURLProd: os.Getenv("DATABASE_URL_PROD"),
	}
}

func (c Config) ActiveDatabaseURL() (string, error) {
	if c.Env == EnvProduction {
		if c.DatabaseURLProd == "" {
			return "", errors.New("DATABASE_URL_PROD is required when APP_ENV=production")
		}
		return c.DatabaseURLProd, nil
	}

	if c.DatabaseURLDev == "" {
		return "", errors.New("DATABASE_URL_DEV is required")
	}
	return c.DatabaseURLDev, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
