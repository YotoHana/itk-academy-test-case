.PHONY: create-migration

MIGRATIONS_DIR="./db/migrations"
DB_DSN="postgres://dev:dev@localhost:5432/dev"

create-migration:
	@echo "Введите имя миграции:"
	@read MIGRATION_NAME; \
	goose -dir $(MIGRATIONS_DIR) create $$MIGRATION_NAME sql

migrations-up:
	goose -dir $(MIGRATIONS_DIR) postgres $(DB_DSN) up

migrations-down:
	goose -dir $(MIGRATIONS_DIR) postgres $(DB_DSN) down

compose-build:
	docker-compose build --no-cache

compose-up:
	docker-compose up -d
	@sleep 2

compose-down:
	docker-compose down -v
	@sleep 2

db-up: compose-build compose-up migrations-up

db-down: migrations-down compose-down

start:
	go run ./cmd/main.go