package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/PrinceM13/knowledge-hub-api/internal/config"

	_ "github.com/lib/pq" // 👈 REQUIRED
)

var DB *sql.DB

func Connect(cfg *config.Config) error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBName,
	)

	var err error

	DB, err = sql.Open("postgres", dsn)

	if err != nil {
		return err
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	DB.SetConnMaxLifetime(time.Hour)

	if err := DB.Ping(); err != nil {
		return err
	}

	log.Println("✅ PostgreSQL connected")
	return nil
}

func Close() error {
	if DB != nil {
		log.Println("🛑 Closing database connection...")
		return DB.Close()
	}
	return nil
}
