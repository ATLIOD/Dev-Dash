package main

import (
	"DevDash/db"
	"DevDash/internal/api"
	"DevDash/internal/api/handlers"
	"DevDash/internal/config"
	"DevDash/internal/repositories"
	"DevDash/internal/services"
	"log"
	"net/http"
)

func main() {
	cfg := config.Load()
	database, err := db.OpenDB(cfg.DB)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	repos := repositories.New(database)
	svcs := services.New(repos)
	h := handlers.New(svcs)

	router := api.NewRouter(h)

	log.Printf("Server starting on :%s", "8080")
	if err := http.ListenAndServe(":"+"8080", router); err != nil {
		log.Fatal(err)
	}
}
