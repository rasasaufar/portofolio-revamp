# Requirements Document

## Introduction

Transform the existing static SvelteKit portfolio website into a dynamic portfolio system with a Go backend, PostgreSQL database, and an admin panel. The public-facing portfolio retains its current futuristic/identity console design but sources all content from a REST API. A protected admin panel enables full content management (CRUD, publish/unpublish, reorder) for all portfolio sections. The project is restructured into `app/` (SvelteKit frontend) and `api/` (Go backend).

## Glossary

- **Frontend**: The SvelteKit + TypeScript application located in the `app/` directory that renders the public portfolio and admin panel
- **Backend**: The Go REST API server located in the `api/` directory that handles data persistence, authentication, and business logic
- **Admin_Panel**: The protected section of the Frontend accessible at `/admin` routes for managing portfolio content
- **Public_Portfolio**: The public-facing section of the Frontend that displays portfolio content to visitors
- **Database**: The PostgreSQL database storing all portfolio content, admin credentials, contact messages, and site settings
- **Auth_System**: The JWT-based authentication system using bcrypt password hashing for admin access
- **Migration_Runner**: The component responsible for executing PostgreSQL schema migrations in order
- **Seed_Runner**: The component responsible for populating the Database with initial portfolio content and a default admin account
- **API_Router**: The Go chi router that handles HTTP request routing, middleware, and CORS configuration

## Requirements

### Requirement 1: Project Structure Migration

**User Story:** As a developer, I want the existing SvelteKit project moved to an `app/` subdirectory and a new Go backend created in `api/`, so that the monorepo structure cleanly separates frontend and backend concerns.

#### Acceptance Criteria

1. WHEN the project is restructured, THE Frontend SHALL reside in the `app/` directory with all existing source files, configurations, and static assets preserved
2. WHEN the project is restructured, THE Backend SHALL reside in the `api/` directory with a Go module, chi router, and PostgreSQL connection setup
3. WHEN the Frontend is moved to `app/`, THE Frontend SHALL build and run without errors using the same commands as before
4. THE Backend SHALL expose a `GET /api/health` endpoint that returns a JSON response indicating server status and database connectivity

### Requirement 2: Database Schema and Migrations

**User Story:** As a developer, I want PostgreSQL migrations that create all necessary tables, so that the database schema is version-controlled and reproducible.

#### Acceptance Criteria

1. THE Migration_Runner SHALL execute SQL migration files in sequential order based on filename numbering
2. WHEN migrations are applied, THE Database SHALL contain tables for: identity_console, capability_snapshots, implementation_strengths, professional_dossier, education, work_experiences, projects, certifications, publications, contact_info, contact_messages, site_settings, and admin_users
3. THE Database SHALL enforce NOT NULL constraints on required fields and use appropriate data types including JSONB for array-like data (bullet_points, tech_tags, skills, gallery_images)
4. THE Database SHALL include `order_number` integer columns on content tables that support manual ordering
5. THE Database SHALL include `is_published` boolean columns on content tables that support publish/unpublish functionality
6. THE Database SHALL include `created_at` and `updated_at` timestamp columns on all content tables

### Requirement 3: Seed Data

**User Story:** As a developer, I want seed data that populates the database with the current portfolio content and a default admin account, so that the portfolio remains populated after initial setup.

#### Acceptance Criteria

1. WHEN the Seed_Runner executes, THE Database SHALL be populated with all existing portfolio content matching the current static data in `portfolio.ts`
2. WHEN the Seed_Runner executes, THE Database SHALL contain a default admin account with email `***REMOVED***` and a bcrypt-hashed password
3. WHEN the Seed_Runner executes, THE Database SHALL contain default site settings including site title, meta description, and theme mode
4. IF the Seed_Runner is executed on a database that already contains data, THEN THE Seed_Runner SHALL skip insertion to avoid duplicate records

### Requirement 4: Authentication System

**User Story:** As an admin, I want to log in with email and password to access the admin panel, so that only authorized users can manage portfolio content.

#### Acceptance Criteria

1. WHEN a valid email and password are submitted to `POST /api/auth/login`, THE Auth_System SHALL return a JWT access token with a configurable expiration time
2. WHEN an invalid email or password is submitted to `POST /api/auth/login`, THE Auth_System SHALL return a 401 Unauthorized response with an error message
3. WHEN a valid JWT token is included in the Authorization header, THE Auth_System SHALL allow access to protected admin API endpoints
4. WHEN an expired or invalid JWT token is included in the Authorization header, THE Auth_System SHALL return a 401 Unauthorized response
5. WHEN `GET /api/auth/me` is called with a valid token, THE Auth_System SHALL return the authenticated admin user profile without the password hash
6. THE Auth_System SHALL hash passwords using bcrypt with a minimum cost factor of 10

### Requirement 5: Public API Endpoints

**User Story:** As a visitor, I want the public portfolio to load content from the API, so that the displayed information is always current with what the admin has published.

#### Acceptance Criteria

1. THE Backend SHALL expose GET endpoints for each portfolio section: `/api/identity`, `/api/capabilities`, `/api/strengths`, `/api/dossier`, `/api/education`, `/api/experiences`, `/api/projects`, `/api/certifications`, `/api/publications`, `/api/contact-info`
2. WHEN a public GET endpoint is called, THE Backend SHALL return only records where `is_published` is true, ordered by `order_number` ascending
3. WHEN `POST /api/contact/messages` is called with valid name, email, and message fields, THE Backend SHALL store the contact message in the Database and return a success response
4. IF `POST /api/contact/messages` is called with missing or invalid fields, THEN THE Backend SHALL return a 400 Bad Request response with validation error details
5. THE Backend SHALL set CORS headers allowing the Frontend origin to make requests

### Requirement 6: Admin CRUD API Endpoints

**User Story:** As an admin, I want full CRUD operations for all portfolio sections through protected API endpoints, so that I can create, read, update, and delete content.

#### Acceptance Criteria

1. THE Backend SHALL expose protected CRUD endpoints (GET list, GET by ID, POST create, PUT update, DELETE) for each content section under the `/api/admin/` prefix
2. WHEN a create or update request is received, THE Backend SHALL validate required fields and return a 400 Bad Request response with details for invalid submissions
3. WHEN a delete request is received for a valid resource ID, THE Backend SHALL remove the record from the Database and return a success response
4. THE Backend SHALL expose `PATCH /api/admin/{resource}/{id}/publish` and `PATCH /api/admin/{resource}/{id}/unpublish` endpoints to toggle the `is_published` field
5. THE Backend SHALL expose `PUT /api/admin/{resource}/reorder` accepting an ordered array of IDs to update `order_number` values for a given section
6. WHEN any admin endpoint is called without a valid JWT token, THE Backend SHALL return a 401 Unauthorized response

### Requirement 7: Frontend Public Portfolio Dynamic Loading

**User Story:** As a visitor, I want the public portfolio to display the same design as before but with data fetched from the API, so that the visual experience is unchanged while content is dynamic.

#### Acceptance Criteria

1. THE Public_Portfolio SHALL fetch all section data from the Backend public API endpoints on page load
2. THE Public_Portfolio SHALL preserve the existing neo-brutalism/identity console visual design, layout, and animations without modification
3. WHILE the Public_Portfolio is loading data from the API, THE Frontend SHALL display loading indicators in each section
4. IF the Backend is unreachable or returns an error, THEN THE Frontend SHALL display a user-friendly error state without crashing
5. WHEN the contact form is submitted, THE Frontend SHALL send the message to `POST /api/contact/messages` instead of directly to Supabase

### Requirement 8: Admin Panel Interface

**User Story:** As an admin, I want a web-based admin panel with a dark console-style theme, so that I can manage all portfolio content through an intuitive interface.

#### Acceptance Criteria

1. THE Admin_Panel SHALL be accessible at the `/admin` route prefix with sub-routes for login, dashboard, and each content section
2. THE Admin_Panel SHALL use a dark theme with console-style aesthetics including a sidebar navigation menu
3. THE Admin_Panel SHALL display a dashboard at `/admin/dashboard` showing summary cards with record counts for each content section
4. WHEN the admin navigates to a content section, THE Admin_Panel SHALL display records in a table list view with columns for key fields, publish status, and action buttons
5. WHEN the admin clicks Add or Edit, THE Admin_Panel SHALL display a form editor with appropriate input types and client-side validation for required fields
6. THE Admin_Panel SHALL provide Delete, Publish, Unpublish, and Reorder controls for each content section

### Requirement 9: Admin Route Protection

**User Story:** As an admin, I want the admin panel routes to be protected, so that unauthenticated users cannot access content management features.

#### Acceptance Criteria

1. WHEN an unauthenticated user navigates to any `/admin` route except `/admin/login`, THE Frontend SHALL redirect the user to `/admin/login`
2. WHEN the admin successfully logs in, THE Frontend SHALL store the JWT token and redirect to `/admin/dashboard`
3. WHEN the admin clicks logout, THE Frontend SHALL clear the stored JWT token and redirect to `/admin/login`
4. IF the JWT token expires during an admin session, THEN THE Frontend SHALL redirect the user to `/admin/login` with a session expired notification

### Requirement 10: Contact Messages Management

**User Story:** As an admin, I want to view and manage contact messages submitted by visitors, so that I can read and track incoming communications.

#### Acceptance Criteria

1. THE Admin_Panel SHALL display contact messages in a table with columns for name, email, message preview, read status, and submission date
2. WHEN the admin opens a contact message, THE Admin_Panel SHALL mark the message as read by calling the Backend
3. THE Admin_Panel SHALL provide a delete action for individual contact messages
4. THE Admin_Panel SHALL display an unread message count indicator in the sidebar navigation

### Requirement 11: Site Settings Management

**User Story:** As an admin, I want to manage global site settings, so that I can update the site title, meta description, and other site-wide configurations.

#### Acceptance Criteria

1. THE Admin_Panel SHALL provide a settings form at `/admin/settings` for editing site_title, meta_description, favicon_url, logo_url, footer_text, theme_mode, and maintenance_mode
2. WHEN site settings are saved, THE Backend SHALL update the single site_settings record in the Database
3. THE Public_Portfolio SHALL use site settings from the API for the page title, meta description, and footer text

### Requirement 12: Image Handling

**User Story:** As an admin, I want to provide image URLs as strings for all image fields, so that I can reference externally hosted images without file upload complexity.

#### Acceptance Criteria

1. THE Admin_Panel SHALL render image fields as text inputs accepting URL strings
2. WHEN an image URL is provided, THE Admin_Panel SHALL display a preview thumbnail next to the input field
3. THE Backend SHALL store image references as string URLs in the Database without performing file upload or storage operations

### Requirement 13: Backend Configuration

**User Story:** As a developer, I want the backend configured through environment variables, so that database credentials, JWT secrets, and server settings are not hardcoded.

#### Acceptance Criteria

1. THE Backend SHALL read configuration from environment variables: DATABASE_URL, JWT_SECRET, JWT_EXPIRY, SERVER_PORT, CORS_ORIGIN, and ADMIN_DEFAULT_PASSWORD
2. IF a required environment variable is missing at startup, THEN THE Backend SHALL log a descriptive error message and exit with a non-zero status code
3. THE Backend SHALL provide a `.env.example` file documenting all required and optional environment variables with descriptions

### Requirement 14: Setup Documentation

**User Story:** As a developer, I want a comprehensive README with setup instructions, so that the project can be set up and run by any team member.

#### Acceptance Criteria

1. THE Backend SHALL include a README documenting: prerequisites, environment setup, database creation, migration execution, seed execution, and server startup commands
2. THE Frontend SHALL include a README documenting: prerequisites, environment setup, dependency installation, development server startup, and build commands
3. THE project root SHALL include a README documenting: overall architecture, project structure, quick start guide, and links to frontend and backend READMEs
