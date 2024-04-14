package http

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"go-ddd/pkg/config"
	"go-ddd/pkg/utils/redis"
	"go-ddd/pkg/utils/validation"
)

func TestServer(t *testing.T) {
	cfg, _ := config.LoadConfig()

	validator := validation.New()
	db, _ := pgxpool.New(context.Background(), cfg.DataBaseURL)
	cache := redis.New(redis.Config{
		Address:  cfg.RedisURI,
		Password: cfg.RedisPassword,
		Database: cfg.RedisDB,
	})

	server, err := Init(validator, db, cache)
	if err != nil {
		t.Fatalf("Failed to initialize server: %v", err)
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			t.Errorf("Could not listen on: %v, %s", cfg.HTTPPort, err)
			return
		}
	}()

	// Wait for the server to start.
	time.Sleep(time.Second)

	// Make a request to the /health endpoint.
	resp, err := http.Get(fmt.Sprintf("http://localhost:%d/health", cfg.HTTPPort))
	if err != nil {
		t.Fatalf("Failed to make GET request: %v", err)
	}
	defer resp.Body.Close()

	// Check the status code.
	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
