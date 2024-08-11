package main

import (
	"context"
	"testing"

	"github.com/Isshinfunada/TodoList/server/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockPgxPool is a mock of pgxpool.Pool
type MockPgxPool struct {
	mock.Mock
	pgxpool.Pool // pgxpool.Poolを埋め込む
}

func (m *MockPgxPool) Close(ctx context.Context) error {
	m.Called()
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

	dbpool, err := initDB(cfg)
	assert.NoError(t, err)
	assert.NotNil(t, dbpool)
}

// func TestStartServer(t *testing.T) {
// 	mockPool := new(MockPgxPool)
// 	startServer(mockPool)
// 	mockPool.AssertExpectations(t)
// }
