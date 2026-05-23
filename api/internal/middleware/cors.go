package middleware

import (
	"net/http"

	"github.com/go-chi/cors"
)

// CORSHandler returns a configured CORS middleware handler.
// It supports multiple origins by also allowing common dev ports.
func CORSHandler(allowedOrigin string) func(http.Handler) http.Handler {
	origins := []string{allowedOrigin}
	// Also allow common Vite dev ports
	if allowedOrigin == "http://localhost:5173" {
		origins = append(origins, "http://localhost:5174", "http://localhost:5175", "http://localhost:4173")
	}

	return cors.Handler(cors.Options{
		AllowedOrigins:   origins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-Request-ID"},
		ExposedHeaders:   []string{"Link", "X-Total-Count"},
		AllowCredentials: true,
		MaxAge:           300,
	})
}
