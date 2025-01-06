package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	DbUser     string
	DbHost     string
	DbPort     string
	DbPassword string
	DbName     string
	JwtSecret  string
	SSLMode    string
	Port       string
}

func LoadEnvFile(env string) error {
	envPath := filepath.Join("env", env+".env")
	fmt.Println("Loading environment file: ", envPath)
	return godotenv.Load(envPath)
}

func LoadConfig() (Config, error) {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}

	err := LoadEnvFile(env)
	if err != nil {
		return Config{}, fmt.Errorf("failed to load environment file: %w", err)
	}

	config := Config{
		DbUser:     os.Getenv("DB_USER"),
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbName:     os.Getenv("DB_NAME"),
		SSLMode:    os.Getenv("DB_SSLMODE"),
		JwtSecret:  os.Getenv("JWT_SECRET"),
		Port:       os.Getenv("PORT"),
	}

	requiredFields := map[string]string{
		"DB_USER":     config.DbUser,
		"DB_HOST":     config.DbHost,
		"DB_PORT":     config.DbPort,
		"DB_PASSWORD": config.DbPassword,
		"DB_NAME":     config.DbName,
		"JWT_SECRET":  config.JwtSecret,
		"PORT":        config.Port,
	}

	for key, value := range requiredFields {
		if value == "" {
			return Config{}, fmt.Errorf("missing required environment variable: %s", key)
		}
	}

	return config, nil
}
