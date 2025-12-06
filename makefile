.PHONY: create-migration

MIGRATIONS_DIR="./db/migrations"
DB_DSN="postgres://dev:dev@localhost:5432/dev"

#---------MIGRATION---------

create-migration:
	@echo "Введите имя миграции:"
	@read MIGRATION_NAME; \
	goose -dir $(MIGRATIONS_DIR) create $$MIGRATION_NAME sql

migrations-up:
	goose -dir $(MIGRATIONS_DIR) postgres $(DB_DSN) up

migrations-down:
	goose -dir $(MIGRATIONS_DIR) postgres $(DB_DSN) down

#---------DOCKER---------

compose-build:
	docker-compose build --no-cache

compose-up:
	docker-compose up -d
	@sleep 2

compose-down:
	docker-compose down -v
	@sleep 2

#---------SERVICE---------

start: compose-build compose-up migrations-up

stop: migrations-down compose-down

local-start:
	go run ./cmd/main.go

#---------TESTING---------

test:
	go test ./internal/service -v
	go test ./internal/handler -v

test-cover:
	go test ./internal/service -cover
	go test ./internal/handler -cover

#---------FORMATTING---------

fmt:
	go fmt ./cmd/
	go fmt ./internal/database/
	go fmt ./internal/errors/
	go fmt ./internal/handler/
	go fmt ./internal/models/
	go fmt ./internal/repository/
	go fmt ./internal/server/
	go fmt ./internal/service/