include .env
export

MIGRATIONS_DIR=internal/pkg/db/migrations
POSTGRES_CONNECT_STRING="host=localhost user=test password=test dbname=urls_db sslmode=disable"

.PHONY: all
all: goose-up
	go run cmd/tinyurl/main.go

.PHONY: compose-up
compose-up:
	docker-compose -p db -f ./build/docker-compose.yml up -d

.PHONY: compose-rm
compose-rm:
	docker-compose -p db -f ./build/docker-compose.yml rm -fvs

.PHONY: goose-status
goose-status:
	goose -dir $(MIGRATIONS_DIR) postgres $(POSTGRES_CONNECT_STRING) status

.PHONY: goose-up
goose-up:
	goose -dir $(MIGRATIONS_DIR) postgres $(POSTGRES_CONNECT_STRING) up

.PHONY: goose-down
goose-down:
	goose -dir $(MIGRATIONS_DIR) postgres $(POSTGRES_CONNECT_STRING) down

.PHONY: proto-generate
proto-generate:
	buf mod update
	buf generate