package user

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	tests := []struct {
		name          string
		email         string
		userName      string
		setupMock     func(*MockRepository)
		expectedError error
		expectUser    bool
	}{
		{
			name:     "success - valid user creation",
			email:    "test@example.com",
			userName: "John Doe",
			setupMock: func(m *MockRepository) {
				m.On("Create", mock.Anything, mock.MatchedBy(func(u *User) bool {
					return u.Email == "test@example.com" && u.Name == "John Doe"
				})).Return(nil).Once()
			},
			expectedError: nil,
			expectUser:    true,
		},
		{
			name:          "error - invalid email format",
			email:         "invalid-email",
			userName:      "John Doe",
			setupMock:     func(m *MockRepository) {}, // repo should NOT be called, validation fails first
			expectedError: ErrInvalidEmail,
			expectUser:    false,
		},
		{
			name:          "error - name too short",
			email:         "test@example.com",
			userName:      "J",
			setupMock:     func(m *MockRepository) {}, // repo should NOT be called, validation fails first
			expectedError: ErrInvalidName,
			expectUser:    false,
		},
		{
			name:          "error - name too long",
			email:         "test@example.com",
			userName:      "This is a very long name that exceeds the maximum allowed length of one hundred characters for testing",
			setupMock:     func(m *MockRepository) {}, // repo should NOT be called, validation fails first
			expectedError: ErrInvalidName,
			expectUser:    false,
		},
		{
			name:     "error - repository failure",
			email:    "test@example.com",
			userName: "John Doe",
			setupMock: func(m *MockRepository) {
				m.On("Create", mock.Anything, mock.Anything).Return(errors.New("database error")).Once()
			},
			expectedError: errors.New("Failed to create user"),
			expectUser:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			mockRepo := new(MockRepository)
			tt.setupMock(mockRepo)
			service := NewService(mockRepo)
			ctx := context.Background()

			// act
			user, err := service.Create(ctx, tt.email, tt.userName)

			// assert
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError.Error())
				assert.Nil(t, user)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, user)
				assert.Equal(t, tt.email, user.Email)
				assert.Equal(t, tt.userName, user.Name)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestFindByID(t *testing.T) {
	tests := []struct {
		name          string
		userID        int64
		setupMock     func(*MockRepository)
		expectedError error
		expectUser    bool
	}{
		{
			name:   "success - user found",
			userID: 1,
			setupMock: func(m *MockRepository) {
				m.On("FindByID", mock.Anything, int64(1)).Return(&User{
					ID:        1,
					Email:     "test@example.com",
					Name:      "John Doe",
					CreatedAt: time.Now(),
				}, nil).Once()
			},
			expectedError: nil,
			expectUser:    true,
		},
		{
			name:   "error - user not found",
			userID: 999,
			setupMock: func(m *MockRepository) {
				m.On("FindByID", mock.Anything, int64(999)).Return(nil, sql.ErrNoRows).Once()
			},
			expectedError: errors.New("not found"),
			expectUser:    false,
		},
		{
			name:   "error - repository failure",
			userID: 1,
			setupMock: func(m *MockRepository) {
				m.On("FindByID", mock.Anything, int64(1)).Return(nil, errors.New("database error")).Once()
			},
			expectedError: errors.New("Failed to fetch user"),
			expectUser:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			mockRepo := new(MockRepository)
			tt.setupMock(mockRepo)
			service := NewService(mockRepo)
			ctx := context.Background()

			// act
			user, err := service.FindByID(ctx, tt.userID)

			// assert
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError.Error())
				assert.Nil(t, user)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, user)
				assert.Equal(t, tt.userID, user.ID)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestListUsers(t *testing.T) {
	tests := []struct {
		name          string
		limit         int
		offset        int
		setupMock     func(*MockRepository)
		expectedError error
		expectedCount int
	}{
		{
			name:   "success - multiple users",
			limit:  10,
			offset: 0,
			setupMock: func(m *MockRepository) {
				m.On("List", mock.Anything, 10, 0).Return([]*User{
					{ID: 1, Email: "user1@example.com", Name: "User 1"},
					{ID: 2, Email: "user2@example.com", Name: "User 2"},
					{ID: 3, Email: "user3@example.com", Name: "User 3"},
				}, nil).Once()
			},
			expectedError: nil,
			expectedCount: 3,
		},
		{
			name:   "success - empty list",
			limit:  10,
			offset: 0,
			setupMock: func(m *MockRepository) {
				m.On("List", mock.Anything, 10, 0).Return([]*User{}, nil).Once()
			},
			expectedError: nil,
			expectedCount: 0,
		},
		{
			name:   "error - repository failure",
			limit:  10,
			offset: 0,
			setupMock: func(m *MockRepository) {
				m.On("List", mock.Anything, 10, 0).Return(nil, errors.New("database error")).Once()
			},
			expectedError: errors.New("Failed to fetch users"),
			expectedCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			mockRepo := new(MockRepository)
			tt.setupMock(mockRepo)
			service := NewService(mockRepo)
			ctx := context.Background()

			// act
			users, err := service.ListUsers(ctx, tt.limit, tt.offset)

			// assert
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError.Error())
				assert.Nil(t, users)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, users)
				assert.Len(t, users, tt.expectedCount)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestRegisterUser(t *testing.T) {
	tests := []struct {
		name          string
		email         string
		userName      string
		setupMock     func(*MockRepository)
		expectedError error
		expectUser    bool
	}{
		{
			name:     "success - new user registered",
			email:    "newuser@example.com",
			userName: "New User",
			setupMock: func(m *MockRepository) {
				// check if user exists (returns not found)
				m.On("FindByEmail", mock.Anything, "newuser@example.com").Return(nil, sql.ErrNoRows).Once()
				// create the user
				m.On("Create", mock.Anything, mock.MatchedBy(func(u *User) bool {
					return u.Email == "newuser@example.com" && u.Name == "New User"
				})).Return(nil).Once()
			},
			expectedError: nil,
			expectUser:    true,
		},
		{
			name:     "error - duplicate email",
			email:    "existing@example.com",
			userName: "Existing User",
			setupMock: func(m *MockRepository) {
				// user already exists
				m.On("FindByEmail", mock.Anything, "existing@example.com").Return(&User{
					ID:    1,
					Email: "existing@example.com",
					Name:  "Existing User",
				}, nil).Once()
			},
			expectedError: ErrDuplicateEmail,
			expectUser:    false,
		},
		{
			name:     "error - checking existing user fails",
			email:    "test@example.com",
			userName: "Test User",
			setupMock: func(m *MockRepository) {
				m.On("FindByEmail", mock.Anything, "test@example.com").Return(nil, errors.New("database error")).Once()
			},
			expectedError: errors.New("Failed to check existing user"),
			expectUser:    false,
		},
		{
			name:     "error - create user fails",
			email:    "newuser@example.com",
			userName: "New User",
			setupMock: func(m *MockRepository) {
				m.On("FindByEmail", mock.Anything, "newuser@example.com").Return(nil, sql.ErrNoRows).Once()
				m.On("Create", mock.Anything, mock.Anything).Return(errors.New("database error")).Once()
			},
			expectedError: errors.New("Failed to create user"),
			expectUser:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			mockRepo := new(MockRepository)
			tt.setupMock(mockRepo)
			service := NewService(mockRepo)
			ctx := context.Background()

			// act
			user, err := service.RegisterUser(ctx, tt.email, tt.userName)

			// assert
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError.Error())
				assert.Nil(t, user)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, user)
				assert.Equal(t, tt.email, user.Email)
				assert.Equal(t, tt.userName, user.Name)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
