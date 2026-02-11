package testutil

import (
	"database/sql"
	"testing"
	"time"

	"github.com/PrinceM13/knowledge-hub-api/internal/user"
)

// CreateTestUser creates a user in the test database and returns it
func CreateTestUser(t *testing.T, db *sql.DB, email, name string) *user.User {
	t.Helper()

	query := `
		INSERT INTO users (email, name, created_at)
		VALUES ($1, $2, $3)
		RETURNING id, email, name, created_at
	`

	u := &user.User{}
	err := db.QueryRow(query, email, name, time.Now()).Scan(
		&u.ID,
		&u.Email,
		&u.Name,
		&u.CreatedAt,
	)

	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}

	return u
}

// CreateMultipleTestUsers creates multiple test users
func CreateMultipleTestUsers(t *testing.T, db *sql.DB, count int) []*user.User {
	t.Helper()

	users := make([]*user.User, count)
	for i := 0; i < count; i++ {
		users[i] = CreateTestUser(
			t,
			db,
			"user"+string(rune(i))+"@test.com",
			"Test User "+string(rune(i)),
		)
	}

	return users
}
