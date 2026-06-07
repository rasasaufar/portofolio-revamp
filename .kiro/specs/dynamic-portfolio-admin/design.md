# Design Document

## Overview

This document details the implementation plan for transforming the static SvelteKit portfolio into a dynamic system with a Go backend, PostgreSQL database, and admin panel. The public-facing portfolio retains its current futuristic/identity console design but sources all content from a REST API. A protected admin panel enables full content management (CRUD, publish/unpublish, reorder) for all portfolio sections. The project is restructured into `app/` (SvelteKit frontend) and `api/` (Go backend).

### Technical Decisions

**Backend (Go)**
- **Router**: `go-chi/chi` v5 — lightweight, idiomatic, middleware-friendly
- **Database driver**: `jackc/pgx` v5 — high-performance PostgreSQL driver with native types
- **Password hashing**: `golang.org/x/crypto/bcrypt` — standard bcrypt implementation
- **JWT**: `golang-jwt/jwt` v5 — widely used JWT library
- **UUID**: `google/uuid` — UUID generation for primary keys
- **Configuration**: Environment variables loaded via `os.Getenv` with validation at startup
- **Migrations**: Custom file-based runner reading `.sql` files from `migrations/` directory
- **Architecture**: Handler → Repository pattern (no service layer for simple CRUD, auth service for JWT/bcrypt logic)

**Frontend (SvelteKit)**
- **Adapter**: `@sveltejs/adapter-node` for SSR deployment
- **API client**: Fetch-based wrapper in `$lib/api/` with automatic JWT header injection
- **Auth state**: Svelte store backed by localStorage for JWT persistence
- **Admin routing**: SvelteKit file-based routing under `src/routes/admin/`
- **Styling**: Existing `neo-brutalism.css` for public, new `admin.css` for admin panel
- **Form handling**: Client-side validation with Svelte reactive statements
- **Supabase removal**: Remove `@supabase/supabase-js` dependency; contact form uses Go API

**Infrastructure**
- **Docker Compose**: PostgreSQL 16 container for local development
- **CORS**: Configured in Go middleware, allowing Frontend origin
- **Environment**: `.env` files for both `app/` and `api/` (gitignored)

## Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                        Browser                               │
├──────────────────────┬──────────────────────────────────────┤
│   Public Portfolio   │         Admin Panel                   │
│   (SvelteKit SSR)    │    (SvelteKit SPA-like)              │
└──────────┬───────────┴──────────────┬───────────────────────┘
           │                          │
           │  HTTP (REST JSON)        │  HTTP (REST JSON + JWT)
           ▼                          ▼
┌─────────────────────────────────────────────────────────────┐
│                    Go Backend (chi)                           │
│  ┌──────────┐  ┌──────────┐  ┌──────────┐  ┌────────────┐ │
│  │  Public   │  │  Admin   │  │   Auth   │  │ Middleware │ │
│  │ Handlers  │  │ Handlers │  │ Handlers │  │ (JWT/CORS) │ │
│  └─────┬────┘  └─────┬────┘  └─────┬────┘  └────────────┘ │
│        └──────────────┼─────────────┘                        │
│                       ▼                                      │
│              ┌─────────────────┐                             │
│              │   Repository    │                             │
│              │     Layer       │                             │
│              └────────┬────────┘                             │
└───────────────────────┼─────────────────────────────────────┘
                        │
                        ▼
              ┌─────────────────┐
              │   PostgreSQL    │
              └─────────────────┘
```

### Implementation Phases

**Phase 1: Project Restructure** — Move SvelteKit to `app/`, init Go module in `api/`, docker-compose with PostgreSQL

**Phase 2: Backend Foundation** — Config, DB connection, health endpoint, migrations, seed, auth service

**Phase 3: Backend CRUD** — Models, repositories, public GET handlers, admin CRUD handlers, contact messages

**Phase 4: Frontend Public Integration** — API client, dynamic data fetching, loading/error states, remove Supabase

**Phase 5: Admin Panel** — Layout, login, dashboard, CRUD pages for all sections, messages, settings, auth guards

**Phase 6: Polish** — Validation, image preview, reorder controls, documentation

## Components and Interfaces

### Project Structure

```
portfolio-svelte/
├── app/                          # Frontend (SvelteKit + TypeScript)
│   ├── src/
│   │   ├── lib/
│   │   │   ├── api/
│   │   │   │   ├── client.ts          # Base HTTP client with auth headers
│   │   │   │   ├── public.ts          # Public API calls
│   │   │   │   └── admin.ts           # Admin API calls (CRUD helpers)
│   │   │   ├── components/
│   │   │   │   ├── public/            # Existing public components (refactored)
│   │   │   │   └── admin/
│   │   │   │       ├── Sidebar.svelte
│   │   │   │       ├── DashboardCard.svelte
│   │   │   │       ├── DataTable.svelte
│   │   │   │       ├── FormEditor.svelte
│   │   │   │       ├── ImageUrlInput.svelte
│   │   │   │       ├── PublishToggle.svelte
│   │   │   │       └── ReorderList.svelte
│   │   │   ├── stores/
│   │   │   │   ├── auth.ts            # JWT token store + helpers
│   │   │   │   └── toast.ts           # Toast notification store
│   │   │   ├── types/
│   │   │   │   ├── portfolio.ts       # Public data types
│   │   │   │   └── admin.ts           # Admin-specific types
│   │   │   ├── utils/
│   │   │   │   ├── validation.ts      # Form validation helpers
│   │   │   │   └── format.ts          # Date/text formatting
│   │   │   ├── styles/
│   │   │   │   ├── neo-brutalism.css  # Existing public styles
│   │   │   │   └── admin.css          # Admin panel styles
│   │   │   ├── assets/
│   │   │   │   └── favicon.svg
│   │   │   └── index.ts
│   │   ├── routes/
│   │   │   ├── +layout.svelte         # Public layout (existing)
│   │   │   ├── +page.svelte           # Public portfolio (modified to fetch API)
│   │   │   └── admin/
│   │   │       ├── +layout.svelte     # Admin layout with sidebar
│   │   │       ├── +page.svelte       # Redirect to /admin/dashboard
│   │   │       ├── login/+page.svelte
│   │   │       ├── dashboard/+page.svelte
│   │   │       ├── identity/          # List, new, [id] pages
│   │   │       ├── capabilities/
│   │   │       ├── strengths/
│   │   │       ├── dossier/
│   │   │       ├── education/
│   │   │       ├── experiences/
│   │   │       ├── projects/
│   │   │       ├── certifications/
│   │   │       ├── publications/
│   │   │       ├── messages/+page.svelte
│   │   │       └── settings/+page.svelte
│   │   ├── app.html
│   │   └── app.d.ts
│   ├── static/
│   ├── package.json
│   ├── svelte.config.js
│   ├── tsconfig.json
│   └── vite.config.ts
│
├── api/                               # Backend (Go + PostgreSQL)
│   ├── cmd/server/main.go            # Entry point
│   ├── internal/
│   │   ├── config/config.go          # Env variable loading
│   │   ├── database/
│   │   │   ├── postgres.go           # Connection pool setup
│   │   │   └── migrate.go            # Migration runner
│   │   ├── middleware/
│   │   │   ├── auth.go               # JWT validation middleware
│   │   │   └── cors.go               # CORS middleware
│   │   ├── models/                    # Go structs for all 13 tables
│   │   ├── repository/               # Data access layer (one per model)
│   │   ├── handler/
│   │   │   ├── auth.go
│   │   │   ├── health.go
│   │   │   ├── public.go             # All public GET handlers
│   │   │   └── admin.go              # All admin CRUD handlers
│   │   ├── service/
│   │   │   └── auth_service.go       # JWT generation/validation, bcrypt
│   │   └── router/router.go          # Chi router setup with all routes
│   ├── migrations/                    # 001-013 SQL migration files
│   ├── seeds/seed.go                  # Seed data from portfolio.ts
│   ├── .env.example
│   ├── go.mod
│   ├── Makefile
│   └── README.md
│
├── .gitignore
├── docker-compose.yml                 # PostgreSQL 16 container
└── README.md
```

### API Endpoint Design

#### Authentication

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | /api/auth/login | Authenticate admin, return JWT |
| GET | /api/auth/me | Get current admin profile (protected) |

#### Public Endpoints (No auth required)

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | /api/identity | Get published identity console data |
| GET | /api/capabilities | Get published capability snapshots |
| GET | /api/strengths | Get published implementation strengths |
| GET | /api/dossier | Get published professional dossier |
| GET | /api/education | Get published education records |
| GET | /api/experiences | Get published work experiences |
| GET | /api/projects | Get published projects |
| GET | /api/certifications | Get published certifications |
| GET | /api/publications | Get published publications |
| GET | /api/contact-info | Get published contact information |
| GET | /api/site-settings | Get site settings (public subset) |
| POST | /api/contact/messages | Submit a contact message |

#### Admin Protected Endpoints (JWT required)

For each resource (identity, capabilities, strengths, dossier, education, experiences, projects, certifications, publications, contact-info):

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | /api/admin/{resource} | List all records (including unpublished) |
| GET | /api/admin/{resource}/{id} | Get single record by ID |
| POST | /api/admin/{resource} | Create new record |
| PUT | /api/admin/{resource}/{id} | Update record |
| DELETE | /api/admin/{resource}/{id} | Delete record |
| PATCH | /api/admin/{resource}/{id}/publish | Set is_published = true |
| PATCH | /api/admin/{resource}/{id}/unpublish | Set is_published = false |
| PUT | /api/admin/{resource}/reorder | Update order_number for all items |

Additional admin endpoints:

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | /api/admin/messages | List all contact messages |
| GET | /api/admin/messages/{id} | Get single message |
| PATCH | /api/admin/messages/{id}/read | Mark message as read |
| DELETE | /api/admin/messages/{id} | Delete message |
| GET | /api/admin/messages/unread-count | Get unread message count |
| GET | /api/admin/settings | Get full site settings |
| PUT | /api/admin/settings | Update site settings |
| GET | /api/admin/dashboard/stats | Get record counts for dashboard |

#### Health

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | /api/health | Server status + DB connectivity check |

## Data Models

### admin_users

| Column | Type | Constraints |
|--------|------|-------------|
| id | UUID | PRIMARY KEY, DEFAULT gen_random_uuid() |
| email | VARCHAR(255) | NOT NULL, UNIQUE |
| password_hash | VARCHAR(255) | NOT NULL |
| name | VARCHAR(100) | NOT NULL |
| created_at | TIMESTAMPTZ | NOT NULL, DEFAULT NOW() |
| updated_at | TIMESTAMPTZ | NOT NULL, DEFAULT NOW() |

### identity_console

| Column | Type | Constraints |
|--------|------|-------------|
| id | UUID | PRIMARY KEY, DEFAULT gen_random_uuid() |
| name | VARCHAR(100) | NOT NULL |
| role | VARCHAR(200) | NOT NULL |
| headline | TEXT | NOT NULL |
| description | TEXT | NOT NULL |
| avatar_url | VARCHAR(500) | |
| current_focus | JSONB | NOT NULL, DEFAULT '[]' |
| availability_text | VARCHAR(200) | |
| cta_buttons | JSONB | NOT NULL, DEFAULT '[]' |
| skill_stack | JSONB | NOT NULL, DEFAULT '[]' |
| order_number | INTEGER | NOT NULL, DEFAULT 0 |
| is_published | BOOLEAN | NOT NULL, DEFAULT true |
| created_at | TIMESTAMPTZ | NOT NULL, DEFAULT NOW() |
| updated_at | TIMESTAMPTZ | NOT NULL, DEFAULT NOW() |

### capability_snapshots

| Column | Type | Constraints |
|--------|------|-------------|
| id | UUID | PRIMARY KEY, DEFAULT gen_random_uuid() |
| label | VARCHAR(100) | NOT NULL |
| value | VARCHAR(50) | NOT NULL |
| description | VARCHAR(255) | |
| order_number | INTEGER | NOT NULL, DEFAULT 0 |
| is_published | BOOLEAN | NOT NULL, DEFAULT true |
| created_at | TIMESTAMPTZ | NOT NULL, DEFAULT NOW() |
| updated_at | TIMESTAMPTZ | NOT NULL, DEFAULT NOW() |

### implementation_strengths

| Column | Type | Constraints |
|--------|------|-------------|
| id | UUID | PRIMARY KEY, DEFAULT gen_random_uuid() |
| title | VARCHAR(200) | NOT NULL |
| description | TEXT | |
| bullet_points | JSONB | NOT NULL, DEFAULT '[]' |
| icon_url | VARCHAR(500) | |
| order_number | INTEGER | NOT NULL, DEFAULT 0 |
| is_published | BOOLEAN | NOT NULL, DEFAULT true |
| created_at | TIMESTAMPTZ | NOT NULL, DEFAULT NOW() |
| updated_at | TIMESTAMPTZ | NOT NULL, DEFAULT NOW() |

### professional_dossier

| Column | Type | Constraints |
|--------|------|-------------|
| id | UUID | PRIMARY KEY, DEFAULT gen_random_uuid() |
| title | VARCHAR(200) | NOT NULL |
| paragraph_1 | TEXT | NOT NULL |
| paragraph_2 | TEXT | |
| paragraph_3 | TEXT | |
| is_published | BOOLEAN | NOT NULL, DEFAULT true |
| created_at | TIMESTAMPTZ | NOT NULL, DEFAULT NOW() |
| updated_at | TIMESTAMPTZ | NOT NULL, DEFAULT NOW() |

### education

| Column | Type | Constraints |
|--------|------|-------------|
| id | UUID | PRIMARY KEY, DEFAULT gen_random_uuid() |
| institution_name | VARCHAR(200) | NOT NULL |
| degree | VARCHAR(200) | NOT NULL |
| major | VARCHAR(200) | |
| years | VARCHAR(50) | NOT NULL |
| gpa | VARCHAR(20) | |
| description | TEXT | |
| image_url | VARCHAR(500) | |
| tags | JSONB | NOT NULL, DEFAULT '[]' |
| order_number | INTEGER | NOT NULL, DEFAULT 0 |
| is_published | BOOLEAN | NOT NULL, DEFAULT true |
| created_at | TIMESTAMPTZ | NOT NULL, DEFAULT NOW() |
| updated_at | TIMESTAMPTZ | NOT NULL, DEFAULT NOW() |

### work_experiences

| Column | Type | Constraints |
|--------|------|-------------|
| id | UUID | PRIMARY KEY, DEFAULT gen_random_uuid() |
| company_name | VARCHAR(200) | NOT NULL |
| position | VARCHAR(200) | NOT NULL |
| dates | VARCHAR(100) | NOT NULL |
| is_current | BOOLEAN | NOT NULL, DEFAULT false |
| description | TEXT | |
| bullet_points | JSONB | NOT NULL, DEFAULT '[]' |
| tech_tags | JSONB | NOT NULL, DEFAULT '[]' |
| logo_url | VARCHAR(500) | |
| gallery_images | JSONB | NOT NULL, DEFAULT '[]' |
| order_number | INTEGER | NOT NULL, DEFAULT 0 |
| is_published | BOOLEAN | NOT NULL, DEFAULT true |
| created_at | TIMESTAMPTZ | NOT NULL, DEFAULT NOW() |
| updated_at | TIMESTAMPTZ | NOT NULL, DEFAULT NOW() |

### projects

| Column | Type | Constraints |
|--------|------|-------------|
| id | UUID | PRIMARY KEY, DEFAULT gen_random_uuid() |
| title | VARCHAR(200) | NOT NULL |
| category | VARCHAR(50) | NOT NULL |
| description | TEXT | |
| tech_tags | JSONB | NOT NULL, DEFAULT '[]' |
| image_url | VARCHAR(500) | |
| demo_url | VARCHAR(500) | |
| repo_url | VARCHAR(500) | |
| is_featured | BOOLEAN | NOT NULL, DEFAULT false |
| order_number | INTEGER | NOT NULL, DEFAULT 0 |
| is_published | BOOLEAN | NOT NULL, DEFAULT true |
| created_at | TIMESTAMPTZ | NOT NULL, DEFAULT NOW() |
| updated_at | TIMESTAMPTZ | NOT NULL, DEFAULT NOW() |

### certifications

| Column | Type | Constraints |
|--------|------|-------------|
| id | UUID | PRIMARY KEY, DEFAULT gen_random_uuid() |
| title | VARCHAR(300) | NOT NULL |
| issuer | VARCHAR(200) | NOT NULL |
| issue_date | VARCHAR(50) | NOT NULL |
| expiry_date | VARCHAR(50) | |
| credential_id | VARCHAR(200) | |
| credential_url | VARCHAR(500) | |
| description | TEXT | |
| skills | JSONB | NOT NULL, DEFAULT '[]' |
| image_url | VARCHAR(500) | |
| category | VARCHAR(50) | NOT NULL, DEFAULT 'core' |
| status | VARCHAR(50) | NOT NULL, DEFAULT 'active' |
| order_number | INTEGER | NOT NULL, DEFAULT 0 |
| is_published | BOOLEAN | NOT NULL, DEFAULT true |
| created_at | TIMESTAMPTZ | NOT NULL, DEFAULT NOW() |
| updated_at | TIMESTAMPTZ | NOT NULL, DEFAULT NOW() |

### publications

| Column | Type | Constraints |
|--------|------|-------------|
| id | UUID | PRIMARY KEY, DEFAULT gen_random_uuid() |
| title | TEXT | NOT NULL |
| journal_name | VARCHAR(300) | NOT NULL |
| published_date | VARCHAR(50) | NOT NULL |
| status | VARCHAR(50) | NOT NULL, DEFAULT 'published' |
| authors | TEXT | NOT NULL |
| description | TEXT | |
| tags | JSONB | NOT NULL, DEFAULT '[]' |
| publication_url | VARCHAR(500) | |
| order_number | INTEGER | NOT NULL, DEFAULT 0 |
| is_published | BOOLEAN | NOT NULL, DEFAULT true |
| created_at | TIMESTAMPTZ | NOT NULL, DEFAULT NOW() |
| updated_at | TIMESTAMPTZ | NOT NULL, DEFAULT NOW() |

### contact_info

| Column | Type | Constraints |
|--------|------|-------------|
| id | UUID | PRIMARY KEY, DEFAULT gen_random_uuid() |
| email | VARCHAR(255) | |
| phone | VARCHAR(50) | |
| whatsapp_url | VARCHAR(500) | |
| github_url | VARCHAR(500) | |
| linkedin_url | VARCHAR(500) | |
| instagram_url | VARCHAR(500) | |
| location | VARCHAR(200) | |
| contact_description | TEXT | |
| is_published | BOOLEAN | NOT NULL, DEFAULT true |
| created_at | TIMESTAMPTZ | NOT NULL, DEFAULT NOW() |
| updated_at | TIMESTAMPTZ | NOT NULL, DEFAULT NOW() |

### contact_messages

| Column | Type | Constraints |
|--------|------|-------------|
| id | UUID | PRIMARY KEY, DEFAULT gen_random_uuid() |
| name | VARCHAR(100) | NOT NULL |
| email | VARCHAR(255) | NOT NULL |
| message | TEXT | NOT NULL |
| is_read | BOOLEAN | NOT NULL, DEFAULT false |
| created_at | TIMESTAMPTZ | NOT NULL, DEFAULT NOW() |

### site_settings

| Column | Type | Constraints |
|--------|------|-------------|
| id | UUID | PRIMARY KEY, DEFAULT gen_random_uuid() |
| site_title | VARCHAR(200) | NOT NULL |
| meta_description | TEXT | |
| favicon_url | VARCHAR(500) | |
| logo_url | VARCHAR(500) | |
| footer_text | VARCHAR(300) | |
| theme_mode | VARCHAR(20) | NOT NULL, DEFAULT 'dark' |
| maintenance_mode | BOOLEAN | NOT NULL, DEFAULT false |
| updated_at | TIMESTAMPTZ | NOT NULL, DEFAULT NOW() |

## Correctness Properties

*A property is a characteristic or behavior that should hold true across all valid executions of a system — essentially, a formal statement about what the system should do. Properties serve as the bridge between human-readable specifications and machine-verifiable correctness guarantees.*

### Property 1: Round-trip data integrity

*For any* valid content record created via the admin API (POST), retrieving that record via the corresponding public or admin GET endpoint should return all field values identical to what was submitted, without loss or corruption.

**Validates: Requirements 6.1, 11.2, 12.3**

### Property 2: Auth isolation

*For any* admin API endpoint and any request without a valid JWT token (missing, expired, or malformed), the backend should return a 401 Unauthorized response. Conversely, for any request with a valid non-expired JWT token, the backend should grant access and not return 401.

**Validates: Requirements 4.1, 4.2, 4.3, 4.4, 6.6**

### Property 3: Publish filtering

*For any* set of content records with mixed `is_published` values, the corresponding public GET endpoint should return only records where `is_published` is true, and the admin GET endpoint should return all records regardless of publish status.

**Validates: Requirements 5.2, 6.4**

### Property 4: Order consistency

*For any* reorder operation submitting a permutation of record IDs, the resulting `order_number` values should form a contiguous zero-based sequence matching the submitted array order, and subsequent public GET requests should return records sorted by that order.

**Validates: Requirements 6.5**

### Property 5: Seed idempotency

*For any* number of seed executions N ≥ 1 on the same database, the resulting database state (record count and content) should be identical to the state after a single execution — no duplicate records should be created.

**Validates: Requirements 3.4**

### Property 6: Password security

*For any* admin user stored in the database, the password_hash field should be a valid bcrypt hash with cost factor ≥ 10, and the GET /api/auth/me endpoint should never include the password_hash field in its response.

**Validates: Requirements 4.5, 4.6**

### Property 7: Input validation

*For any* create or update request to an admin CRUD endpoint with missing or invalid required fields, the backend should return a 400 Bad Request response with validation error details, and the database should remain unchanged.

**Validates: Requirements 5.3, 5.4, 6.2**

### Property 8: Configuration fail-fast

*For any* required environment variable (DATABASE_URL, JWT_SECRET, SERVER_PORT, CORS_ORIGIN), removing it at startup should cause the backend to log a descriptive error message and exit with a non-zero status code without accepting connections.

**Validates: Requirements 13.2**

## Error Handling

### Backend Error Responses

All API errors follow a consistent JSON structure:

```json
{
  "error": "Human-readable error message",
  "details": {}  // Optional: validation errors, field-level details
}
```

**HTTP Status Codes:**
- `400 Bad Request` — Invalid input, missing required fields, malformed JSON
- `401 Unauthorized` — Missing, expired, or invalid JWT token
- `404 Not Found` — Resource with given ID does not exist
- `500 Internal Server Error` — Unexpected server errors (logged, not exposed to client)

### Backend Error Handling Strategy

- **Repository layer**: Returns Go errors with context; does not handle HTTP concerns
- **Handler layer**: Catches repository errors, maps to appropriate HTTP status codes
- **Middleware**: JWT validation errors short-circuit with 401 before reaching handlers
- **Database errors**: Connection failures return 500; constraint violations return 400
- **Panic recovery**: Chi middleware recovers panics and returns 500

### Frontend Error Handling Strategy

- **API client**: Wraps fetch with error detection; throws typed errors for non-2xx responses
- **Page-level**: Each page catches API errors and displays contextual error messages
- **Auth errors**: 401 responses trigger automatic redirect to login page
- **Network errors**: Displayed as "Unable to connect to server" with retry option
- **Form validation**: Client-side validation prevents invalid submissions; server-side errors displayed inline

### Startup Errors

- Missing required environment variables: log descriptive message, exit with code 1
- Database connection failure: retry with backoff (3 attempts), then exit with code 1
- Migration failure: log which migration failed, exit with code 1

## Testing Strategy

### Dual Testing Approach

This feature uses both unit tests and property-based tests for comprehensive coverage.

**Unit tests** cover:
- Specific examples demonstrating correct behavior (e.g., login with known credentials)
- Integration points between components (e.g., handler → repository interaction)
- Edge cases and error conditions (e.g., empty payloads, max-length fields)
- Frontend component rendering and interaction

**Property-based tests** cover:
- Universal properties that hold across all valid inputs
- Comprehensive input coverage through randomization
- Round-trip integrity, auth isolation, publish filtering, ordering

### Property-Based Testing Configuration

- **Library**: Go — `pgregory.net/rapid` (fast, idiomatic Go PBT library)
- **Minimum iterations**: 100 per property test
- **Tag format**: `Feature: dynamic-portfolio-admin, Property {N}: {property_text}`

### Test Organization

```
api/
├── internal/
│   ├── handler/
│   │   ├── auth_test.go          # Auth handler unit + property tests
│   │   ├── public_test.go        # Public handler tests
│   │   └── admin_test.go         # Admin CRUD handler tests
│   ├── repository/
│   │   └── *_repo_test.go        # Repository integration tests
│   ├── service/
│   │   └── auth_service_test.go  # JWT/bcrypt property tests
│   └── database/
│       └── migrate_test.go       # Migration ordering tests
app/
├── src/
│   └── lib/
│       ├── api/
│       │   └── *.test.ts         # API client unit tests
│       ├── utils/
│       │   └── *.test.ts         # Validation/format unit tests
│       └── stores/
│           └── *.test.ts         # Store logic tests
```

### Key Test Scenarios

1. **Auth flow**: Valid login returns JWT; invalid login returns 401; expired token rejected
2. **CRUD lifecycle**: Create → Read → Update → Read → Delete → 404
3. **Publish filtering**: Mix of published/unpublished records; public API returns only published
4. **Reorder**: Submit permutation; verify order_number assignment and GET ordering
5. **Seed idempotency**: Run seed twice; verify no duplicates
6. **Input validation**: Missing required fields; invalid types; boundary values
7. **CORS**: Allowed origin succeeds; disallowed origin rejected
