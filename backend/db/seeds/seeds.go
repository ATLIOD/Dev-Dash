package seeds

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

// SeedDatabase is the entry point for seeding dummy data
func SeedDatabase(pool *pgxpool.Pool) error {
	log.Println("Seeding database...")
	ctx := context.Background()

	if err := seedUsers(ctx, pool); err != nil {
		return fmt.Errorf("failed to seed data: %w", err)
	}

	log.Println("Seeding completed successfully!")
	return nil
}

func seedUsers(ctx context.Context, pool *pgxpool.Pool) error {
	password, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)

	users := []struct {
		Name         string
		Email        string
		PasswordHash string
	}{
		{"Admin User", "admin@devdash.com", string(password)},
		{"John Doe", "john@example.com", string(password)},
		{"Jane Smith", "jane@example.com", string(password)},
	}

	log.Println("Seeding database...")

	for _, u := range users {
		_, err := pool.Exec(ctx, `
			INSERT INTO users (name, email, password_hash) 
			VALUES ($1, $2, $3)
			ON CONFLICT (email) DO NOTHING`,
			u.Name, u.Email, u.PasswordHash,
		)
		if err != nil {
			return fmt.Errorf("failed to insert user %s: %w", u.Email, err)
		}
	}

	return nil
}
