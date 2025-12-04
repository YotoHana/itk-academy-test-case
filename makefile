.PHONY: create-migration

MIGRATIONS_DIR="./db/migrations"
DB_DSN="postgresql://dev:dev@localhost:5432/dev"

create-migration:
	@echo "Введите имя миграции:"
	@read MIGRATION_NAME; \
	goose -dir $(MIGRATIONS_DIR) create $$MIGRATION_NAME sql

migrations-up:
	goose -dir $(MIGRATIONS_DIR) postgresql $(DB_DSN) up

migrations-down:
	goose -dir $(MIGRATIONS_DIR) postgresql $(DB_DSN) down
