package db

import (
	"context"
	"fmt"
	"log"

	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	once    sync.Once
	DB      *gorm.DB
	PGXPool *pgxpool.Pool
)

type DBConfig struct {
	DbUser     string
	DbHost     string
	DbPort     string
	DbPassword string
	DbName     string
	SSLMode    string
}

func connectPGX(cfg *DBConfig) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbName, cfg.SSLMode,
	)
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Printf("Unable to connect to pgx database: %v\n", err)
		return nil, err
	}

	// Test query to verify connection
	var greeting string
	err = pool.QueryRow(context.Background(), "SELECT 'Hello, world!'").Scan(&greeting)
	if err != nil {
		log.Printf("QueryRow failed: %v\n", err)
		pool.Close()
		return nil, err
	}

	return pool, nil
}

func connectGORMDB(cfg *DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbName, cfg.SSLMode,
	)

	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn}), gormConfig)
	if err != nil {
		log.Printf("Unable to open gorm DB: %v", err)
		return nil, err
	}

	return gormDB, nil
}

func ConnectDB(cfg *DBConfig) error {
	var err error

	once.Do(func() {
		PGXPool, err = connectPGX(cfg)
		if err != nil {
			log.Fatalf("pgx connection failed: %v", err)
			return
		}

		// Verify pgx connection
		if err = PGXPool.Ping(context.Background()); err != nil {
			log.Fatalf("pgx ping failed: %v", err)
			PGXPool.Close()
			return
		}

		DB, err = connectGORMDB(cfg)
		if err != nil {
			log.Fatalf("gorm connection failed: %v", err)
			PGXPool.Close()
			return
		}

		sqlDB, err := DB.DB()
		if err != nil {
			log.Fatalf("Error fetching SQL DB from GORM: %v", err)
			PGXPool.Close()
			return
		}

		// Configure connection pooling
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)

		log.Println("Successfully connected to Postgres with GORM and PGX")
	})

	return err
}

func DisConnectDB() error {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			log.Printf("Error fetching SQL DB from GORM: %v", err)
			return err
		}

		err = sqlDB.Close()
		if err != nil {
			log.Printf("Error closing SQL DB connection: %v", err)
			return err
		}
	}

	if PGXPool != nil {
		PGXPool.Close()
		log.Println("PGXPool connection closed")
	}

	return nil
}
