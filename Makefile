COMPOSE_FILE=docker/docker-compose.yml
ENV_FILE=docker/.env

include $(ENV_FILE)
export

# Construct DB URL for migrate
MIGRATE_DB=postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable
MIGRATIONS_PATH=apps/api-go/migrations

.PHONY: \
	db-up db-down db-restart db-logs db-psql \
	migrate-up migrate-down migrate-force \
	test test-unit test-integration test-unit-verbose test-coverage \
	install dev dev-next build build-next lint lint-next clean

# =========================
# Database (Docker)
# =========================

db-up:
	docker compose -f $(COMPOSE_FILE) up -d

db-down:
	docker compose -f $(COMPOSE_FILE) down

db-restart:
	docker compose -f $(COMPOSE_FILE) down
	docker compose -f $(COMPOSE_FILE) up -d

db-logs:
	docker compose -f $(COMPOSE_FILE) logs -f postgres

db-psql:
	docker exec -it knowledge_hub_postgres \
	psql -U $(POSTGRES_USER) -d $(POSTGRES_DB)

# =========================
# Migrations
# =========================

migrate-up:
	migrate -path $(MIGRATIONS_PATH) -database "$(MIGRATE_DB)" up

migrate-down:
	migrate -path $(MIGRATIONS_PATH) -database "$(MIGRATE_DB)" down 1

migrate-force:
	migrate -path $(MIGRATIONS_PATH) -database "$(MIGRATE_DB)" force $(version)

migrate-create:
	migrate create -ext sql -dir $(MIGRATIONS_PATH) $(name)

# =========================
# Testing
# =========================

test:
	cd apps/api-go && go test ./...

test-unit:
	cd apps/api-go && go test -short ./...

test-integration:
	cd apps/api-go && go test -run Integration ./...

test-unit-verbose:
	cd apps/api-go && go test -v -short ./...

test-integration-verbose:
	cd apps/api-go && go test -v -run Integration ./...

test-coverage:
	cd apps/api-go && go test -cover ./... && \
	go test -coverprofile=coverage.out ./... && \
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: apps/api-go/coverage.html"

# =========================
# Node / pnpm
# =========================

install:
	pnpm install

# =========================
# Next.js
# =========================

dev:
	pnpm --filter web-next dev

dev-next:
	pnpm --filter web-next dev

build-next:
	pnpm --filter web-next build

lint-next:
	pnpm --filter web-next lint

clean:
	rm -rf node_modules
	rm -rf apps/web-next/node_modules
	rm -rf apps/web-next/.next
	rm -rf pnpm-lock.yaml
