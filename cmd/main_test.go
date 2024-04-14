package main

import (
    "testing"
    "github.com/jackc/pgx/v5/pgxpool"
    "github.com/stretchr/testify/mock"
)

type MockDatabase struct {
    mock.Mock
}

func (m *MockDatabase) New(dsn string) (*pgxpool.Pool, error) {
    args := m.Called(dsn)
    return args.Get(0).(*pgxpool.Pool), args.Error(1)
}

func TestInitDBConnection(t *testing.T) {
    mockDatabase := new(MockDatabase)
    mockDatabase.On("New", "test_dsn").Return(&pgxpool.Pool{}, nil)

    db, err := initDBConnection("test_dsn")
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }

    if db == nil {
        t.Errorf("Expected db to be not nil")
    }

    mockDatabase.AssertExpectations(t)
}