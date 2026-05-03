package utils

import (
	"DevDash/db"
	"DevDash/internal/config"
	"DevDash/internal/repositories"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func Setup() (*repositories.Repository, func()) {
	LoadEnv()
	cfg := config.Load()

	database, err := db.OpenDB(cfg.DB)
	if err != nil {
		log.Fatal(err)
	}
	repo := repositories.New(database)
	return repo, database.Close
}

func LoadEnv() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	path := cwd
	for {
		envPath := filepath.Join(path, ".env")
		if _, err := os.Stat(envPath); err == nil {
			err := godotenv.Load(envPath)
			if err != nil {
				log.Fatalf("Error loading .env file from %s: %v", envPath, err)
			}
			return
		}

		parent := filepath.Dir(path)
		if parent == path {
			break
		}
		path = parent
	}
	log.Println("Warning: .env file not found in any parent directory")
}
