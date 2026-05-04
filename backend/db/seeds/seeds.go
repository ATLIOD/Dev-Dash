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
		return fmt.Errorf("failed to seed users: %w", err)
	}

	if err := seedProjects(ctx, pool); err != nil {
		return fmt.Errorf("failed to seed projects: %w", err)
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

	log.Println("Seeding users...")
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

func seedProjects(ctx context.Context, pool *pgxpool.Pool) error {
	var userIDs []int64
	rows, err := pool.Query(ctx, "SELECT id FROM users LIMIT 3")
	if err != nil {
		return fmt.Errorf("failed to fetch users for seeding projects: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return fmt.Errorf("failed to scan user id: %w", err)
		}
		userIDs = append(userIDs, id)
	}

	if len(userIDs) == 0 {
		log.Println("No users found to associate with projects. Skipping project seeding.")
		return nil
	}

	projects := []struct {
		Name          string
		Description   string
		Status        string
		Stack         string
		RepositoryURL string
		DeploymentURL string
		UserID        int64
	}{
		{
			Name:          "Dev-Dash Backend",
			Description:   "Go backend for developer dashboard",
			Status:        "active",
			Stack:         "Go, PostgreSQL, pgx",
			RepositoryURL: "https://github.com/example/dev-dash",
			DeploymentURL: "https://api.dev-dash.com",
			UserID:        userIDs[0],
		},
		{
			Name:          "Dev-Dash Frontend",
			Description:   "React frontend for developer dashboard",
			Status:        "active",
			Stack:         "React, TypeScript, Tailwind",
			RepositoryURL: "https://github.com/example/dev-dash-ui",
			DeploymentURL: "https://dev-dash.com",
			UserID:        userIDs[0],
		},
	}

	if len(userIDs) > 1 {
		projects = append(projects, struct {
			Name          string
			Description   string
			Status        string
			Stack         string
			RepositoryURL string
			DeploymentURL string
			UserID        int64
		}{
			Name:          "Personal Portfolio",
			Description:   "A static portfolio site",
			Status:        "completed",
			Stack:         "HTML, CSS, JS",
			RepositoryURL: "https://github.com/johndoe/portfolio",
			DeploymentURL: "https://johndoe.me",
			UserID:        userIDs[1],
		})
	}

	log.Println("Seeding projects...")
	for _, p := range projects {
		_, err := pool.Exec(ctx, `
			INSERT INTO projects (name, description, status, stack, repository_url, deployment_url, user_id) 
			VALUES ($1, $2, $3, $4, $5, $6, $7)
			ON CONFLICT DO NOTHING`,
			p.Name, p.Description, p.Status, p.Stack, p.RepositoryURL, p.DeploymentURL, p.UserID,
		)
		if err != nil {
			return fmt.Errorf("failed to insert project %s: %w", p.Name, err)
		}
	}

	return nil
}
