package handler

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// UploadHandler handles file upload endpoints.
type UploadHandler struct {
	uploadDir string
	baseURL   string
}

// NewUploadHandler creates a new UploadHandler.
// uploadDir is the local directory to store files.
// baseURL is the public base URL for serving uploaded files (e.g. "http://localhost:8081").
func NewUploadHandler(uploadDir, baseURL string) *UploadHandler {
	// Ensure upload directory exists
	os.MkdirAll(uploadDir, 0755)
	return &UploadHandler{uploadDir: uploadDir, baseURL: baseURL}
}

// Upload handles POST /api/upload
// Accepts multipart form with field "file".
// Returns JSON with the public URL of the uploaded file.
func (h *UploadHandler) Upload(w http.ResponseWriter, r *http.Request) {
	// Max 10MB
	r.ParseMultipartForm(10 << 20)

	file, header, err := r.FormFile("file")
	if err != nil {
		writeError(w, http.StatusBadRequest, "No file provided. Use form field 'file'.")
		return
	}
	defer file.Close()

	// Validate file type
	ext := strings.ToLower(filepath.Ext(header.Filename))
	allowed := map[string]bool{
		".jpg": true, ".jpeg": true, ".png": true, ".gif": true,
		".webp": true, ".svg": true, ".ico": true,
	}
	if !allowed[ext] {
		writeError(w, http.StatusBadRequest, "File type not allowed. Use: jpg, jpeg, png, gif, webp, svg, ico")
		return
	}

	// Generate unique filename
	randBytes := make([]byte, 16)
	rand.Read(randBytes)
	filename := hex.EncodeToString(randBytes) + ext

	// Save file
	destPath := filepath.Join(h.uploadDir, filename)
	dest, err := os.Create(destPath)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to save file")
		return
	}
	defer dest.Close()

	if _, err := io.Copy(dest, file); err != nil {
		writeError(w, http.StatusInternalServerError, "Failed to write file")
		return
	}

	// Return public URL
	publicURL := fmt.Sprintf("%s/uploads/%s", h.baseURL, filename)
	writeJSON(w, http.StatusOK, map[string]string{
		"url":      publicURL,
		"filename": filename,
	})
}

// ServeFiles returns an http.Handler that serves uploaded files.
func (h *UploadHandler) ServeFiles() http.Handler {
	return http.StripPrefix("/uploads/", http.FileServer(http.Dir(h.uploadDir)))
}
