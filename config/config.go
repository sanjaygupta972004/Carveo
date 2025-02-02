package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/joho/godotenv"
)

// Global config variable with mutex for safe concurrent access
var (
	config     Config
	configOnce sync.Once
	isLoaded   bool
	mu         sync.RWMutex
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
	return MailgunConfig{
		Domain:        os.Getenv("MAILGUN_DOMAIN"),
		PrivateAPIKey: os.Getenv("MAILGUN_API_KEY"),
		SenderEmail:   os.Getenv("MAILGUN_SENDER_EMAIL"),
	}, nil
}

func LoadConfig() error {
	var loadErr error
	configOnce.Do(func() {
		env := os.Getenv("APP_ENV")
		if env == "" {
			env = "dev"
		}

		if err := LoadEnvFile(env); err != nil {
			loadErr = fmt.Errorf("failed to load environment file: %w", err)
			return
		}

		mailgunConfig, err := loadENVForMailgun()
		if err != nil {
			loadErr = fmt.Errorf("failed to load mailgun config: %w", err)
			return
		}

		mu.Lock()
		defer mu.Unlock()

		config = Config{
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

		// Validate required fields
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
				loadErr = fmt.Errorf("missing required environment variable: %s", key)
				return
			}
		}

		isLoaded = true
	})

	return loadErr
}

func GetConfig() *Config {
	mu.RLock()
	defer mu.RUnlock()

	if !isLoaded {
		log.Fatal("Config is not initialized. Call LoadConfig() first.")
	}
	return &config
}
