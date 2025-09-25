package database

import (
	"log"

	"github.com/RodrigoMattosoSilveira/ContasCorrentes/internal/config"
	"github.com/RodrigoMattosoSilveira/ContasCorrentes/internal/modules/users"

	"gorm.io/driver/sqlite"
	"fmt"
	"gorm.io/gorm"
	"log/slog"
	"os"
)

func InitDatabase(cfg *config.Config) (*gorm.DB, error) {
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		panic(fmt.Sprintf("DBInit: Invalid DB_NAME environment%svariable", " "))
	}

	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect to SQLite %s database", dbName))
	}
	
	slog.Info(fmt.Sprintf("DBInit: connected successfully to %s", dbName))
	return db, nil
}

func RunMigrations(db *gorm.DB) error {
	log.Println("Running database migrations...")
	err := db.AutoMigrate(&users.User{})
	if err != nil {
		return err
	}
	log.Println("Database migrations completed.")
	return nil
}
