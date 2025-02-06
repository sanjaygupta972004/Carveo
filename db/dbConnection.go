package db

import (
	"context"
	"fmt"
	"log"
	"os"
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

func connectPGX(cfg *DBConfig) *pgxpool.Pool {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbName, cfg.SSLMode,
	)
	PGXPool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Printf("Unable to connect to pgx database : %v\n", err)
		os.Exit(1)

	}
	var greeting string
	err = PGXPool.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	return PGXPool
}

func connectGORMDB(cfg *DBConfig) *gorm.DB {
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
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), gormConfig)
	if err != nil {
		log.Printf("Unable to open gorm DB :%v", err)
		os.Exit(1)
	}

	return gormDB
}

func ConnectDB(cfg *DBConfig) error {
	// initialized pgx native postgres driver
	PGXPool := connectPGX(cfg)
	defer PGXPool.Close()

	if err := PGXPool.Ping(context.Background()); err != nil {
		log.Fatal("pgx connection failed:", err)
		return err
	}

	// connected gorm db
	gormDB := connectGORMDB(cfg)
	sqlDB, err := gormDB.DB()
	if err != nil {
		log.Printf("Unable to open sqlDB from gorm :%v", err)
		return err
	}

	if sqlDB.Ping() == nil {
		DB = gormDB
	} else {
		log.Fatalf("Error connecting to database: %v", err)
		return err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("Successfully connected to postgres with gorm and pgx")

	return nil
}

func DisConnectDB() error {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			log.Fatalf("Error fetching SQL DB from GORM: %v", err)
			return err
		}

		err = sqlDB.Close()
		if err != nil {
			log.Fatalf("Error closing SQL DB connection: %v", err)
			return err
		}

	}

	return nil
}
