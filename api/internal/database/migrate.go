package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Migrate runs all SQL migration files from the given directory in order.
// It tracks applied migrations in a schema_migrations table.
func Migrate(pool *pgxpool.Pool, migrationsDir string) error {
	ctx := context.Background()

	// Create migrations tracking table
	_, err := pool.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS schema_migrations (
			id SERIAL PRIMARY KEY,
			filename VARCHAR(255) NOT NULL UNIQUE,
			applied_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create schema_migrations table: %w", err)
	}

	// Read migration files
	entries, err := os.ReadDir(migrationsDir)
	if err != nil {
		return fmt.Errorf("failed to read migrations directory: %w", err)
	}

	// Filter and sort SQL files
	var files []string
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".sql") {
			files = append(files, entry.Name())
		}
	}
	sort.Strings(files)

	// Apply each migration
	for _, filename := range files {
		// Check if already applied
		var count int
		err := pool.QueryRow(ctx,
			"SELECT COUNT(*) FROM schema_migrations WHERE filename = $1",
			filename,
		).Scan(&count)
		if err != nil {
			return fmt.Errorf("failed to check migration status for %s: %w", filename, err)
		}
		if count > 0 {
			continue
		}

		// Read and execute migration
		filePath := filepath.Join(migrationsDir, filename)
		content, err := os.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("failed to read migration file %s: %w", filename, err)
		}

		_, err = pool.Exec(ctx, string(content))
		if err != nil {
			return fmt.Errorf("failed to execute migration %s: %w", filename, err)
		}

		// Record migration
		_, err = pool.Exec(ctx,
			"INSERT INTO schema_migrations (filename) VALUES ($1)",
			filename,
		)
		if err != nil {
			return fmt.Errorf("failed to record migration %s: %w", filename, err)
		}

		log.Printf("✅ Applied migration: %s", filename)
	}

	log.Println("✅ All migrations applied successfully")
	return nil
}
