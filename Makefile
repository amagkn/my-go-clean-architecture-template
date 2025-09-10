export

DB_MIGRATE_URL = postgres://postgres:postgres@localhost:5432/market?sslmode=disable
MIGRATE_PATH = ./migration/postgres

mod:
	go mod tidy

run: mod
	go run ./cmd/app

migrate-install:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.18.1

migrate-create:
	migrate create -ext sql -dir "$(MIGRATE_PATH)" $(name)

migrate-up:
	migrate -database "$(DB_MIGRATE_URL)" -path "$(MIGRATE_PATH)" up

migrate-down:
	migrate -database "$(DB_MIGRATE_URL)" -path "$(MIGRATE_PATH)" down -all

seed:
	psql $(DB_MIGRATE_URL) -f ./migration/postgres/seed.sql