package config

import (
	"DevDash/internal/api/middleware"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

type Config struct {
	DB         DBConfig
	CorsConfig *cors.Cors
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

	allowedOrigins := strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",")
	allowedMethods := strings.Split(os.Getenv("ALLOWED_METHODS"), ",")
	allowedHeaders := strings.Split(os.Getenv("ALLOWED_HEADERS"), ",")
	exposedHeaders := strings.Split(os.Getenv("EXPOSED_HEADERS"), ",")

	corsConfig := middleware.GetCorsConfig(allowedOrigins, allowedMethods, allowedHeaders, exposedHeaders)

	return &Config{
		DB: DBConfig{
			Dsn:             pgDSN,
			MaxConns:        MaxConns,
			MaxConnIdleTime: MaxConnIdleTime,
			MinConns:        MinConns,
			Seed:            Seed,
		},
		CorsConfig: corsConfig,
	}
}
