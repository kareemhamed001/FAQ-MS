# FAQ API

A REST API for managing FAQs, categories, translations, and merchant stores. Built with Gin, Gorm, and PostgreSQL; migrations via Goose.

## Stack

- Go (1.25.3) and Gin
- Gorm with Postgres
- Goose for migrations
- Docker + docker-compose for local infra

## Project Layout

- cmd/api/main.go – entrypoint and router wiring
- internal/config – env config loader
- internal/DB – DB init
- internal/handlers – HTTP handlers
- internal/services – business logic
- internal/middlewares – auth/role middleware
- internal/models – Gorm models
- internal/migrations – Goose migrations
- internal/helpers – utility functions
- internal/types – shared types and constants
- internal/DTOs – request/response structs
- internal/requests – requests structs
- internal/responses – response structs
- internal/errors – custom error types
- internal/logger – logging setup

## Prerequisites

- Go 1.25.3
- Docker + docker-compose

## Installation

```
git clone https://github.com/kareemhamed001/FAQ-MS.git
```

```
go mod vendor
```

```
go mod tidy
```

## Environment Variables

```
APP_PORT=8080
APP_ENV=local           # local|production
DB_DRIVER=postgres
DB_HOST=localhost       # or postgres when using docker-compose
DB_PORT=5432            # 5433 exposed on host via compose
DB_USER=admin
DB_PASSWORD=admin
DB_NAME=faq_db
JWT_PRIVATE_KEY=change-me

# Goose (used by provided commands)
GOOSE_DRIVER=postgres
GOOSE_DBSTRING=postgres://admin:admin@localhost:5433/faq_db?sslmode=disable
GOOSE_MIGRATION_DIR=./internal/migrations
```

## Running with Docker

```
docker-compose up -d --build
```

- App: http://localhost:8080
- Postgres: localhost:5433
- pgAdmin: http://localhost:5050 (admin@example.com / admin123)

## Running Locally (no Docker for app)

1. Start Postgres (from compose or your own instance) and ensure DB vars match.
2. Install deps: `go mod vendor && go mod tidy`
3. Run migrations: `GOOSE_DRIVER=postgres GOOSE_DBSTRING="postgres://admin:admin@localhost:5433/faq_db?sslmode=disable" GOOSE_MIGRATION_DIR=./internal/migrations goose up`
4. Start API: `go run ./cmd/api/main.go`

## Migrations

- Tool: Goose
- Directory: internal/migrations
- Apply: `goose up`
- Rollback last: `goose down`

## Authentication & Roles

- Auth endpoints: `/auth/register`, `/auth/login`
- Bearer token required for protected routes: `Authorization: Bearer <token>`
- Roles: admin, merchant, customer

## API Surface (high level)

- Health: `GET /health`
- Auth: `POST /auth/register`, `POST /auth/login`
- FAQ Categories (admin only): CRUD under `/api/faq-categories`
- FAQs (admin/merchant): CRUD under `/api/faqs`
- Stores: `GET /api/stores` and `GET /api/stores/:id` (public)

## Development Tips

- Use `APP_ENV=local` to keep Gin in debug mode; production sets release mode.
- Keep JWT secret non-default in any shared environment.
- When changing migrations, reset DB or create forward-only migrations instead of editing existing ones.

## Assumptions

-merchant can have only one store.
-merchant can list all faqs and global ones.
-admin can edit merchants faqs.
-user see faqs with its local language only.
