package main

import (
	"log"

	"github.com/RodrigoMattosoSilveira/ContasCorrentes/internal/config"
	"github.com/RodrigoMattosoSilveira/ContasCorrentes/internal/database"
	"github.com/RodrigoMattosoSilveira/ContasCorrentes/internal/server"
	"github.com/RodrigoMattosoSilveira/ContasCorrentes/internal/validator"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	db, err := database.InitDatabase(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	if err := database.RunMigrations(db); err != nil {
		log.Fatalf("Failed to run database migrations: %v", err)
	}

	validator.InitValidator()

	app := server.NewServer(cfg, db)
	log.Printf("Starting server on port %s...", cfg.Port)
	if err := app.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
