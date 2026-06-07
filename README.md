# 🎨 Portfolio Revamp

> A polished portfolio platform for showcasing projects, credentials, and professional milestones with a secure content dashboard.

![Svelte](https://img.shields.io/badge/Svelte-FF3E00?style=for-the-badge&logo=svelte&logoColor=white)
![TypeScript](https://img.shields.io/badge/TypeScript-3178C6?style=for-the-badge&logo=typescript&logoColor=white)
![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-4169E1?style=for-the-badge&logo=postgresql&logoColor=white)
![Nginx](https://img.shields.io/badge/Nginx-009639?style=for-the-badge&logo=nginx&logoColor=white)

**Portfolio Revamp** is a personal portfolio website and project showcase built with SvelteKit and Go, featuring a dynamic ***REMOVED*** dashboard for managing projects, certifications, work experience, and education records.

Production domain: [rasasaufar.site](https://rasasaufar.site), secured with Cloudflare SSL Full Strict.

## ✨ Features

**Project showcase**  
Display portfolio projects with clear descriptions, technologies used, images, and links. The showcase is designed to make each project easy to scan while still giving enough context for visitors to understand the work.

**Certification management**  
Add, edit, and delete professional certifications from the ***REMOVED*** dashboard. Certification records can be organized as part of the public portfolio to highlight verified learning, achievements, and industry credentials.

**Work experience timeline**  
Manage work history entries in a structured timeline format. Each entry can capture role details, company information, dates, descriptions, and supporting highlights.

**Education records**  
Track educational background in a dedicated content area. Education entries help present academic history alongside project work and professional experience.

**Admin dashboard**  
Secure JWT-authenticated ***REMOVED*** panel for CRUD operations across portfolio content. The dashboard centralizes content management for projects, certifications, work experience, education records, and other site data.

**File uploads**  
Support for uploading project images, avatars, and other visual assets used across the portfolio. Uploaded files can be managed through backend services and displayed on the frontend.

**Responsive design**  
Mobile-first interface with a clean UI built for readability and usability across devices. The public site and dashboard are designed to remain comfortable on both desktop and smaller screens.

## 🧰 Tech Stack

| Layer | Technology |
|---|---|
| Frontend | SvelteKit, Svelte 5, adapter-node, TypeScript |
| Backend | Go, chi router, pgx/v5, JWT authentication |
| Database | PostgreSQL 16 Alpine |
| Deployment | Docker Compose, Nginx reverse proxy, Let's Encrypt SSL, GitHub Actions CI/CD |
| Domain & SSL | rasasaufar.site, Cloudflare SSL Full Strict |

## 🧱 Docker Architecture

| Service | Role |
|---|---|
| Frontend | SvelteKit app served via Nginx reverse proxy |
| API | Go backend served via Nginx reverse proxy |
| PostgreSQL | Database (internal only) |
| pgAdmin | Database management UI (internal only) |

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
Frontend             API
                       |
                       v
                  PostgreSQL
```

## 📋 Prerequisites

- Node.js
- Go
- Docker and Docker Compose
- Nginx for production reverse proxy
- Let's Encrypt for SSL certificates
- Cloudflare configured for `rasasaufar.site`

## 🚀 Installation & Setup

### 1. Clone the repository

```bash
git clone <repository-url>
cd portofolio-revamp
```

### 2. Configure environment files

Copy the example environment files:

```bash
cp api/.env.example api/.env
cp app/.env.example app/.env
cp .env.example .env
```

Configure your environment variables in `.env` (see `.env.example` for reference).

Required environment variables include:

```text
DATABASE_URL
JWT_SECRET
ADMIN_PASSWORD
CORS_ORIGIN
PUBLIC_API_BASE_URL
```

### 3. Start with Docker Compose

```bash
docker compose up --build
```

### 4. Run locally for development

Install and run the frontend:

```bash
cd app
npm install
npm run dev
```

Run the backend:

```bash
cd api
go mod download
go run ./cmd/server
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

## 🔌 API Overview

The API provides RESTful endpoints for managing projects, certifications, work experience, and education records. It also supports authentication, dashboard content management, file uploads, and public portfolio data delivery.

Full API documentation is available in `docs/API.md`.

## 🖼️ Screenshots

Screenshots can be added here as the interface evolves.

| Page | Preview |
|---|---|
| Homepage | `docs/screenshots/homepage.png` |
| Project Showcase | `docs/screenshots/projects.png` |
| Admin Dashboard | `docs/screenshots/***REMOVED***-dashboard.png` |
| Content Editor | `docs/screenshots/content-editor.png` |

## 🤝 Contributing

Contributions are welcome. Please keep changes focused, documented, and easy to review.

```bash
git checkout -b feature/your-feature-name
```

Before opening a pull request, run the relevant checks for the frontend and backend:

```bash
cd app
npm run check
npm run build

cd ../api
go vet ./...
go test ./...
```

## 📄 License

This project is licensed under the MIT License.
