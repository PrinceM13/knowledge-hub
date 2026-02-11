package userhttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PrinceM13/knowledge-hub-api/internal/app"
	userdb "github.com/PrinceM13/knowledge-hub-api/internal/db/user"
	"github.com/PrinceM13/knowledge-hub-api/internal/middleware"
	"github.com/PrinceM13/knowledge-hub-api/internal/testutil"
	"github.com/PrinceM13/knowledge-hub-api/internal/user"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTestRouter(t *testing.T) (*gin.Engine, *app.App) {
	t.Helper()

	// Setup test database
	db := testutil.SetupTestDB(t)
	testutil.TruncateTable(t, db, "users")

	// Setup repositories and services
	userRepo := userdb.NewPostgresRepository(db)
	userSvc := user.NewService(userRepo)
	application := app.New(userSvc)

	// Setup Gin router
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(middleware.ErrorHandler())

	// Register routes
	v1 := router.Group("/api/v1")
	Register(v1, application)

	return router, application
}

func TestIntegrationUserHandler_RegisterUser(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	router, _ := setupTestRouter(t)

	tests := []struct {
		name           string
		requestBody    interface{}
		expectedStatus int
		checkResponse  func(t *testing.T, body map[string]interface{})
	}{
		{
			name: "success - register new user",
			requestBody: map[string]string{
				"email": "newuser@example.com",
				"name":  "New User",
			},
			expectedStatus: http.StatusCreated,
			checkResponse: func(t *testing.T, body map[string]interface{}) {
				assert.NotZero(t, body["id"])
				assert.Equal(t, "newuser@example.com", body["email"])
				assert.Equal(t, "New User", body["name"])
				assert.NotEmpty(t, body["createdAt"])
			},
		},
		{
			name: "error - invalid email format",
			requestBody: map[string]string{
				"email": "invalid-email",
				"name":  "Test User",
			},
			expectedStatus: http.StatusBadRequest,
			checkResponse: func(t *testing.T, body map[string]interface{}) {
				assert.Contains(t, body, "error")
				assert.Contains(t, body, "message")
			},
		},
		{
			name: "error - missing email",
			requestBody: map[string]string{
				"name": "Test User",
			},
			expectedStatus: http.StatusBadRequest,
			checkResponse: func(t *testing.T, body map[string]interface{}) {
				assert.Contains(t, body, "error")
			},
		},
		{
			name: "error - missing name",
			requestBody: map[string]string{
				"email": "test@example.com",
			},
			expectedStatus: http.StatusBadRequest,
			checkResponse: func(t *testing.T, body map[string]interface{}) {
				assert.Contains(t, body, "error")
			},
		},
		{
			name: "error - name too short",
			requestBody: map[string]string{
				"email": "test@example.com",
				"name":  "A",
			},
			expectedStatus: http.StatusBadRequest,
			checkResponse: func(t *testing.T, body map[string]interface{}) {
				assert.Contains(t, body, "error")
			},
		},
		{
			name:           "error - empty request body",
			requestBody:    map[string]string{},
			expectedStatus: http.StatusBadRequest,
			checkResponse: func(t *testing.T, body map[string]interface{}) {
				assert.Contains(t, body, "error")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			jsonBody, err := json.Marshal(tt.requestBody)
			require.NoError(t, err)

			req := httptest.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			// act
			router.ServeHTTP(w, req)

			// assert
			assert.Equal(t, tt.expectedStatus, w.Code)

			var response map[string]interface{}
			err = json.Unmarshal(w.Body.Bytes(), &response)
			require.NoError(t, err)

			if tt.checkResponse != nil {
				tt.checkResponse(t, response)
			}
		})
	}
}

func TestIntegrationUserHandler_RegisterUser_DuplicateEmail(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	router, _ := setupTestRouter(t)

	// arrange - create first user
	firstUser := map[string]string{
		"email": "duplicate@example.com",
		"name":  "First User",
	}
	jsonBody, _ := json.Marshal(firstUser)
	req1 := httptest.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(jsonBody))
	req1.Header.Set("Content-Type", "application/json")
	w1 := httptest.NewRecorder()
	router.ServeHTTP(w1, req1)
	require.Equal(t, http.StatusCreated, w1.Code)

	// act - try to create second user with same email
	secondUser := map[string]string{
		"email": "duplicate@example.com",
		"name":  "Second User",
	}
	jsonBody, _ = json.Marshal(secondUser)
	req2 := httptest.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(jsonBody))
	req2.Header.Set("Content-Type", "application/json")
	w2 := httptest.NewRecorder()
	router.ServeHTTP(w2, req2)

	// assert
	assert.Equal(t, http.StatusConflict, w2.Code)

	var response map[string]interface{}
	json.Unmarshal(w2.Body.Bytes(), &response)
	assert.Contains(t, response, "error")
	assert.Contains(t, response, "message")
}

func TestIntegrationUserHandler_GetUserByID(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	router, _ := setupTestRouter(t)
	db := testutil.SetupTestDB(t)

	// arrange - create test user
	testUser := testutil.CreateTestUser(t, db, "getme@example.com", "Get Me")

	tests := []struct {
		name           string
		userID         string
		expectedStatus int
		checkResponse  func(t *testing.T, body map[string]interface{})
	}{
		{
			name:           "success - get existing user",
			userID:         fmt.Sprintf("%d", testUser.ID),
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, body map[string]interface{}) {
				assert.Equal(t, float64(testUser.ID), body["id"])
				assert.Equal(t, testUser.Email, body["email"])
				assert.Equal(t, testUser.Name, body["name"])
				assert.NotEmpty(t, body["createdAt"])
			},
		},
		{
			name:           "error - user not found",
			userID:         "99999",
			expectedStatus: http.StatusNotFound,
			checkResponse: func(t *testing.T, body map[string]interface{}) {
				assert.Contains(t, body, "error")
				assert.Contains(t, body, "message")
			},
		},
		{
			name:           "error - invalid user ID format",
			userID:         "invalid",
			expectedStatus: http.StatusBadRequest,
			checkResponse: func(t *testing.T, body map[string]interface{}) {
				assert.Contains(t, body, "error")
				assert.Contains(t, body["message"], "id")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			url := fmt.Sprintf("/api/v1/users/%s", tt.userID)
			req := httptest.NewRequest(http.MethodGet, url, nil)
			w := httptest.NewRecorder()

			// act
			router.ServeHTTP(w, req)

			// assert
			assert.Equal(t, tt.expectedStatus, w.Code)

			var response map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			require.NoError(t, err)

			if tt.checkResponse != nil {
				tt.checkResponse(t, response)
			}
		})
	}
}

func TestIntegrationUserHandler_ListUsers(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	router, _ := setupTestRouter(t)
	db := testutil.SetupTestDB(t)

	// arrange - create test users
	testutil.CreateTestUser(t, db, "user1@test.com", "User 1")
	testutil.CreateTestUser(t, db, "user2@test.com", "User 2")
	testutil.CreateTestUser(t, db, "user3@test.com", "User 3")

	tests := []struct {
		name           string
		queryParams    string
		expectedStatus int
		checkResponse  func(t *testing.T, body []interface{})
	}{
		{
			name:           "success - list all users with default params",
			queryParams:    "",
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, body []interface{}) {
				assert.Len(t, body, 3)
			},
		},
		{
			name:           "success - list with limit",
			queryParams:    "?limit=2",
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, body []interface{}) {
				assert.Len(t, body, 2)
			},
		},
		{
			name:           "success - list with offset",
			queryParams:    "?offset=1",
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, body []interface{}) {
				assert.Len(t, body, 2)
			},
		},
		{
			name:           "success - list with limit and offset",
			queryParams:    "?limit=1&offset=1",
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, body []interface{}) {
				assert.Len(t, body, 1)
			},
		},
		{
			name:           "error - invalid limit",
			queryParams:    "?limit=invalid",
			expectedStatus: http.StatusBadRequest,
			checkResponse:  nil, // Error response is different structure
		},
		{
			name:           "error - invalid offset",
			queryParams:    "?offset=invalid",
			expectedStatus: http.StatusBadRequest,
			checkResponse:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			url := "/api/v1/users" + tt.queryParams
			req := httptest.NewRequest(http.MethodGet, url, nil)
			w := httptest.NewRecorder()

			// act
			router.ServeHTTP(w, req)

			// assert
			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedStatus == http.StatusOK && tt.checkResponse != nil {
				var response []interface{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				require.NoError(t, err)
				tt.checkResponse(t, response)
			} else if tt.expectedStatus != http.StatusOK {
				var response map[string]interface{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				require.NoError(t, err)
				assert.Contains(t, response, "error")
			}
		})
	}
}

func TestIntegrationUserHandler_ListUsers_EmptyTable(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	router, _ := setupTestRouter(t)

	// arrange - no users in database
	req := httptest.NewRequest(http.MethodGet, "/api/v1/users", nil)
	w := httptest.NewRecorder()

	// act
	router.ServeHTTP(w, req)

	// assert
	assert.Equal(t, http.StatusOK, w.Code)

	var response []interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)
	assert.Empty(t, response)
}
