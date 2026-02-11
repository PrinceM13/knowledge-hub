package userdb

import (
	"context"
	"database/sql"
	"testing"

	"github.com/PrinceM13/knowledge-hub-api/internal/testutil"
	"github.com/PrinceM13/knowledge-hub-api/internal/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIntegrationPostgresRepository_Create(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	db := testutil.SetupTestDB(t)
	defer testutil.TruncateTable(t, db, "users")

	repo := NewPostgresRepository(db)
	ctx := context.Background()

	tests := []struct {
		name    string
		user    *user.User
		wantErr bool
	}{
		{
			name: "create user successfully",
			user: &user.User{
				Email: "test@example.com",
				Name:  "Test User",
			},
			wantErr: false,
		},
		{
			name: "create another user",
			user: &user.User{
				Email: "another@example.com",
				Name:  "Another User",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// act
			err := repo.Create(ctx, tt.user)

			// assert
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotZero(t, tt.user.ID)
				assert.NotZero(t, tt.user.CreatedAt)
				assert.Equal(t, tt.user.Email, tt.user.Email)
				assert.Equal(t, tt.user.Name, tt.user.Name)
			}
		})
	}
}

func TestIntegrationPostgresRepository_Create_DuplicateEmail(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	db := testutil.SetupTestDB(t)
	defer testutil.TruncateTable(t, db, "users")

	repo := NewPostgresRepository(db)
	ctx := context.Background()

	// arrange - create first user
	user1 := &user.User{
		Email: "duplicate@example.com",
		Name:  "First User",
	}
	err := repo.Create(ctx, user1)
	require.NoError(t, err)

	// act - try to create user with same email
	user2 := &user.User{
		Email: "duplicate@example.com",
		Name:  "Second User",
	}
	err = repo.Create(ctx, user2)

	// assert - should fail due to unique constraint
	assert.Error(t, err)
}

func TestIntegrationPostgresRepository_FindByID(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	db := testutil.SetupTestDB(t)
	defer testutil.TruncateTable(t, db, "users")

	repo := NewPostgresRepository(db)
	ctx := context.Background()

	// arrange - create a test user
	createdUser := testutil.CreateTestUser(t, db, "findme@example.com", "Find Me")

	tests := []struct {
		name    string
		userID  int64
		wantErr bool
		errType error
	}{
		{
			name:    "find existing user",
			userID:  createdUser.ID,
			wantErr: false,
		},
		{
			name:    "user not found",
			userID:  99999,
			wantErr: true,
			errType: sql.ErrNoRows,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// act
			foundUser, err := repo.FindByID(ctx, tt.userID)

			// assert
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, foundUser)
				if tt.errType != nil {
					assert.Equal(t, tt.errType, err)
				}
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, foundUser)
				assert.Equal(t, createdUser.ID, foundUser.ID)
				assert.Equal(t, createdUser.Email, foundUser.Email)
				assert.Equal(t, createdUser.Name, foundUser.Name)
			}
		})
	}
}

func TestIntegrationPostgresRepository_FindByEmail(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	db := testutil.SetupTestDB(t)
	defer testutil.TruncateTable(t, db, "users")

	repo := NewPostgresRepository(db)
	ctx := context.Background()

	// arrange
	createdUser := testutil.CreateTestUser(t, db, "findme@example.com", "Find Me")

	tests := []struct {
		name    string
		email   string
		wantErr bool
		errType error
	}{
		{
			name:    "find existing user by email",
			email:   "findme@example.com",
			wantErr: false,
		},
		{
			name:    "email not found",
			email:   "notfound@example.com",
			wantErr: true,
			errType: sql.ErrNoRows,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// act
			foundUser, err := repo.FindByEmail(ctx, tt.email)

			// assert
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, foundUser)
				if tt.errType != nil {
					assert.Equal(t, tt.errType, err)
				}
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, foundUser)
				assert.Equal(t, createdUser.ID, foundUser.ID)
				assert.Equal(t, createdUser.Email, foundUser.Email)
				assert.Equal(t, createdUser.Name, foundUser.Name)
			}
		})
	}
}

func TestIntegrationPostgresRepository_List(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	db := testutil.SetupTestDB(t)
	defer testutil.TruncateTable(t, db, "users")

	repo := NewPostgresRepository(db)
	ctx := context.Background()

	// arrange - create multiple users
	testutil.CreateTestUser(t, db, "user1@test.com", "User 1")
	testutil.CreateTestUser(t, db, "user2@test.com", "User 2")
	testutil.CreateTestUser(t, db, "user3@test.com", "User 3")
	testutil.CreateTestUser(t, db, "user4@test.com", "User 4")
	testutil.CreateTestUser(t, db, "user5@test.com", "User 5")

	tests := []struct {
		name          string
		limit         int
		offset        int
		expectedCount int
		wantErr       bool
	}{
		{
			name:          "list all users",
			limit:         10,
			offset:        0,
			expectedCount: 5,
			wantErr:       false,
		},
		{
			name:          "list with limit",
			limit:         3,
			offset:        0,
			expectedCount: 3,
			wantErr:       false,
		},
		{
			name:          "list with offset",
			limit:         10,
			offset:        2,
			expectedCount: 3,
			wantErr:       false,
		},
		{
			name:          "list with limit and offset",
			limit:         2,
			offset:        1,
			expectedCount: 2,
			wantErr:       false,
		},
		{
			name:          "empty result with high offset",
			limit:         10,
			offset:        100,
			expectedCount: 0,
			wantErr:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// act
			users, err := repo.List(ctx, tt.limit, tt.offset)

			// assert
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				if tt.expectedCount == 0 {
					// Empty result can be nil or empty slice - both are valid
					assert.Len(t, users, 0)
				} else {
					assert.NotNil(t, users)
					assert.Len(t, users, tt.expectedCount)

					// verify users are ordered by created_at DESC
					if len(users) > 1 {
						for i := 0; i < len(users)-1; i++ {
							assert.True(t, users[i].CreatedAt.After(users[i+1].CreatedAt) || users[i].CreatedAt.Equal(users[i+1].CreatedAt),
								"Users should be ordered by created_at DESC")
						}
					}
				}
			}
		})
	}
}

func TestIntegrationPostgresRepository_List_EmptyTable(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	db := testutil.SetupTestDB(t)
	defer testutil.TruncateTable(t, db, "users")

	repo := NewPostgresRepository(db)
	ctx := context.Background()

	// act - list from empty table
	users, err := repo.List(ctx, 10, 0)

	// assert
	assert.NoError(t, err)
	assert.Len(t, users, 0) // Can be nil or empty slice - both valid
}
