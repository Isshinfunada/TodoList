package main

import (
	"database/sql"
	"testing"

	firebase "firebase.google.com/go"
	"github.com/Isshinfunada/TodoList/server/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockDB is a mock of sql.DB
type MockDB struct {
	mock.Mock
	DB *sql.DB
}

func (m *MockDB) Close() error {
	m.Called()
	return nil
}

func (m *MockDB) Ping() error {
	return nil
}

func TestInitDB(t *testing.T) {
	cfg := &config.Config{
		DBUser:     "testuser",
		DBPassword: "testpassword",
		DBName:     "testdb",
		DBHost:     "localhost",
		DBPort:     "5432",
	}

	db, err := initDB(cfg)
	assert.NoError(t, err)
	assert.NotNil(t, db)
}

func TestStartServer(t *testing.T) {
	mockDB := &MockDB{DB: &sql.DB{}}
	mockDB.On("Close").Return(nil)

	// Firebaseのモックを作成
	mockFirebaseApp := &firebase.App{}

	startServer(mockDB.DB, mockFirebaseApp)
	mockDB.AssertExpectations(t)
}
