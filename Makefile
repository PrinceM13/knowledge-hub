COMPOSE_FILE=docker/docker-compose.yml
ENV_FILE=docker/.env

include $(ENV_FILE)
export

# Construct DB URL for migrate
MIGRATE_DB=postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable
MIGRATIONS_PATH=apps/api-go/migrations

.PHONY: \
	db-up db-down db-restart db-logs db-psql \
	migrate-up migrate-down migrate-force

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
