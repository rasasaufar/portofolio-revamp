package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// PublicHandler handles all public portfolio API endpoints.
type PublicHandler struct {
	db *pgxpool.Pool
}

// NewPublicHandler creates a new PublicHandler.
func NewPublicHandler(db *pgxpool.Pool) *PublicHandler {
	return &PublicHandler{db: db}
}

// helper to write JSON response
func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// helper to write error JSON
func writeError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

// GetIdentity handles GET /api/identity
func (h *PublicHandler) GetIdentity(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	rows, err := h.db.Query(ctx,
		`SELECT id, name, role, headline, description, avatar_url, current_focus,
		availability_text, cta_primary_label, cta_primary_link, cta_secondary_label, cta_secondary_link
		FROM identity_console WHERE is_published = true ORDER BY order_number ASC LIMIT 1`)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to fetch identity data")
		return
	}
	defer rows.Close()

	var result []map[string]interface{}
	for rows.Next() {
		var id, name, role, headline, description string
		var avatarURL, availabilityText, ctaPrimaryLabel, ctaPrimaryLink, ctaSecondaryLabel, ctaSecondaryLink *string
		var currentFocus json.RawMessage

		err := rows.Scan(&id, &name, &role, &headline, &description, &avatarURL, &currentFocus,
			&availabilityText, &ctaPrimaryLabel, &ctaPrimaryLink, &ctaSecondaryLabel, &ctaSecondaryLink)
		if err != nil {
			writeError(w, http.StatusInternalServerError, "Failed to parse identity data")
			return
		}

		item := map[string]interface{}{
			"id":                  id,
			"name":                name,
			"role":                role,
			"headline":            headline,
			"description":         description,
			"avatar_url":          ptrStr(avatarURL),
			"current_focus":       currentFocus,
			"availability_text":   ptrStr(availabilityText),
			"cta_primary_label":   ptrStr(ctaPrimaryLabel),
			"cta_primary_link":    ptrStr(ctaPrimaryLink),
			"cta_secondary_label": ptrStr(ctaSecondaryLabel),
			"cta_secondary_link":  ptrStr(ctaSecondaryLink),
		}
		result = append(result, item)
	}

	if len(result) == 0 {
		writeJSON(w, http.StatusOK, nil)
		return
	}
	writeJSON(w, http.StatusOK, result[0])
}

// GetCapabilities handles GET /api/capabilities
func (h *PublicHandler) GetCapabilities(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	rows, err := h.db.Query(ctx,
		`SELECT id, label, value, description FROM capability_snapshots
		WHERE is_published = true ORDER BY order_number ASC`)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to fetch capabilities")
		return
	}
	defer rows.Close()

	var result []map[string]interface{}
	for rows.Next() {
		var id, label, value string
		var description *string
		if err := rows.Scan(&id, &label, &value, &description); err != nil {
			writeError(w, http.StatusInternalServerError, "Failed to parse capabilities")
			return
		}
		result = append(result, map[string]interface{}{
			"id": id, "label": label, "value": value, "description": ptrStr(description),
		})
	}

	writeJSON(w, http.StatusOK, ensureArray(result))
}

// GetStrengths handles GET /api/strengths
func (h *PublicHandler) GetStrengths(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	rows, err := h.db.Query(ctx,
		`SELECT id, title, description, bullet_points, icon_url FROM implementation_strengths
		WHERE is_published = true ORDER BY order_number ASC`)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to fetch strengths")
		return
	}
	defer rows.Close()

	var result []map[string]interface{}
	for rows.Next() {
		var id, title string
		var description, iconURL *string
		var bulletPoints json.RawMessage
		if err := rows.Scan(&id, &title, &description, &bulletPoints, &iconURL); err != nil {
			writeError(w, http.StatusInternalServerError, "Failed to parse strengths")
			return
		}
		result = append(result, map[string]interface{}{
			"id": id, "title": title, "description": ptrStr(description),
			"bullet_points": bulletPoints, "icon_url": ptrStr(iconURL),
		})
	}

	writeJSON(w, http.StatusOK, ensureArray(result))
}

// GetDossier handles GET /api/dossier
func (h *PublicHandler) GetDossier(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	rows, err := h.db.Query(ctx,
		`SELECT id, title, paragraph_1, paragraph_2, paragraph_3 FROM professional_dossier
		WHERE is_published = true LIMIT 1`)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to fetch dossier")
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id, title, p1 string
		var p2, p3 *string
		if err := rows.Scan(&id, &title, &p1, &p2, &p3); err != nil {
			writeError(w, http.StatusInternalServerError, "Failed to parse dossier")
			return
		}
		writeJSON(w, http.StatusOK, map[string]interface{}{
			"id": id, "title": title, "paragraph_1": p1, "paragraph_2": ptrStr(p2), "paragraph_3": ptrStr(p3),
		})
		return
	}

	writeJSON(w, http.StatusOK, nil)
}

// GetEducation handles GET /api/education
func (h *PublicHandler) GetEducation(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	rows, err := h.db.Query(ctx,
		`SELECT id, institution_name, degree, major, start_year, end_year, gpa, description, image_url, tags
		FROM education WHERE is_published = true ORDER BY order_number ASC`)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to fetch education")
		return
	}
	defer rows.Close()

	var result []map[string]interface{}
	for rows.Next() {
		var id, institutionName, degree string
		var major, startYear, endYear, gpa, description, imageURL *string
		var tags json.RawMessage
		if err := rows.Scan(&id, &institutionName, &degree, &major, &startYear, &endYear, &gpa, &description, &imageURL, &tags); err != nil {
			writeError(w, http.StatusInternalServerError, "Failed to parse education")
			return
		}
		result = append(result, map[string]interface{}{
			"id": id, "institution_name": institutionName, "degree": degree, "major": ptrStr(major),
			"start_year": ptrStr(startYear), "end_year": ptrStr(endYear), "gpa": ptrStr(gpa),
			"description": ptrStr(description), "image_url": ptrStr(imageURL), "tags": tags,
		})
	}

	writeJSON(w, http.StatusOK, ensureArray(result))
}

// GetExperiences handles GET /api/experiences
func (h *PublicHandler) GetExperiences(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	rows, err := h.db.Query(ctx,
		`SELECT id, company_name, position, start_date, end_date, is_current, description,
		bullet_points, tech_tags, logo_url, gallery_images
		FROM work_experiences WHERE is_published = true ORDER BY order_number ASC`)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to fetch experiences")
		return
	}
	defer rows.Close()

	var result []map[string]interface{}
	for rows.Next() {
		var id, companyName, position string
		var startDate, endDate, description, logoURL *string
		var isCurrent bool
		var bulletPoints, techTags, galleryImages json.RawMessage
		if err := rows.Scan(&id, &companyName, &position, &startDate, &endDate, &isCurrent, &description,
			&bulletPoints, &techTags, &logoURL, &galleryImages); err != nil {
			writeError(w, http.StatusInternalServerError, "Failed to parse experiences")
			return
		}
		result = append(result, map[string]interface{}{
			"id": id, "company_name": companyName, "position": position,
			"start_date": ptrStr(startDate), "end_date": ptrStr(endDate), "is_current": isCurrent,
			"description": ptrStr(description), "bullet_points": bulletPoints,
			"tech_tags": techTags, "logo_url": ptrStr(logoURL), "gallery_images": galleryImages,
		})
	}

	writeJSON(w, http.StatusOK, ensureArray(result))
}

// GetProjects handles GET /api/projects
func (h *PublicHandler) GetProjects(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	rows, err := h.db.Query(ctx,
		`SELECT id, title, category, description, tech_tags, image_url, demo_url, repo_url, is_featured
		FROM projects WHERE is_published = true ORDER BY order_number ASC`)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to fetch projects")
		return
	}
	defer rows.Close()

	var result []map[string]interface{}
	for rows.Next() {
		var id, title, category string
		var description, imageURL, demoURL, repoURL *string
		var isFeatured bool
		var techTags json.RawMessage
		if err := rows.Scan(&id, &title, &category, &description, &techTags, &imageURL, &demoURL, &repoURL, &isFeatured); err != nil {
			writeError(w, http.StatusInternalServerError, "Failed to parse projects")
			return
		}
		result = append(result, map[string]interface{}{
			"id": id, "title": title, "category": category, "description": ptrStr(description),
			"tech_tags": techTags, "image_url": ptrStr(imageURL), "demo_url": ptrStr(demoURL),
			"repo_url": ptrStr(repoURL), "is_featured": isFeatured,
		})
	}

	writeJSON(w, http.StatusOK, ensureArray(result))
}

// GetCertifications handles GET /api/certifications
func (h *PublicHandler) GetCertifications(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	rows, err := h.db.Query(ctx,
		`SELECT id, title, issuer, issued_date, expired_date, credential_id, credential_url,
		description, skills, image_url, category, status
		FROM certifications WHERE is_published = true ORDER BY order_number ASC`)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to fetch certifications")
		return
	}
	defer rows.Close()

	var result []map[string]interface{}
	for rows.Next() {
		var id, title, issuer, issuedDate, category, status string
		var expiredDate, credentialID, credentialURL, description, imageURL *string
		var skills json.RawMessage
		if err := rows.Scan(&id, &title, &issuer, &issuedDate, &expiredDate, &credentialID, &credentialURL,
			&description, &skills, &imageURL, &category, &status); err != nil {
			writeError(w, http.StatusInternalServerError, "Failed to parse certifications")
			return
		}
		result = append(result, map[string]interface{}{
			"id": id, "title": title, "issuer": issuer, "issued_date": issuedDate,
			"expired_date": ptrStr(expiredDate), "credential_id": ptrStr(credentialID),
			"credential_url": ptrStr(credentialURL), "description": ptrStr(description),
			"skills": skills, "image_url": ptrStr(imageURL), "category": category, "status": status,
		})
	}

	writeJSON(w, http.StatusOK, ensureArray(result))
}

// GetPublications handles GET /api/publications
func (h *PublicHandler) GetPublications(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	rows, err := h.db.Query(ctx,
		`SELECT id, title, journal_name, published_date, status, authors, description, tags, publication_url
		FROM publications WHERE is_published = true ORDER BY order_number ASC`)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to fetch publications")
		return
	}
	defer rows.Close()

	var result []map[string]interface{}
	for rows.Next() {
		var id, title, journalName, publishedDate, status, authors string
		var description, publicationURL *string
		var tags json.RawMessage
		if err := rows.Scan(&id, &title, &journalName, &publishedDate, &status, &authors, &description, &tags, &publicationURL); err != nil {
			writeError(w, http.StatusInternalServerError, "Failed to parse publications")
			return
		}
		result = append(result, map[string]interface{}{
			"id": id, "title": title, "journal_name": journalName, "published_date": publishedDate,
			"status": status, "authors": authors, "description": ptrStr(description),
			"tags": tags, "publication_url": ptrStr(publicationURL),
		})
	}

	writeJSON(w, http.StatusOK, ensureArray(result))
}

// GetContactInfo handles GET /api/contact-info
func (h *PublicHandler) GetContactInfo(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	rows, err := h.db.Query(ctx,
		`SELECT id, email, phone, whatsapp_url, github_url, linkedin_url, instagram_url, location, contact_description
		FROM contact_info WHERE is_published = true LIMIT 1`)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to fetch contact info")
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		var email, phone, whatsappURL, githubURL, linkedinURL, instagramURL, location, contactDescription *string
		if err := rows.Scan(&id, &email, &phone, &whatsappURL, &githubURL, &linkedinURL, &instagramURL, &location, &contactDescription); err != nil {
			writeError(w, http.StatusInternalServerError, "Failed to parse contact info")
			return
		}
		writeJSON(w, http.StatusOK, map[string]interface{}{
			"id": id, "email": ptrStr(email), "phone": ptrStr(phone),
			"whatsapp_url": ptrStr(whatsappURL), "github_url": ptrStr(githubURL),
			"linkedin_url": ptrStr(linkedinURL), "instagram_url": ptrStr(instagramURL),
			"location": ptrStr(location), "contact_description": ptrStr(contactDescription),
		})
		return
	}

	writeJSON(w, http.StatusOK, nil)
}

// GetSiteSettings handles GET /api/site-settings
func (h *PublicHandler) GetSiteSettings(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	var siteTitle string
	var metaDescription, faviconURL, logoURL, footerText *string
	var themeMode string
	var maintenanceMode bool

	err := h.db.QueryRow(ctx,
		`SELECT site_title, meta_description, favicon_url, logo_url, footer_text, theme_mode, maintenance_mode
		FROM site_settings LIMIT 1`).Scan(&siteTitle, &metaDescription, &faviconURL, &logoURL, &footerText, &themeMode, &maintenanceMode)
	if err != nil {
		writeJSON(w, http.StatusOK, map[string]interface{}{
			"site_title": "Portfolio", "theme_mode": "dark", "maintenance_mode": false,
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"site_title": siteTitle, "meta_description": ptrStr(metaDescription),
		"favicon_url": ptrStr(faviconURL), "logo_url": ptrStr(logoURL),
		"footer_text": ptrStr(footerText), "theme_mode": themeMode, "maintenance_mode": maintenanceMode,
	})
}

// CreateContactMessage handles POST /api/contact/messages
func (h *PublicHandler) CreateContactMessage(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name    string `json:"name"`
		Email   string `json:"email"`
		Message string `json:"message"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.Name == "" || req.Email == "" || req.Message == "" {
		writeError(w, http.StatusBadRequest, "Name, email, and message are required")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	_, err := h.db.Exec(ctx,
		`INSERT INTO contact_messages (name, email, message) VALUES ($1, $2, $3)`,
		req.Name, req.Email, req.Message)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to save message")
		return
	}

	writeJSON(w, http.StatusCreated, map[string]string{"message": "Message sent successfully"})
}

// --- Helpers ---

func ptrStr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func ensureArray(arr []map[string]interface{}) []map[string]interface{} {
	if arr == nil {
		return []map[string]interface{}{}
	}
	return arr
}
