package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	DB DBConfig
}

type DBConfig struct {
	Dsn             string
	MaxConns        int
	MaxConnIdleTime time.Duration
	MinConns        int
}

func Load() *Config {
	if os.Getenv("APP_ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Println("No .env fie found, continuin..")
		}
	}
	log.Println("environment: ", os.Getenv("APP_ENV"))

	pgDSN := os.Getenv("DATABASE_URL")
	MaxConns, err := strconv.Atoi(os.Getenv("DATABASE_MAX_CONNS"))
	if err != nil {
		MaxConns = 10
		log.Println("DATABASE_MAX_CONNS:", err)
	}
	MaxConnIdleTime, err := time.ParseDuration(os.Getenv("DATABASE_MAX_CONN_IDLE_TIME"))
	if err != nil {
		MaxConnIdleTime = 20
		log.Println("DATABASE_MAX_CONN_IDLE_TIME:", err)
	}
	MinConns, err := strconv.Atoi(os.Getenv("DATABASE_MIN_CONNS"))
	if err != nil {
		MinConns = 10
		log.Println("DATABASE_MIN_CONNS:", err)
	}

	return &Config{
		DB: DBConfig{
			Dsn:             pgDSN,
			MaxConns:        MaxConns,
			MaxConnIdleTime: MaxConnIdleTime,
			MinConns:        MinConns,
		},
	}
}
