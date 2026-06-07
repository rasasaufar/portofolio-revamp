# 🎨 Portfolio Revamp

> A modern personal portfolio and content dashboard for showcasing projects, credentials, and professional milestones.

![Svelte](https://img.shields.io/badge/Svelte-FF3E00?style=for-the-badge&logo=svelte&logoColor=white)
![TypeScript](https://img.shields.io/badge/TypeScript-3178C6?style=for-the-badge&logo=typescript&logoColor=white)
![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-4169E1?style=for-the-badge&logo=postgresql&logoColor=white)
![Nginx](https://img.shields.io/badge/Nginx-009639?style=for-the-badge&logo=nginx&logoColor=white)

**Portfolio Revamp** is a personal portfolio website and project showcase built with SvelteKit and Go, featuring a dynamic ***REMOVED*** dashboard for managing projects, certifications, work experience, and education records.

Production domain: [rasasaufar.site](https://rasasaufar.site) with Cloudflare SSL Full Strict.

## ✨ Features

**Project showcase**  
Display portfolio projects with rich descriptions, technology tags, preview images, live demo links, and repository links. Featured projects can be highlighted and reordered from the ***REMOVED*** dashboard.

**Certification management**  
Add, edit, delete, publish, and organize professional certifications. Each certification can include issuer details, credential IDs, credential URLs, skill tags, status, and supporting images.

**Work experience timeline**  
Manage work history entries as a clean professional timeline. Each experience supports roles, company details, start and end dates, current-position flags, bullet points, tech tags, logos, and gallery images.

**Education records**  
Track educational background with institution data, degree information, major, year range, GPA, descriptions, tags, and image assets.

**Admin dashboard**  
Secure JWT-authenticated ***REMOVED*** panel for CRUD operations across portfolio content, including identity, capabilities, strengths, dossier, education, experiences, projects, certifications, publications, contact information, messages, and site settings.

**File uploads**  
Upload and serve project images, avatars, logos, certification images, and other visual assets used throughout the portfolio.

**Responsive design**  
Mobile-first SvelteKit frontend with a clean UI, responsive layouts, and polished ***REMOVED*** workflows for managing content on any screen size.

## 🧰 Tech Stack

| Layer | Technology |
|---|---|
| Frontend | SvelteKit, Svelte 5, TypeScript, adapter-node |
| Backend | Go, chi router, pgx/v5, JWT authentication |
| Database | PostgreSQL 16 Alpine |
| Deployment | Docker Compose, Nginx reverse proxy, Let's Encrypt SSL |
| CI/CD | GitHub Actions |
| Domain & SSL | rasasaufar.site, Cloudflare SSL Full Strict |

## 🧱 Docker Architecture

| Service | Role |
|---|---|
| Frontend | SvelteKit app served via Nginx reverse proxy |
| API | Go backend served via Nginx reverse proxy |
| PostgreSQL | Database used by the API, available internally to the Docker network |
| pgAdmin | Database management UI, intended for internal ***REMOVED***istration only |

```text
Client
  |
  v
Cloudflare SSL Full Strict
  |
  v
Nginx Reverse Proxy
  |--------------------|
  v                    v
SvelteKit Frontend   Go API
                       |
                       v
                  PostgreSQL
```

## 📋 Prerequisites

- Node.js 18 or newer
- Go 1.26 or compatible with the version declared in `api/go.mod`
- Docker and Docker Compose
- PostgreSQL client tools, optional for manual database inspection
- Nginx and Let's Encrypt tooling for production deployment
- Cloudflare account configured for `rasasaufar.site`

## 🚀 Installation & Setup

### 1. Clone the repository

```bash
git clone https://github.com/your-username/portofolio-revamp.git
cd portofolio-revamp
```

### 2. Configure environment variables

Create environment files for the API, frontend, and Docker Compose deployment.

```bash
cp api/.env.example api/.env
cp app/.env.example app/.env
touch .env
```

Recommended production-oriented variables:

```env
# Backend
DATABASE_URL=postgres://portfolio:change_me@postgres/portfolio_db?sslmode=disable
JWT_SECRET=replace-with-a-long-random-secret
ADMIN_PASSWORD=replace-with-a-secure-***REMOVED***-password
ADMIN_DEFAULT_PASSWORD=replace-with-a-secure-***REMOVED***-password
CORS_ORIGIN=https://rasasaufar.site

# Frontend
PUBLIC_API_BASE_URL=https://rasasaufar.site
VITE_API_BASE_URL=https://rasasaufar.site
```

`ADMIN_DEFAULT_PASSWORD` and `VITE_API_BASE_URL` are used by the current codebase. `ADMIN_PASSWORD` and `PUBLIC_API_BASE_URL` are documented as deployment-facing aliases when standardizing environment naming.

### 3. Start with Docker Compose

```bash
docker compose up -d --build
```

Run migrations and seed the first ***REMOVED*** user when needed:

```bash
docker compose exec api ./server -migrate
docker compose exec api ./server -seed
```

### 4. Run locally for development

Start the backend:

```bash
cd api
cp .env.example .env
make migrate
make seed
make run
```

Start the frontend in another terminal:

```bash
cd app
cp .env.example .env
npm install
npm run dev
```

### 5. Verify the API

```bash
curl https://rasasaufar.site/api/health
```

Example login request:

```bash
curl -X POST https://rasasaufar.site/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "***REMOVED***",
    "password": "your-***REMOVED***-password"
  }'
```

Example authenticated project creation:

```bash
curl -X POST https://rasasaufar.site/api/***REMOVED***/projects \
  -H "Authorization: Bearer <jwt-token>" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Portfolio Revamp",
    "category": "Web Application",
    "description": "A SvelteKit and Go portfolio platform.",
    "tech_tags": ["SvelteKit", "Go", "PostgreSQL"],
    "demo_url": "https://rasasaufar.site",
    "is_featured": true,
    "order_number": 1,
    "is_published": true
  }'
```

## 📁 Project Structure

```text
portofolio-revamp/
├── .github/workflows/deploy.yml
├── api/
│   ├── cmd/server/
│   ├── internal/
│   ├── migrations/
│   ├── seeds/
│   ├── uploads/
│   ├── Dockerfile
│   └── go.mod
├── app/
│   ├── src/lib/
│   ├── src/routes/
│   ├── static/
│   ├── Dockerfile
│   └── package.json
├── docker-compose.yml
└── README.md
```

## 🔌 API Endpoints

### Public

| Method | Endpoint | Description |
|---|---|---|
| GET | `/api/health` | Check API and database health |
| POST | `/api/auth/login` | Authenticate ***REMOVED*** user and return a JWT |
| GET | `/api/identity` | Get published identity and hero profile data |
| GET | `/api/capabilities` | Get published capability snapshots |
| GET | `/api/strengths` | Get published implementation strengths |
| GET | `/api/dossier` | Get published professional dossier content |
| GET | `/api/education` | Get published education records |
| GET | `/api/experiences` | Get published work experience entries |
| GET | `/api/projects` | Get published portfolio projects |
| GET | `/api/certifications` | Get published certifications |
| GET | `/api/publications` | Get published publications |
| GET | `/api/contact-info` | Get published contact information |
| GET | `/api/site-settings` | Get public site settings |
| POST | `/api/contact/messages` | Submit a contact message |
| GET | `/uploads/*` | Serve uploaded public assets |

### Authenticated

| Method | Endpoint | Description |
|---|---|---|
| GET | `/api/auth/me` | Get the currently authenticated ***REMOVED*** profile |
| PUT | `/api/auth/password` | Change the authenticated ***REMOVED*** password |
| POST | `/api/upload` | Upload images or files for portfolio content |

### Admin

Admin endpoints require a valid JWT in the `Authorization: Bearer <token>` header.

| Method | Endpoint | Description |
|---|---|---|
| GET | `/api/***REMOVED***/dashboard/stats` | Get dashboard statistics and unread message count |
| GET | `/api/***REMOVED***/{resource}` | List records for an ***REMOVED*** resource |
| POST | `/api/***REMOVED***/{resource}` | Create a new resource record |
| PUT | `/api/***REMOVED***/{resource}/reorder` | Reorder records for sortable resources |
| GET | `/api/***REMOVED***/{resource}/{id}` | Get a single resource record by ID |
| PUT | `/api/***REMOVED***/{resource}/{id}` | Update a resource record by ID |
| DELETE | `/api/***REMOVED***/{resource}/{id}` | Delete a resource record by ID |
| PATCH | `/api/***REMOVED***/{resource}/{id}/publish` | Publish a resource record |
| PATCH | `/api/***REMOVED***/{resource}/{id}/unpublish` | Unpublish a resource record |
| PATCH | `/api/***REMOVED***/messages/{id}/read` | Mark a contact message as read |

Supported ***REMOVED*** resources:

```text
identity
capabilities
strengths
dossier
education
experiences
projects
certifications
publications
contact
messages
settings
```

## 🖼️ Screenshots

Screenshots can be added here as the UI stabilizes.

| Page | Preview |
|---|---|
| Homepage | `docs/screenshots/homepage.png` |
| Project Showcase | `docs/screenshots/projects.png` |
| Admin Dashboard | `docs/screenshots/***REMOVED***-dashboard.png` |
| Content Editor | `docs/screenshots/content-editor.png` |

## 🤝 Contributing

Contributions are welcome. For a clean workflow:

1. Fork the repository.
2. Create a feature branch.
3. Make focused changes with clear commit messages.
4. Run frontend and backend checks before opening a pull request.
5. Open a pull request with a concise description, screenshots for UI changes, and notes for any migration or deployment impact.

Useful checks:

```bash
cd app
npm run check
npm run build

cd ../api
go vet ./...
go test ./...
```

## 📄 License

This project is licensed under the MIT License. See `LICENSE` for details.
