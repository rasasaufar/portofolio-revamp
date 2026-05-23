package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

// AdminHandler handles all ***REMOVED*** CRUD endpoints using a generic approach.
type AdminHandler struct {
	db *pgxpool.Pool
}

// NewAdminHandler creates a new AdminHandler.
func NewAdminHandler(db *pgxpool.Pool) *AdminHandler {
	return &AdminHandler{db: db}
}

// ResourceConfig defines the table and columns for a resource.
type ResourceConfig struct {
	Table       string
	Columns     []string // columns for INSERT/UPDATE (excluding id, created_at, updated_at)
	HasOrder    bool
	HasPublish  bool
	ListOrderBy string
}

// resourceConfigs maps resource names to their configurations.
var resourceConfigs = map[string]ResourceConfig{
	"identity": {
		Table:       "identity_console",
		Columns:     []string{"name", "role", "headline", "description", "avatar_url", "current_focus", "availability_text", "cta_primary_label", "cta_primary_link", "cta_secondary_label", "cta_secondary_link", "order_number", "is_published"},
		HasOrder:    true,
		HasPublish:  true,
		ListOrderBy: "order_number ASC",
	},
	"capabilities": {
		Table:       "capability_snapshots",
		Columns:     []string{"label", "value", "description", "order_number", "is_published"},
		HasOrder:    true,
		HasPublish:  true,
		ListOrderBy: "order_number ASC",
	},
	"strengths": {
		Table:       "implementation_strengths",
		Columns:     []string{"title", "description", "bullet_points", "icon_url", "order_number", "is_published"},
		HasOrder:    true,
		HasPublish:  true,
		ListOrderBy: "order_number ASC",
	},
	"dossier": {
		Table:       "professional_dossier",
		Columns:     []string{"title", "paragraph_1", "paragraph_2", "paragraph_3", "is_published"},
		HasOrder:    false,
		HasPublish:  true,
		ListOrderBy: "created_at DESC",
	},
	"education": {
		Table:       "education",
		Columns:     []string{"institution_name", "degree", "major", "start_year", "end_year", "gpa", "description", "image_url", "tags", "order_number", "is_published"},
		HasOrder:    true,
		HasPublish:  true,
		ListOrderBy: "order_number ASC",
	},
	"experiences": {
		Table:       "work_experiences",
		Columns:     []string{"company_name", "position", "start_date", "end_date", "is_current", "description", "bullet_points", "tech_tags", "logo_url", "gallery_images", "order_number", "is_published"},
		HasOrder:    true,
		HasPublish:  true,
		ListOrderBy: "order_number ASC",
	},
	"projects": {
		Table:       "projects",
		Columns:     []string{"title", "category", "description", "tech_tags", "image_url", "demo_url", "repo_url", "is_featured", "order_number", "is_published"},
		HasOrder:    true,
		HasPublish:  true,
		ListOrderBy: "order_number ASC",
	},
	"certifications": {
		Table:       "certifications",
		Columns:     []string{"title", "issuer", "issued_date", "expired_date", "credential_id", "credential_url", "description", "skills", "image_url", "category", "status", "order_number", "is_published"},
		HasOrder:    true,
		HasPublish:  true,
		ListOrderBy: "order_number ASC",
	},
	"publications": {
		Table:       "publications",
		Columns:     []string{"title", "journal_name", "published_date", "status", "authors", "description", "tags", "publication_url", "order_number", "is_published"},
		HasOrder:    true,
		HasPublish:  true,
		ListOrderBy: "order_number ASC",
	},
	"contact": {
		Table:       "contact_info",
		Columns:     []string{"email", "phone", "whatsapp_url", "github_url", "linkedin_url", "instagram_url", "location", "contact_description", "is_published"},
		HasOrder:    false,
		HasPublish:  true,
		ListOrderBy: "created_at DESC",
	},
	"messages": {
		Table:       "contact_messages",
		Columns:     []string{"name", "email", "message", "is_read"},
		HasOrder:    false,
		HasPublish:  false,
		ListOrderBy: "created_at DESC",
	},
	"settings": {
		Table:       "site_settings",
		Columns:     []string{"site_title", "meta_description", "favicon_url", "logo_url", "footer_text", "theme_mode", "maintenance_mode"},
		HasOrder:    false,
		HasPublish:  false,
		ListOrderBy: "updated_at DESC",
	},
}

// DashboardStats handles GET /api/***REMOVED***/dashboard/stats
func (h *AdminHandler) DashboardStats(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	stats := map[string]int{}
	tables := map[string]string{
		"identity":       "identity_console",
		"capabilities":   "capability_snapshots",
		"strengths":      "implementation_strengths",
		"dossier":        "professional_dossier",
		"education":      "education",
		"experiences":    "work_experiences",
		"projects":       "projects",
		"certifications": "certifications",
		"publications":   "publications",
		"contact":        "contact_info",
		"messages":       "contact_messages",
	}

	for key, table := range tables {
		var count int
		err := h.db.QueryRow(ctx, fmt.Sprintf("SELECT COUNT(*) FROM %s", table)).Scan(&count)
		if err != nil {
			stats[key] = 0
		} else {
			stats[key] = count
		}
	}

	// Unread messages count
	var unread int
	h.db.QueryRow(ctx, "SELECT COUNT(*) FROM contact_messages WHERE is_read = false").Scan(&unread)
	stats["unread_messages"] = unread

	writeJSON(w, http.StatusOK, stats)
}

// List handles GET /api/***REMOVED***/{resource}
func (h *AdminHandler) List(w http.ResponseWriter, r *http.Request) {
	resource := chi.URLParam(r, "resource")
	cfg, ok := resourceConfigs[resource]
	if !ok {
		writeError(w, http.StatusNotFound, "Resource not found")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	query := fmt.Sprintf("SELECT * FROM %s ORDER BY %s", cfg.Table, cfg.ListOrderBy)
	rows, err := h.db.Query(ctx, query)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to fetch records")
		return
	}
	defer rows.Close()

	results, err := scanRowsToMaps(rows)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to parse records")
		return
	}

	if results == nil {
		results = []map[string]interface{}{}
	}

	writeJSON(w, http.StatusOK, results)
}

// GetByID handles GET /api/***REMOVED***/{resource}/{id}
func (h *AdminHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	resource := chi.URLParam(r, "resource")
	id := chi.URLParam(r, "id")

	cfg, ok := resourceConfigs[resource]
	if !ok {
		writeError(w, http.StatusNotFound, "Resource not found")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", cfg.Table)
	rows, err := h.db.Query(ctx, query, id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to fetch record")
		return
	}
	defer rows.Close()

	results, err := scanRowsToMaps(rows)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to parse record")
		return
	}

	if len(results) == 0 {
		writeError(w, http.StatusNotFound, "Record not found")
		return
	}

	writeJSON(w, http.StatusOK, results[0])
}

// Create handles POST /api/***REMOVED***/{resource}
func (h *AdminHandler) Create(w http.ResponseWriter, r *http.Request) {
	resource := chi.URLParam(r, "resource")
	cfg, ok := resourceConfigs[resource]
	if !ok {
		writeError(w, http.StatusNotFound, "Resource not found")
		return
	}

	var body map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	// Build INSERT query dynamically
	cols := ""
	placeholders := ""
	values := []interface{}{}
	idx := 1

	for _, col := range cfg.Columns {
		if val, exists := body[col]; exists {
			if cols != "" {
				cols += ", "
				placeholders += ", "
			}
			cols += col
			placeholders += fmt.Sprintf("$%d", idx)
			// Handle JSONB fields
			switch v := val.(type) {
			case []interface{}, map[string]interface{}:
				jsonBytes, _ := json.Marshal(v)
				values = append(values, string(jsonBytes))
			default:
				values = append(values, val)
			}
			idx++
		}
	}

	if cols == "" {
		writeError(w, http.StatusBadRequest, "No valid fields provided")
		return
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) RETURNING id", cfg.Table, cols, placeholders)
	var newID string
	err := h.db.QueryRow(ctx, query, values...).Scan(&newID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to create record: %v", err))
		return
	}

	writeJSON(w, http.StatusCreated, map[string]string{"id": newID, "message": "Record created successfully"})
}

// Update handles PUT /api/***REMOVED***/{resource}/{id}
func (h *AdminHandler) Update(w http.ResponseWriter, r *http.Request) {
	resource := chi.URLParam(r, "resource")
	id := chi.URLParam(r, "id")

	cfg, ok := resourceConfigs[resource]
	if !ok {
		writeError(w, http.StatusNotFound, "Resource not found")
		return
	}

	var body map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	// Build UPDATE query dynamically
	setClauses := ""
	values := []interface{}{}
	idx := 1

	for _, col := range cfg.Columns {
		if val, exists := body[col]; exists {
			if setClauses != "" {
				setClauses += ", "
			}
			setClauses += fmt.Sprintf("%s = $%d", col, idx)
			switch v := val.(type) {
			case []interface{}, map[string]interface{}:
				jsonBytes, _ := json.Marshal(v)
				values = append(values, string(jsonBytes))
			default:
				values = append(values, val)
			}
			idx++
		}
	}

	if setClauses == "" {
		writeError(w, http.StatusBadRequest, "No valid fields provided")
		return
	}

	// Add updated_at if table has it
	if resource != "messages" {
		setClauses += fmt.Sprintf(", updated_at = $%d", idx)
		values = append(values, time.Now())
		idx++
	}

	values = append(values, id)
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", cfg.Table, setClauses, idx)

	result, err := h.db.Exec(ctx, query, values...)
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to update record: %v", err))
		return
	}

	if result.RowsAffected() == 0 {
		writeError(w, http.StatusNotFound, "Record not found")
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"message": "Record updated successfully"})
}

// Delete handles DELETE /api/***REMOVED***/{resource}/{id}
func (h *AdminHandler) Delete(w http.ResponseWriter, r *http.Request) {
	resource := chi.URLParam(r, "resource")
	id := chi.URLParam(r, "id")

	cfg, ok := resourceConfigs[resource]
	if !ok {
		writeError(w, http.StatusNotFound, "Resource not found")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", cfg.Table)
	result, err := h.db.Exec(ctx, query, id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to delete record")
		return
	}

	if result.RowsAffected() == 0 {
		writeError(w, http.StatusNotFound, "Record not found")
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"message": "Record deleted successfully"})
}

// Publish handles PATCH /api/***REMOVED***/{resource}/{id}/publish
func (h *AdminHandler) Publish(w http.ResponseWriter, r *http.Request) {
	h.setPublishState(w, r, true)
}

// Unpublish handles PATCH /api/***REMOVED***/{resource}/{id}/unpublish
func (h *AdminHandler) Unpublish(w http.ResponseWriter, r *http.Request) {
	h.setPublishState(w, r, false)
}

func (h *AdminHandler) setPublishState(w http.ResponseWriter, r *http.Request, published bool) {
	resource := chi.URLParam(r, "resource")
	id := chi.URLParam(r, "id")

	cfg, ok := resourceConfigs[resource]
	if !ok || !cfg.HasPublish {
		writeError(w, http.StatusNotFound, "Resource not found or does not support publish")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	query := fmt.Sprintf("UPDATE %s SET is_published = $1, updated_at = $2 WHERE id = $3", cfg.Table)
	result, err := h.db.Exec(ctx, query, published, time.Now(), id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to update publish state")
		return
	}

	if result.RowsAffected() == 0 {
		writeError(w, http.StatusNotFound, "Record not found")
		return
	}

	state := "published"
	if !published {
		state = "unpublished"
	}
	writeJSON(w, http.StatusOK, map[string]string{"message": fmt.Sprintf("Record %s successfully", state)})
}

// Reorder handles PUT /api/***REMOVED***/{resource}/reorder
func (h *AdminHandler) Reorder(w http.ResponseWriter, r *http.Request) {
	resource := chi.URLParam(r, "resource")

	cfg, ok := resourceConfigs[resource]
	if !ok || !cfg.HasOrder {
		writeError(w, http.StatusNotFound, "Resource not found or does not support reorder")
		return
	}

	var body struct {
		IDs []string `json:"ids"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if len(body.IDs) == 0 {
		writeError(w, http.StatusBadRequest, "IDs array is required")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	for i, id := range body.IDs {
		query := fmt.Sprintf("UPDATE %s SET order_number = $1, updated_at = $2 WHERE id = $3", cfg.Table)
		_, err := h.db.Exec(ctx, query, i, time.Now(), id)
		if err != nil {
			writeError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to reorder: %v", err))
			return
		}
	}

	writeJSON(w, http.StatusOK, map[string]string{"message": "Reorder successful"})
}

// MarkRead handles PATCH /api/***REMOVED***/messages/{id}/read
func (h *AdminHandler) MarkRead(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	_, err := h.db.Exec(ctx, "UPDATE contact_messages SET is_read = true WHERE id = $1", id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to mark as read")
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"message": "Message marked as read"})
}

// convertPgValue converts pgx-specific types to JSON-friendly Go types.
func convertPgValue(val interface{}) interface{} {
	if val == nil {
		return nil
	}
	switch v := val.(type) {
	case [16]uint8:
		// UUID bytes → string format
		u := pgtype.UUID{Bytes: v, Valid: true}
		str, _ := u.Value()
		return str
	case pgtype.UUID:
		if !v.Valid {
			return nil
		}
		str, _ := v.Value()
		return str
	case time.Time:
		return v.Format(time.RFC3339)
	case pgtype.Timestamptz:
		if !v.Valid {
			return nil
		}
		return v.Time.Format(time.RFC3339)
	case pgtype.Text:
		if !v.Valid {
			return nil
		}
		return v.String
	case pgtype.Bool:
		if !v.Valid {
			return nil
		}
		return v.Bool
	case pgtype.Int4:
		if !v.Valid {
			return nil
		}
		return v.Int32
	default:
		return val
	}
}

// scanRowsToMaps converts pgx rows to a slice of maps with proper type conversion.
func scanRowsToMaps(rows interface {
	FieldDescriptions() []pgconn.FieldDescription
	Next() bool
	Values() ([]interface{}, error)
}) ([]map[string]interface{}, error) {
	fieldDescs := rows.FieldDescriptions()
	var results []map[string]interface{}

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return nil, err
		}

		item := make(map[string]interface{})
		for i, fd := range fieldDescs {
			item[string(fd.Name)] = convertPgValue(values[i])
		}
		results = append(results, item)
	}

	return results, nil
}
