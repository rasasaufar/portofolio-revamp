# Portfolio - Dynamic Portfolio System

A dynamic portfolio website with a Go REST API backend and SvelteKit frontend with ***REMOVED*** panel.

## Architecture

```
portfolio-svelte/
├── app/          # Frontend - SvelteKit + TypeScript
├── api/          # Backend - Go + PostgreSQL
└── docker-compose.yml
```

## Quick Start

### Prerequisites

- Node.js 18+
- Go 1.21+
- Docker & Docker Compose (for PostgreSQL)

### 1. Start Database

```bash
docker-compose up -d
```

### 2. Start Backend

```bash
cd api
cp .env.example .env
make run
```

API will be available at `http://localhost:8080`
Health check: `GET http://localhost:8080/api/health`

### 3. Start Frontend

```bash
cd app
cp .env.example .env
npm install
npm run dev
```

Frontend will be available at `http://localhost:5173`

## Development

- **Frontend docs**: See `app/README.md`
- **Backend docs**: See `api/README.md`
