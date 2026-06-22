package main

import (
	"API/internal/config"
	"API/internal/database"
	"API/internal/router"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	cfg, err := config.Load()
	if err != nil {
		log.Fatal("error loading config")
	}

	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("error connecting to database")
	}
	defer db.Close()

	mux := router.SetupRoutes(db)
	log.Printf("Server started on %s", cfg.ServerAddr)
	if err := http.ListenAndServe(cfg.ServerAddr, mux); err != nil {
		log.Fatal(err)
	}

}
