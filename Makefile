.PHONY: help
help:
	@echo "Available commands:"
	@echo "  make db-up              - Start PostgreSQL database using docker-compose"
	@echo "  make db-down            - Stop PostgreSQL database"
	@echo "  make seed-admin          - Seed default admin user (admin@example.com / admin123)"
	@echo "  make seed-admin-custom   - Seed custom admin user (use NAME=, EMAIL=, PASSWORD=)"
	@echo "  make run                 - Run the API server"
	@echo "  make dev                 - Run in development mode with hot reload"
	@echo ""
	@echo "Examples:"
	@echo "  make seed-admin"
	@echo "  make seed-admin-custom NAME='John Doe' EMAIL='john@example.com' PASSWORD='secret123'"

.PHONY: db-up
db-up:
	docker-compose up -d
	@echo "Database started"

.PHONY: db-down
db-down:
	docker-compose down
	@echo "Database stopped"

.PHONY: seed-admin
seed-admin:
	@echo "Seeding default admin user..."
	DB_HOST=localhost go run cmd/seeder/main.go seed:admin -default

.PHONY: seed-admin-custom
seed-admin-custom:
	@if [ -z "$(NAME)" ] || [ -z "$(EMAIL)" ] || [ -z "$(PASSWORD)" ]; then \
		echo "Error: NAME, EMAIL, and PASSWORD are required"; \
		echo "Usage: make seed-admin-custom NAME='John Doe' EMAIL='john@example.com' PASSWORD='secret123'"; \
		exit 1; \
	fi
	@echo "Seeding custom admin user..."
	DB_HOST=localhost go run cmd/seeder/main.go seed:admin -name "$(NAME)" -email "$(EMAIL)" -password "$(PASSWORD)"

.PHONY: run
run:
	go run cmd/api/main.go

.PHONY: build
build:
	go build -o bin/api cmd/api/main.go
	@echo "Built binary: bin/api"


.PHONY: tidy
tidy:
	go mod tidy

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: lint
lint:
	golint ./...
