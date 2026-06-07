package router

import (
	"fmt"
	"os"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/portfolio/api/internal/handler"
	"github.com/portfolio/api/internal/middleware"
	"github.com/portfolio/api/internal/repository"
	"github.com/portfolio/api/internal/service"
)

// New creates and configures the chi router with all routes and middleware.
func New(db *pgxpool.Pool, authService *service.AuthService, corsOrigin string) *chi.Mux {
	r := chi.NewRouter()

	// Global middleware
	r.Use(chimiddleware.Logger)
	r.Use(chimiddleware.Recoverer)
	r.Use(chimiddleware.RequestID)
	r.Use(chimiddleware.RealIP)
	r.Use(middleware.CORSHandler(corsOrigin))

	// Repositories
	userRepo := repository.NewUserRepository(db)

	// Upload handler
	uploadDir := getUploadDir()
	baseURL := os.Getenv("PUBLIC_URL")
	if baseURL == "" {
		serverPort := os.Getenv("SERVER_PORT")
		if serverPort == "" {
			serverPort = "8080"
		}
		baseURL = fmt.Sprintf("http://localhost:%s", serverPort)
	}
	uploadHandler := handler.NewUploadHandler(uploadDir, baseURL)

	// Handlers
	healthHandler := handler.NewHealthHandler(db)
	authHandler := handler.NewAuthHandler(userRepo, authService)
	publicHandler := handler.NewPublicHandler(db)
	adminHandler := handler.NewAdminHandler(db)

	// Serve uploaded files (public)
	r.Handle("/uploads/*", uploadHandler.ServeFiles())

	// Routes
	r.Route("/api", func(r chi.Router) {
		// Health check (public)
		r.Get("/health", healthHandler.Check)

		// Auth routes
		r.Post("/auth/login", authHandler.Login)

		// Protected auth routes
		r.Group(func(r chi.Router) {
			r.Use(middleware.JWTAuth(authService))
			r.Get("/auth/me", authHandler.Me)
			r.Put("/auth/password", authHandler.ChangePassword)
		})

		// Upload (protected)
		r.Group(func(r chi.Router) {
			r.Use(middleware.JWTAuth(authService))
			r.Post("/upload", uploadHandler.Upload)
		})

		// Public portfolio endpoints
		r.Get("/identity", publicHandler.GetIdentity)
		r.Get("/capabilities", publicHandler.GetCapabilities)
		r.Get("/strengths", publicHandler.GetStrengths)
		r.Get("/dossier", publicHandler.GetDossier)
		r.Get("/education", publicHandler.GetEducation)
		r.Get("/experiences", publicHandler.GetExperiences)
		r.Get("/projects", publicHandler.GetProjects)
		r.Get("/certifications", publicHandler.GetCertifications)
		r.Get("/publications", publicHandler.GetPublications)
		r.Get("/contact-info", publicHandler.GetContactInfo)
		r.Get("/site-settings", publicHandler.GetSiteSettings)
		r.Post("/contact/messages", publicHandler.CreateContactMessage)

		// Admin protected routes
		r.Route("/admin", func(r chi.Router) {
			r.Use(middleware.JWTAuth(authService))

			// Dashboard
			r.Get("/dashboard/stats", adminHandler.DashboardStats)

			// Generic CRUD for all resources
			r.Get("/{resource}", adminHandler.List)
			r.Post("/{resource}", adminHandler.Create)
			r.Put("/{resource}/reorder", adminHandler.Reorder)
			r.Get("/{resource}/{id}", adminHandler.GetByID)
			r.Put("/{resource}/{id}", adminHandler.Update)
			r.Delete("/{resource}/{id}", adminHandler.Delete)
			r.Patch("/{resource}/{id}/publish", adminHandler.Publish)
			r.Patch("/{resource}/{id}/unpublish", adminHandler.Unpublish)

			// Messages specific
			r.Patch("/messages/{id}/read", adminHandler.MarkRead)
		})
	})

	return r
}

func getUploadDir() string {
	if _, err := os.Stat("uploads"); err == nil {
		return "uploads"
	}
	if _, err := os.Stat("../../uploads"); err == nil {
		return "../../uploads"
	}
	os.MkdirAll("uploads", 0755)
	return "uploads"
}
