package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	DbUser        string
	DbHost        string
	DbPort        string
	DbPassword    string
	DbName        string
	JwtSecret     string
	SSLMode       string
	Port          string
	MailgunConfig MailgunConfig
}

type MailgunConfig struct {
	Domain        string
	PrivateAPIKey string
	SenderEmail   string
}

func LoadEnvFile(env string) error {
	envPath := filepath.Join("env", env+".env")
	fmt.Println("Loading environment file: ", envPath)
	return godotenv.Load(envPath)
}

func loadENVForMailgun() (MailgunConfig, error) {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}

	err := LoadEnvFile(env)
	if err != nil {
		return MailgunConfig{}, fmt.Errorf("failed to load environment file: %w", err)
	}

	config := MailgunConfig{
		Domain:        os.Getenv("MAILGUN_DOMAIN"),
		PrivateAPIKey: os.Getenv("MAILGUN_PRIVATE_API_KEY"),
		SenderEmail:   os.Getenv("MAILGUN_SENDER_EMAIL"),
	}

	requiredFields := map[string]string{
		"MAILGUN_DOMAIN":          config.Domain,
		"MAILGUN_PRIVATE_API_KEY": config.PrivateAPIKey,
		"MAILGUN_SENDER_EMAIL":    config.SenderEmail,
	}

	for key, value := range requiredFields {
		if value == "" {
			return MailgunConfig{}, fmt.Errorf("missing required environment variable: %s", key)
		}
	}

	return config, nil
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
	mailgunConfig, err := loadENVForMailgun()
	if err != nil {
		return Config{}, fmt.Errorf("failed to load mailgun config: %w", err)
	}
	config := Config{
		DbUser:        os.Getenv("DB_USER"),
		DbHost:        os.Getenv("DB_HOST"),
		DbPort:        os.Getenv("DB_PORT"),
		DbPassword:    os.Getenv("DB_PASSWORD"),
		DbName:        os.Getenv("DB_NAME"),
		SSLMode:       os.Getenv("DB_SSLMODE"),
		JwtSecret:     os.Getenv("JWT_SECRET"),
		Port:          os.Getenv("PORT"),
		MailgunConfig: mailgunConfig,
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
