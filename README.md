# FAQ API

A REST API for managing FAQs, categories, translations, and merchant stores. Built with Gin, Gorm, and PostgreSQL; migrations via Goose.

## Stack

- Go 1.25.3 and Gin
- Gorm with PostgreSQL
- Goose for database migrations
- Docker and docker-compose for local infrastructure

## Project Structure

```
cmd/api/main.go              # Application entrypoint and router setup
internal/
    ├── config/                # Environment configuration loader
    ├── db/                     # Database initialization
    ├── handlers/               # HTTP request handlers
    ├── services/               # Business logic layer
    ├── middlewares/            # Authentication and authorization middleware
    ├── models/                 # Gorm data models
    ├── migrations/             # Goose database migrations
    ├── helpers/                # Utility functions
    ├── types/                  # Shared types and constants
    ├── dtos/                   # Data transfer objects
    ├── requests/               # Request structures
    ├── responses/              # Response structures
    ├── errors/                 # Custom error types
    └── logger/                 # Logging configuration
```

## Prerequisites

- Go 1.25.3
- Docker and docker-compose

## Quick Start

### Clone and Setup

```bash
git clone https://github.com/kareemhamed001/FAQ-MS.git
cd FAQ-MS
go mod vendor && go mod tidy
```

### Environment Variables

```env
APP_PORT=8080
APP_ENV=local                # local or production
DB_DRIVER=postgres
DB_HOST=localhost            # Use 'postgres' with docker-compose
DB_PORT=5432
DB_USER=admin
DB_PASSWORD=admin
DB_NAME=faq_db
JWT_PRIVATE_KEY=your-secret-key

GOOSE_DRIVER=postgres
GOOSE_DBSTRING=postgres://admin:admin@localhost:5433/faq_db?sslmode=disable
GOOSE_MIGRATION_DIR=./internal/migrations
```

### Running with Docker

```bash
docker-compose up -d --build
```

- API: http://localhost:8080
- PostgreSQL: localhost:5433
- pgAdmin: http://localhost:5050 (admin@example.com / admin123)

### Running Locally

1. Start PostgreSQL and configure environment variables
2. Install dependencies: `go mod vendor && go mod tidy`
3. Run migrations: `goose up`
4. Seed database (optional): `make seed-admin` or `make seed-admin-custom NAME="Admin" EMAIL="admin@email.com" PASSWORD="password"`
5. Start server: `go run ./cmd/api/main.go`

## Database Seeding

The project includes a database seeder to initialize admin users:

### Quick Setup with Make

```bash
# Start database
make db-up

# Seed default admin user (admin@example.com / admin123)
make seed-admin

# Seed custom admin user
make seed-admin-custom NAME="John Doe" EMAIL="john@example.com" PASSWORD="secret123"

# Stop database
make db-down
```

### Direct CLI Commands

```bash
# Seed default admin user
DB_HOST=localhost go run cmd/seeder/main.go seed:admin -default

# Seed custom admin user
DB_HOST=localhost go run cmd/seeder/main.go seed:admin -name "Admin Name" -email "admin@email.com" -password "password"
```

For detailed seeding documentation, see [SEEDER.md](SEEDER.md).

## Migrations

Goose manages all database schema changes:

- Apply: `goose up`
- Rollback: `goose down`
- Location: `internal/migrations/`

## Authentication

- Endpoints: `POST /auth/register`, `POST /auth/login`
- Protected routes require: `Authorization: Bearer <token>`
- Roles: `admin`, `merchant`, `customer`

## API Endpoints

| Endpoint              | Method | Access         | Description           |
| --------------------- | ------ | -------------- | --------------------- |
| `/health`             | GET    | Public         | Health check          |
| `/auth/register`      | POST   | Public         | User registration     |
| `/auth/login`         | POST   | Public         | User login            |
| `/api/faq-categories` | All    | Admin          | Manage FAQ categories |
| `/api/faqs`           | All    | Admin/Merchant | Manage FAQs           |
| `/api/stores`         | GET    | Public         | List stores           |
| `/api/stores/:id`     | GET    | Public         | Get store details     |

## Key Assumptions

- Each merchant owns exactly one store
- Merchants can view all their FAQs and global FAQs
- Admins can edit merchant FAQs
- Users see FAQs in their preferred language only
- Repository pattern not required for this project scope

## Development Notes

- Use `APP_ENV=local` for debug mode; production uses release mode
- Keep JWT secret confidential in shared environments
- Create forward-only migrations; avoid editing existing migrations

## Testing

Use the Postman collection in the root directory to test all endpoints.
