package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Connect establishes a connection pool to PostgreSQL.
// It retries up to 3 times with exponential backoff on failure.
func Connect(databaseURL string) (*pgxpool.Pool, error) {
	var pool *pgxpool.Pool
	var err error

	for attempt := 1; attempt <= 3; attempt++ {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		config, parseErr := pgxpool.ParseConfig(databaseURL)
		if parseErr != nil {
			cancel()
			return nil, fmt.Errorf("unable to parse database URL: %w", parseErr)
		}

		config.MaxConns = 20
		config.MinConns = 2
		config.MaxConnLifetime = 1 * time.Hour
		config.MaxConnIdleTime = 30 * time.Minute

		pool, err = pgxpool.NewWithConfig(ctx, config)
		if err != nil {
			cancel()
			log.Printf("Database connection attempt %d failed: %v", attempt, err)
			time.Sleep(time.Duration(attempt) * 2 * time.Second)
			continue
		}

		// Verify connection
		if pingErr := pool.Ping(ctx); pingErr != nil {
			cancel()
			pool.Close()
			log.Printf("Database ping attempt %d failed: %v", attempt, pingErr)
			time.Sleep(time.Duration(attempt) * 2 * time.Second)
			continue
		}

		cancel()
		log.Println("✅ Database connected successfully")
		return pool, nil
	}

	return nil, fmt.Errorf("failed to connect to database after 3 attempts: %w", err)
}

// HealthCheck verifies the database connection is alive.
func HealthCheck(pool *pgxpool.Pool) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return pool.Ping(ctx)
}
