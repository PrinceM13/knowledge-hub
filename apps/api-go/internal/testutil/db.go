package testutil

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

// TestDBConfig holds test database configuration
type TestDBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// GetTestDBConfig loads test database configuration from environment
func GetTestDBConfig() *TestDBConfig {
	return &TestDBConfig{
		Host:     getEnv("POSTGRES_HOST", "localhost"),
		Port:     getEnv("POSTGRES_PORT", "5433"),
		User:     getEnv("POSTGRES_USER", "kh_user"),
		Password: getEnv("POSTGRES_PASSWORD", "kh_password"),
		DBName:   getEnv("POSTGRES_DB", "knowledge_hub_test"),
	}
}

// SetupTestDB creates a test database connection
func SetupTestDB(t *testing.T) *sql.DB {
	t.Helper()

	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	cfg := GetTestDBConfig()
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	// Verify connection
	if err := db.Ping(); err != nil {
		t.Fatalf("Failed to ping test database: %v", err)
	}

	// Clean up function
	t.Cleanup(func() {
		db.Close()
	})

	return db
}

// CleanupTestDB truncates all tables for a clean test environment
func CleanupTestDB(t *testing.T, db *sql.DB, tables ...string) {
	t.Helper()

	for _, table := range tables {
		query := fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE", table)
		if _, err := db.Exec(query); err != nil {
			t.Logf("Warning: Failed to truncate table %s: %v", table, err)
		}
	}
}

// TruncateTable is a convenience function to truncate a single table
func TruncateTable(t *testing.T, db *sql.DB, table string) {
	t.Helper()
	CleanupTestDB(t, db, table)
}

// getEnv gets environment variable with a fallback default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
