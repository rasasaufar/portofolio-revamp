package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/portfolio/api/internal/config"
	"github.com/portfolio/api/internal/database"
	"github.com/portfolio/api/internal/router"
	"github.com/portfolio/api/internal/service"
	"github.com/portfolio/api/seeds"
)

func main() {
	// Load .env file (ignore error if not present, e.g. in production)
	_ = godotenv.Load()

	// Load and validate configuration
	cfg := loadConfig()

	// Connect to database
	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Run migrations if --migrate flag is passed
	if hasFlag("-migrate") || hasFlag("--migrate") {
		migrationsDir := getMigrationsDir()
		if err := database.Migrate(db, migrationsDir); err != nil {
			log.Fatalf("❌ Migration failed: %v", err)
		}
		if !hasFlag("-serve") {
			log.Println("✅ Migrations complete. Exiting.")
			return
		}
	}

	// Run seed if --seed flag is passed
	if hasFlag("-seed") || hasFlag("--seed") {
		if err := seeds.Run(db, cfg.AdminDefaultPassword); err != nil {
			log.Fatalf("❌ Seed failed: %v", err)
		}
		if !hasFlag("-serve") {
			log.Println("✅ Seed complete. Exiting.")
			return
		}
	}

	// Initialize services
	authService := service.NewAuthService(cfg.JWTSecret, cfg.JWTExpiry)

	// Set up router
	r := router.New(db, authService, cfg.CORSOrigin)

	// Start server
	log.Printf("🚀 Portfolio API server starting on port %s", cfg.ServerPort)
	log.Printf("   Health check: http://localhost:%s/api/health", cfg.ServerPort)
	log.Printf("   CORS origin:  %s", cfg.CORSOrigin)

	if err := http.ListenAndServe(":"+cfg.ServerPort, r); err != nil {
		log.Fatalf("❌ Server failed to start: %v", err)
	}
}

// loadConfig loads configuration with panic recovery for better error messages.
func loadConfig() *config.Config {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("❌ Configuration error: %v", r)
		}
	}()
	return config.Load()
}

// hasFlag checks if a CLI flag is present in os.Args.
func hasFlag(flag string) bool {
	for _, arg := range os.Args[1:] {
		if arg == flag {
			return true
		}
	}
	return false
}

// getMigrationsDir returns the path to the migrations directory.
func getMigrationsDir() string {
	// Check if running from api/ directory or from cmd/server/
	if _, err := os.Stat("migrations"); err == nil {
		return "migrations"
	}
	if _, err := os.Stat("../../migrations"); err == nil {
		return "../../migrations"
	}
	return "migrations"
}
