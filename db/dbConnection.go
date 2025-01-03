package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

var (
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
	TimeZone   string
}

func ConnectDB(cfg *DBConfig) error {

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s timezone=%s", cfg.DbUser, cfg.DbPassword, cfg.DbName, cfg.DbHost, cfg.DbPort, cfg.SSLMode, cfg.TimeZone)

	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatalf("Unable to perse pgx config ; %v", err)
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	PGXPool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		log.Fatalf("Unable to initialize pgx pool ; %v", err)
		return err
	}

	err = PGXPool.Ping(ctx)
	if err != nil {
		log.Fatalf("Unable to connect to database with pgx: %v", err)
	}

	log.Println("Successfully connected to postgres with pgx :", PGXPool)

	sqlDB := stdlib.OpenDB(*PGXPool.Config().ConnConfig)
	defer sqlDB.Close()

	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}

	DB, err = gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), gormConfig)
	if err != nil {
		log.Fatalf("Failed to initialize GORM: %v", err)
		return err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("Successfully connected to postgres with gorm ")
	return nil
}

func DisConnectDB() error {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			log.Fatalf("failed to get sqlDB form gorm DB :%v", err)
			return err
		}
		err = sqlDB.Close()
		if err != nil {
			log.Fatalf("error closing database :%v", err)
			return err
		}
		log.Println("Database connection closed successfully")

	}
	return nil
}
