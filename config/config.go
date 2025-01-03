package config

import (
	"log"
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

func LoadConfig() (Config, error) {

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}
	envPath := filepath.Join("..", "config", "env", env+".env")
	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("Error coming while loading env file : %v", err)
		return Config{}, err
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

	if config.DbHost == "" && config.DbName == "" && config.DbPassword == "" && config.DbUser == "" && config.DbPort == "" {
		log.Fatalf("Essential credential not found from env files")
		return Config{}, err
	}
	return config, nil
}
