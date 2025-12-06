.PHONY: create-migration

MIGRATIONS_DIR="./db/migrations"
DB_DSN="postgres://dev:dev@localhost:5432/dev"
RPS_SCRIPT="./post.lua"
WALLET_UUID="73fd60a4-ef54-484c-9de8-4beb3808da26"

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

test-rps:
	wrk -t12 -c400 -d30s -s $(RPS_SCRIPT) http://localhost:8080/api/v1/wallet
	wrk -t12 -c400 -d30s http://localhost:8080/api/v1/wallets/$(WALLET_UUID)

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