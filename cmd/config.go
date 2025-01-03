package main

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
	TimeZone   string
	Port       string
}

func LoadConfig() (Config, error) {

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}
	envPath := filepath.Join(env + ".env")
	log.Println("absolute path ", envPath)
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
		JwtSecret:  os.Getenv("JWT_SECRET"),
		SSLMode:    os.Getenv("DB_SSLMODE"),
		TimeZone:   os.Getenv("DB_TIMEZONE"),
		Port:       os.Getenv("PORT"),
	}

	if config.DbHost == "" && config.DbName == "" && config.DbPassword == "" && config.DbUser == "" && config.DbPort == "" && config.Port == "" && config.SSLMode == "" && config.TimeZone == "" {
		log.Fatalf("Essential credential not found from env files")
		return Config{}, err
	}
	return config, nil
}
