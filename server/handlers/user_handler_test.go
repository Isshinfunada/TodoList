package handlers

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"firebase.google.com/go/auth"
	"github.com/Isshinfunada/TodoList/server/models"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockQueries is a mock for the database queries
type MockQueries struct {
	mock.Mock
}

func (m *MockQueries) CreateUser(ctx context.Context, arg models.CreateUserParams) error {
	args := m.Called(ctx, arg)
	return args.Error(0)
}

// Implement the interface methods for MockQueries
func (m *MockQueries) GetUserByFirebaseUID(ctx context.Context, firebaseUid string) (models.GetUserByFirebaseUIDRow, error) {
	args := m.Called(ctx, firebaseUid)
	return args.Get(0).(models.GetUserByFirebaseUIDRow), args.Error(1)
}

func TestRegisterUser(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(`{"Username": "testuser", "Email": "test@example.com", "Password": "password"}`))
	req.Header.Set(echo.HeaderAuthorization, "Bearer valid_token")
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockAuthClient := new(MockAuthClient)
	mockAuthClient.On("VerifyIDToken", mock.Anything, "valid_token").Return(&auth.Token{
		UID: "test_uid",
	}, nil)

	mockQueries := new(MockQueries)
	mockQueries.On("CreateUser", mock.Anything, models.CreateUserParams{
		Username:    "testuser",
		Email:       "test@example.com",
		Password:    "password",
		FirebaseUid: "test_uid",
	}).Return(nil)

	// Test
	h := RegisterUser(mockQueries, mockAuthClient)
	if assert.NoError(t, h(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Contains(t, rec.Body.String(), "testuser")
		assert.Contains(t, rec.Body.String(), "test@example.com")
	}

	// Test invalid token
	req = httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(`{"Username": "testuser", "Email": "test@example.com", "Password": "password"}`))
	req.Header.Set(echo.HeaderAuthorization, "Bearer invalid_token")
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	mockAuthClient.On("VerifyIDToken", mock.Anything, "invalid_token").Return((*auth.Token)(nil), errors.New("invalid token"))

	h = RegisterUser(mockQueries, mockAuthClient)
	if assert.NoError(t, h(c)) {
		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.Contains(t, rec.Body.String(), "無効なトークン")
	}
}
