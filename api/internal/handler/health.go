package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/portfolio/api/internal/database"
)

// HealthHandler handles the health check endpoint.
type HealthHandler struct {
	db *pgxpool.Pool
}

// NewHealthHandler creates a new HealthHandler.
func NewHealthHandler(db *pgxpool.Pool) *HealthHandler {
	return &HealthHandler{db: db}
}

// healthResponse represents the health check JSON response.
type healthResponse struct {
	Status   string `json:"status"`
	Message  string `json:"message"`
	Database string `json:"database"`
}

// Check handles GET /api/health
func (h *HealthHandler) Check(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dbStatus := "connected"
	if err := database.HealthCheck(h.db); err != nil {
		dbStatus = "disconnected"
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(healthResponse{
			Status:   "error",
			Message:  "Database connection failed",
			Database: dbStatus,
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(healthResponse{
		Status:   "ok",
		Message:  "Portfolio API is running",
		Database: dbStatus,
	})
}
