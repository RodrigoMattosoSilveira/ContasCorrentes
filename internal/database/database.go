package database

import (
	"log"

	"github.com/RodrigoMattosoSilveira/ContasCorrentes/internal/config"
	"github.com/RodrigoMattosoSilveira/ContasCorrentes/internal/modules/associates"
	"github.com/RodrigoMattosoSilveira/ContasCorrentes/internal/modules/users"

	"fmt"
	"log/slog"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDatabase(cfg *config.Config) (*gorm.DB, error) {
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		panic(fmt.Sprintf("DBInit: Invalid DB_NAME environment %s variable", dbName))
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
	err = db.AutoMigrate(&associates.Associate{})
	if err != nil {
		return err
	}
	log.Println("Database migrations completed.")
	return nil
}
