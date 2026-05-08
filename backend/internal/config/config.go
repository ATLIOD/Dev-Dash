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
	Seed            bool
}

func Load() *Config {
	if os.Getenv("APP_ENV") != "production" {
		_ = godotenv.Load()
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
		MaxConnIdleTime = 20 * time.Second
		log.Println("DATABASE_MAX_CONN_IDLE_TIME:", err)
	}
	MinConns, err := strconv.Atoi(os.Getenv("DATABASE_MIN_CONNS"))
	if err != nil {
		MinConns = 10
		log.Println("DATABASE_MIN_CONNS:", err)
	}

	seedStr := os.Getenv("SEED")
	Seed, err := strconv.ParseBool(seedStr)
	if err != nil {
		Seed = false
		log.Println("SEED:", err)
	}

	return &Config{
		DB: DBConfig{
			Dsn:             pgDSN,
			MaxConns:        MaxConns,
			MaxConnIdleTime: MaxConnIdleTime,
			MinConns:        MinConns,
			Seed:            Seed,
		},
	}
}
