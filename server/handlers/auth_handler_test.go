package handlers

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"firebase.google.com/go/auth"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockAuthClient is a mock for the Firebase Auth client
type MockAuthClient struct {
	mock.Mock
}

func (m *MockAuthClient) VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error) {
	args := m.Called(ctx, idToken)
	return args.Get(0).(*auth.Token), args.Error(1)
}

func TestLogin(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/login", nil)
	req.Header.Set(echo.HeaderAuthorization, "Bearer valid_token")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockAuthClient := new(MockAuthClient)
	mockAuthClient.On("VerifyIDToken", mock.Anything, "valid_token").Return(&auth.Token{
		UID: "test_uid",
		Claims: map[string]interface{}{
			"email": "test@example.com",
		},
	}, nil)

	// Test
	h := Login(mockAuthClient)
	if assert.NoError(t, h(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "test_uid")
		assert.Contains(t, rec.Body.String(), "test@example.com")
	}

	// Test invalid token
	req = httptest.NewRequest(http.MethodPost, "/login", nil)
	req.Header.Set(echo.HeaderAuthorization, "Bearer invalid_token")
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	mockAuthClient.On("VerifyIDToken", mock.Anything, "invalid_token").Return((*auth.Token)(nil), errors.New("invalid token"))

	h = Login(mockAuthClient)
	if assert.NoError(t, h(c)) {
		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.Contains(t, rec.Body.String(), "無効なトークン")
	}
}
