# Implementation Plan: Dynamic Portfolio Admin

## Overview

Transform the existing static SvelteKit portfolio into a dynamic system with a Go backend (chi router, pgx, JWT auth), PostgreSQL database, and an admin panel. The project is restructured into `app/` (SvelteKit frontend) and `api/` (Go backend). Implementation follows an incremental approach: project restructure → backend foundation → models/repositories → handlers → seed data → frontend API client → public portfolio integration → admin panel → validation/polish → documentation.

## Tasks

- [ ] 1. Project Restructure
  - [ ] 1.1 Move existing SvelteKit project files into `app/` subdirectory (src, static, package.json, svelte.config.js, tsconfig.json, vite.config.ts, .npmrc, Dockerfile)
    - Preserve all existing source files, configurations, and static assets
    - _Requirements: 1.1, 1.3_
  - [ ] 1.2 Update any path references in app configuration files to work from the new `app/` location
    - _Requirements: 1.3_
  - [ ] 1.3 Verify the frontend builds and runs successfully from `app/` directory
    - _Requirements: 1.3_
  - [ ] 1.4 Initialize Go module in `api/` directory with `go mod init` and install dependencies (chi, pgx, bcrypt, jwt, uuid)
    - _Requirements: 1.2_
  - [ ] 1.5 Create `api/cmd/server/main.go` entry point with basic server startup
    - _Requirements: 1.2_
  - [ ] 1.6 Create `docker-compose.yml` at project root with PostgreSQL 16 service
    - _Requirements: 2.1_
  - [ ] 1.7 Create root `.gitignore` updated for both app/ and api/ directories
    - _Requirements: 1.1, 1.2_
  - [ ] 1.8 Remove Supabase dependency from `app/package.json` and delete `supabaseClient.ts`
    - _Requirements: 7.5_

- [ ] 2. Backend Foundation
  - [ ] 2.1 Create `api/internal/config/config.go` to load environment variables (DATABASE_URL, JWT_SECRET, JWT_EXPIRY, SERVER_PORT, CORS_ORIGIN, ADMIN_DEFAULT_PASSWORD) with validation
    - Exit with descriptive error if required vars are missing
    - _Requirements: 13.1, 13.2_
  - [ ] 2.2 Create `api/internal/database/postgres.go` with connection pool setup using pgx
    - _Requirements: 1.2_
  - [ ] 2.3 Create `api/internal/database/migrate.go` with file-based migration runner that executes SQL files in order
    - _Requirements: 2.1_
  - [ ] 2.4 Create migration file `001_create_admin_users.sql`
    - _Requirements: 2.2_
  - [ ] 2.5 Create migration file `002_create_identity_console.sql`
    - Include order_number, is_published, created_at, updated_at columns
    - _Requirements: 2.2, 2.3, 2.4, 2.5, 2.6_
  - [ ] 2.6 Create migration file `003_create_capability_snapshots.sql`
    - _Requirements: 2.2, 2.4, 2.5, 2.6_
  - [ ] 2.7 Create migration file `004_create_implementation_strengths.sql`
    - Include JSONB column for bullet_points
    - _Requirements: 2.2, 2.3, 2.4, 2.5, 2.6_
  - [ ] 2.8 Create migration file `005_create_professional_dossier.sql`
    - _Requirements: 2.2, 2.5, 2.6_
  - [ ] 2.9 Create migration file `006_create_education.sql`
    - Include JSONB column for tags
    - _Requirements: 2.2, 2.3, 2.4, 2.5, 2.6_
  - [ ] 2.10 Create migration file `007_create_work_experiences.sql`
    - Include JSONB columns for bullet_points, tech_tags, gallery_images
    - _Requirements: 2.2, 2.3, 2.4, 2.5, 2.6_
  - [ ] 2.11 Create migration file `008_create_projects.sql`
    - Include JSONB column for tech_tags
    - _Requirements: 2.2, 2.3, 2.4, 2.5, 2.6_
  - [ ] 2.12 Create migration file `009_create_certifications.sql`
    - Include JSONB column for skills
    - _Requirements: 2.2, 2.3, 2.4, 2.5, 2.6_
  - [ ] 2.13 Create migration file `010_create_publications.sql`
    - Include JSONB column for tags
    - _Requirements: 2.2, 2.3, 2.4, 2.5, 2.6_
  - [ ] 2.14 Create migration file `011_create_contact_info.sql`
    - _Requirements: 2.2, 2.5, 2.6_
  - [ ] 2.15 Create migration file `012_create_contact_messages.sql`
    - _Requirements: 2.2_
  - [ ] 2.16 Create migration file `013_create_site_settings.sql`
    - _Requirements: 2.2_
  - [ ] 2.17 Create `api/internal/service/auth_service.go` with JWT generation, JWT validation, and bcrypt hash/compare functions
    - Use bcrypt cost factor >= 10
    - _Requirements: 4.1, 4.6_
  - [ ] 2.18 Create `api/internal/middleware/auth.go` with JWT authentication middleware for chi
    - Return 401 for missing, expired, or invalid tokens
    - _Requirements: 4.3, 4.4_
  - [ ] 2.19 Create `api/internal/middleware/cors.go` with CORS middleware configuration
    - Allow Frontend origin from CORS_ORIGIN env var
    - _Requirements: 5.5_
  - [ ] 2.20 Create `api/internal/handler/health.go` with GET /api/health endpoint returning server status and DB connectivity
    - _Requirements: 1.4_
  - [ ] 2.21 Create `api/internal/handler/auth.go` with POST /api/auth/login and GET /api/auth/me handlers
    - Never return password_hash in /me response
    - _Requirements: 4.1, 4.2, 4.5_
  - [ ] 2.22 Create `api/internal/router/router.go` with chi router setup, middleware registration, and route mounting
    - _Requirements: 1.2, 5.5_
  - [ ] 2.23 Wire up main.go to load config, connect DB, run migrations, set up router, and start HTTP server
    - _Requirements: 1.2, 13.1_
  - [ ] 2.24 Create `api/.env.example` with all environment variables documented
    - _Requirements: 13.3_
  - [ ] 2.25 Create `api/Makefile` with targets: run, build, migrate, seed
    - _Requirements: 14.1_

- [ ] 3. Checkpoint - Backend foundation verification
  - Ensure all tests pass, ask the user if questions arise.

- [ ] 4. Backend CRUD - Models and Repositories
  - [ ] 4.1 Create `api/internal/models/` with Go structs for all 13 database tables (user, identity, capability, strength, dossier, education, experience, project, certification, publication, contact, message, settings)
    - Include JSON tags for serialization
    - _Requirements: 2.2, 2.3_
  - [ ] 4.2 Create `api/internal/repository/user_repo.go` with FindByEmail and Create methods
    - _Requirements: 4.1_
  - [ ] 4.3 Create `api/internal/repository/identity_repo.go` with full CRUD + publish/unpublish/reorder + list published
    - _Requirements: 5.1, 5.2, 6.1, 6.4, 6.5_
  - [ ] 4.4 Create `api/internal/repository/capability_repo.go` with full CRUD + publish/unpublish/reorder + list published
    - _Requirements: 5.1, 5.2, 6.1, 6.4, 6.5_
  - [ ] 4.5 Create `api/internal/repository/strength_repo.go` with full CRUD + publish/unpublish/reorder + list published
    - _Requirements: 5.1, 5.2, 6.1, 6.4, 6.5_
  - [ ] 4.6 Create `api/internal/repository/dossier_repo.go` with full CRUD + publish/unpublish + list published
    - _Requirements: 5.1, 5.2, 6.1, 6.4_
  - [ ] 4.7 Create `api/internal/repository/education_repo.go` with full CRUD + publish/unpublish/reorder + list published
    - _Requirements: 5.1, 5.2, 6.1, 6.4, 6.5_
  - [ ] 4.8 Create `api/internal/repository/experience_repo.go` with full CRUD + publish/unpublish/reorder + list published
    - _Requirements: 5.1, 5.2, 6.1, 6.4, 6.5_
  - [ ] 4.9 Create `api/internal/repository/project_repo.go` with full CRUD + publish/unpublish/reorder + list published
    - _Requirements: 5.1, 5.2, 6.1, 6.4, 6.5_
  - [ ] 4.10 Create `api/internal/repository/certification_repo.go` with full CRUD + publish/unpublish/reorder + list published
    - _Requirements: 5.1, 5.2, 6.1, 6.4, 6.5_
  - [ ] 4.11 Create `api/internal/repository/publication_repo.go` with full CRUD + publish/unpublish/reorder + list published
    - _Requirements: 5.1, 5.2, 6.1, 6.4, 6.5_
  - [ ] 4.12 Create `api/internal/repository/contact_repo.go` with CRUD for contact_info
    - _Requirements: 5.1, 6.1_
  - [ ] 4.13 Create `api/internal/repository/message_repo.go` with List, GetByID, Create, MarkRead, Delete, UnreadCount
    - _Requirements: 5.3, 10.1, 10.2, 10.3_
  - [ ] 4.14 Create `api/internal/repository/settings_repo.go` with Get and Update methods
    - _Requirements: 11.1, 11.2_

- [ ] 5. Backend CRUD - Handlers
  - [ ] 5.1 Create `api/internal/handler/public.go` with GET handlers for all public endpoints (identity, capabilities, strengths, dossier, education, experiences, projects, certifications, publications, contact-info, site-settings)
    - Return only published records ordered by order_number ascending
    - _Requirements: 5.1, 5.2_
  - [ ] 5.2 Create public POST /api/contact/messages handler with request validation
    - Validate name, email, message fields; return 400 for invalid input
    - _Requirements: 5.3, 5.4_
  - [ ] 5.3 Create `api/internal/handler/admin.go` with admin CRUD handlers for identity section (list all, get by ID, create, update, delete, publish, unpublish, reorder)
    - Validate required fields on create/update
    - _Requirements: 6.1, 6.2, 6.3, 6.4, 6.5_
  - [ ] 5.4 Add admin CRUD handlers for capabilities, strengths, dossier, education, experiences, projects, certifications, publications, contact-info sections
    - _Requirements: 6.1, 6.2, 6.3, 6.4, 6.5_
  - [ ] 5.5 Create admin message handlers (list, get, mark read, delete, unread count)
    - _Requirements: 10.1, 10.2, 10.3_
  - [ ] 5.6 Create admin settings handlers (get, update)
    - _Requirements: 11.1, 11.2_
  - [ ] 5.7 Create admin dashboard stats handler (GET /api/admin/dashboard/stats returning record counts)
    - _Requirements: 8.3_
  - [ ] 5.8 Register all new handlers in router.go with appropriate middleware (public routes open, admin routes behind JWT middleware)
    - _Requirements: 6.6_

- [ ] 6. Checkpoint - Backend CRUD verification
  - Ensure all tests pass, ask the user if questions arise.

- [ ]* 6.1 Write property test for round-trip data integrity
  - **Property 1: Round-trip data integrity**
  - **Validates: Requirements 6.1, 11.2, 12.3**

- [ ]* 6.2 Write property test for auth isolation
  - **Property 2: Auth isolation**
  - **Validates: Requirements 4.1, 4.2, 4.3, 4.4, 6.6**

- [ ]* 6.3 Write property test for publish filtering
  - **Property 3: Publish filtering**
  - **Validates: Requirements 5.2, 6.4**

- [ ]* 6.4 Write property test for order consistency
  - **Property 4: Order consistency**
  - **Validates: Requirements 6.5**

- [ ]* 6.5 Write property test for input validation
  - **Property 7: Input validation**
  - **Validates: Requirements 5.3, 5.4, 6.2**

- [ ]* 6.6 Write property test for configuration fail-fast
  - **Property 8: Configuration fail-fast**
  - **Validates: Requirements 13.2**

- [ ] 7. Seed Data
  - [ ] 7.1 Create `api/seeds/seed.go` with seed function that populates all tables from existing portfolio.ts data
    - _Requirements: 3.1_
  - [ ] 7.2 Include default admin user creation (***REMOVED***) with bcrypt-hashed password from ADMIN_DEFAULT_PASSWORD env var
    - _Requirements: 3.2_
  - [ ] 7.3 Include default site_settings record with current meta title and description
    - _Requirements: 3.3_
  - [ ] 7.4 Add idempotency check (skip if data already exists) to prevent duplicate seeding
    - _Requirements: 3.4_
  - [ ] 7.5 Wire seed command into main.go (triggered by CLI flag or Makefile target)
    - _Requirements: 3.1_

- [ ]* 7.6 Write property test for seed idempotency
  - **Property 5: Seed idempotency**
  - **Validates: Requirements 3.4**

- [ ]* 7.7 Write property test for password security
  - **Property 6: Password security**
  - **Validates: Requirements 4.5, 4.6**

- [ ] 8. Checkpoint - Backend complete verification
  - Ensure all tests pass, ask the user if questions arise.

- [ ] 9. Frontend - API Client and Types
  - [ ] 9.1 Create `app/src/lib/types/portfolio.ts` with TypeScript interfaces matching all API response shapes
    - _Requirements: 7.1_
  - [ ] 9.2 Create `app/src/lib/types/admin.ts` with admin-specific types (login request/response, CRUD payloads, dashboard stats)
    - _Requirements: 8.4, 8.5_
  - [ ] 9.3 Create `app/src/lib/api/client.ts` with base fetch wrapper (base URL from env, automatic JSON headers, JWT injection from store)
    - Redirect to login on 401 responses
    - _Requirements: 9.4_
  - [ ] 9.4 Create `app/src/lib/api/public.ts` with functions for all public API calls
    - _Requirements: 7.1_
  - [ ] 9.5 Create `app/src/lib/api/admin.ts` with functions for all admin API calls (CRUD, publish, reorder, messages, settings)
    - _Requirements: 8.4, 8.5, 8.6_
  - [ ] 9.6 Create `app/src/lib/stores/auth.ts` with JWT token store (localStorage persistence, login/logout helpers, isAuthenticated derived)
    - _Requirements: 9.2, 9.3_
  - [ ] 9.7 Create `app/src/lib/stores/toast.ts` with toast notification store
    - _Requirements: 8.5_

- [ ] 10. Frontend - Public Portfolio Integration
  - [ ] 10.1 Modify `app/src/routes/+page.svelte` to fetch data from public API endpoints instead of importing static portfolio.ts
    - _Requirements: 7.1_
  - [ ] 10.2 Add loading states (skeleton/spinner) while API data is being fetched
    - _Requirements: 7.3_
  - [ ] 10.3 Add error state handling when API is unreachable
    - _Requirements: 7.4_
  - [ ] 10.4 Update contact form submission to use POST /api/contact/messages
    - _Requirements: 7.5_
  - [ ] 10.5 Update `app/src/routes/+layout.svelte` to fetch site settings and navigation from API
    - _Requirements: 11.3_
  - [ ] 10.6 Remove the static `$lib/data/portfolio.ts` file (keep as reference/backup until confirmed working)
    - _Requirements: 7.1, 7.2_

- [ ] 11. Frontend - Admin Panel Layout and Auth
  - [ ] 11.1 Create `app/src/lib/styles/admin.css` with dark console-style theme (dark background, monospace accents, card-based layout)
    - _Requirements: 8.2_
  - [ ] 11.2 Create `app/src/routes/admin/+layout.svelte` with sidebar navigation, auth guard (redirect to login if no token), and admin CSS import
    - _Requirements: 9.1_
  - [ ] 11.3 Create `app/src/lib/components/admin/Sidebar.svelte` with navigation links to all admin sections and unread message count badge
    - _Requirements: 8.2, 10.4_
  - [ ] 11.4 Create `app/src/routes/admin/login/+page.svelte` with email/password form, login API call, and redirect to dashboard on success
    - _Requirements: 9.2_
  - [ ] 11.5 Create `app/src/routes/admin/+page.svelte` that redirects to /admin/dashboard
    - _Requirements: 8.1_
  - [ ] 11.6 Create `app/src/routes/admin/dashboard/+page.svelte` with summary cards showing record counts for each section
    - _Requirements: 8.3_

- [ ] 12. Frontend - Admin CRUD Pages
  - [ ] 12.1 Create reusable admin components: DataTable.svelte, FormEditor.svelte, ImageUrlInput.svelte, PublishToggle.svelte, ReorderList.svelte, DashboardCard.svelte
    - _Requirements: 8.4, 8.5, 8.6, 12.1, 12.2_
  - [ ] 12.2 Create admin identity section pages (list, new, edit) with form fields matching the identity_console schema
    - _Requirements: 8.4, 8.5_
  - [ ] 12.3 Create admin capabilities section pages (list, new, edit)
    - _Requirements: 8.4, 8.5_
  - [ ] 12.4 Create admin strengths section pages (list, new, edit)
    - _Requirements: 8.4, 8.5_
  - [ ] 12.5 Create admin dossier section pages (list, new, edit)
    - _Requirements: 8.4, 8.5_
  - [ ] 12.6 Create admin education section pages (list, new, edit)
    - _Requirements: 8.4, 8.5_
  - [ ] 12.7 Create admin experiences section pages (list, new, edit) including JSONB array editors for bullet_points, tech_tags, gallery_images
    - _Requirements: 8.4, 8.5_
  - [ ] 12.8 Create admin projects section pages (list, new, edit)
    - _Requirements: 8.4, 8.5_
  - [ ] 12.9 Create admin certifications section pages (list, new, edit)
    - _Requirements: 8.4, 8.5_
  - [ ] 12.10 Create admin publications section pages (list, new, edit)
    - _Requirements: 8.4, 8.5_
  - [ ] 12.11 Create admin messages page with table view, read/unread status, and delete action
    - _Requirements: 10.1, 10.2, 10.3_
  - [ ] 12.12 Create admin settings page with form for all site_settings fields
    - _Requirements: 11.1_

- [ ] 13. Checkpoint - Frontend admin verification
  - Ensure all tests pass, ask the user if questions arise.

- [ ] 14. Frontend - Validation and Polish
  - [ ] 14.1 Create `app/src/lib/utils/validation.ts` with form validation helpers (required, email, URL, min/max length)
    - _Requirements: 8.5_
  - [ ] 14.2 Add client-side validation to all admin forms with inline error messages
    - _Requirements: 8.5_
  - [ ] 14.3 Add image URL preview thumbnails to ImageUrlInput component
    - _Requirements: 12.2_
  - [ ] 14.4 Add confirmation dialogs for delete actions
    - _Requirements: 8.6_
  - [ ] 14.5 Add toast notifications for all CRUD operations (success/error feedback)
    - _Requirements: 8.5_
  - [ ] 14.6 Handle JWT expiration in API client (redirect to login on 401 responses)
    - _Requirements: 9.4_

- [ ] 15. Documentation and Final Setup
  - [ ] 15.1 Create `api/README.md` with prerequisites, env setup, database creation, migration, seed, and run instructions
    - _Requirements: 14.1_
  - [ ] 15.2 Create `app/README.md` with prerequisites, env setup, install, dev, and build instructions
    - _Requirements: 14.2_
  - [ ] 15.3 Create root `README.md` with architecture overview, project structure, quick start guide, and links to sub-READMEs
    - _Requirements: 14.3_
  - [ ] 15.4 Create `app/.env.example` with PUBLIC_API_URL variable
    - _Requirements: 14.2_
  - [ ] 15.5 Update `docker-compose.yml` to include volume persistence and optional pgAdmin service
    - _Requirements: 14.1_
  - [ ] 15.6 Final verification: ensure frontend builds, backend compiles, migrations run, seeds execute, and health endpoint responds
    - _Requirements: 14.1, 14.2, 14.3_

- [ ] 16. Final checkpoint - Ensure all tests pass
  - Ensure all tests pass, ask the user if questions arise.

## Notes

- Tasks marked with `*` are optional and can be skipped for faster MVP
- Each task references specific requirements for traceability
- Checkpoints ensure incremental validation between major phases
- Property tests validate universal correctness properties from the design document using Go's `pgregory.net/rapid` library
- Unit tests validate specific examples and edge cases
- The Go backend uses Handler → Repository pattern without a service layer for simple CRUD
- Frontend uses SvelteKit file-based routing with stores for auth state management
- All image fields are URL strings (no file upload handling needed)
- JSONB columns are used for array-like data (bullet_points, tech_tags, skills, gallery_images)

## Task Dependency Graph

```json
{
  "waves": [
    { "id": 0, "tasks": ["1.1", "1.4", "1.6", "1.7"] },
    { "id": 1, "tasks": ["1.2", "1.5", "1.8"] },
    { "id": 2, "tasks": ["1.3", "2.1", "2.24", "2.25"] },
    { "id": 3, "tasks": ["2.2", "2.17"] },
    { "id": 4, "tasks": ["2.3", "2.18", "2.19"] },
    { "id": 5, "tasks": ["2.4", "2.5", "2.6", "2.7", "2.8", "2.9", "2.10", "2.11", "2.12", "2.13", "2.14", "2.15", "2.16"] },
    { "id": 6, "tasks": ["2.20", "2.21", "2.22"] },
    { "id": 7, "tasks": ["2.23"] },
    { "id": 8, "tasks": ["4.1"] },
    { "id": 9, "tasks": ["4.2", "4.3", "4.4", "4.5", "4.6", "4.7", "4.8", "4.9", "4.10", "4.11", "4.12", "4.13", "4.14"] },
    { "id": 10, "tasks": ["5.1", "5.2", "5.5", "5.6", "5.7"] },
    { "id": 11, "tasks": ["5.3", "5.4", "5.8"] },
    { "id": 12, "tasks": ["6.1", "6.2", "6.3", "6.4", "6.5", "6.6"] },
    { "id": 13, "tasks": ["7.1", "7.2", "7.3"] },
    { "id": 14, "tasks": ["7.4", "7.5"] },
    { "id": 15, "tasks": ["7.6", "7.7"] },
    { "id": 16, "tasks": ["9.1", "9.2"] },
    { "id": 17, "tasks": ["9.3", "9.6", "9.7"] },
    { "id": 18, "tasks": ["9.4", "9.5"] },
    { "id": 19, "tasks": ["10.1", "10.5"] },
    { "id": 20, "tasks": ["10.2", "10.3", "10.4"] },
    { "id": 21, "tasks": ["10.6", "11.1"] },
    { "id": 22, "tasks": ["11.2", "11.3", "11.4"] },
    { "id": 23, "tasks": ["11.5", "11.6"] },
    { "id": 24, "tasks": ["12.1"] },
    { "id": 25, "tasks": ["12.2", "12.3", "12.4", "12.5", "12.6", "12.7", "12.8", "12.9", "12.10", "12.11", "12.12"] },
    { "id": 26, "tasks": ["14.1"] },
    { "id": 27, "tasks": ["14.2", "14.3", "14.4", "14.5", "14.6"] },
    { "id": 28, "tasks": ["15.1", "15.2", "15.3", "15.4", "15.5"] },
    { "id": 29, "tasks": ["15.6"] }
  ]
}
```
