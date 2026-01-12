COMPOSE_FILE=docker/docker-compose.yml
ENV_FILE=docker/.env

include $(ENV_FILE)
export

.PHONY: db-up db-down db-restart db-logs db-psql

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
