include .env
export

# Database Migration Variables
MIGRATE=$(shell go env GOPATH)/bin/migrate
MIGRATION_DIR=db/migrations

.PHONY: run build test migrate-up migrate-down migrate-create clean

run:
	@echo "Starting Gophermart Server..."
	go run cmd/api/main.go cmd/api/router.go

build:
	@echo "Building Gophermart Server..."
	go build -o bin/api cmd/api/main.go cmd/api/router.go

test:
	@echo "Running Tests..."
	go test -v ./... -cover

migrate-up:
	@echo "Running up migrations..."
	$(MIGRATE) -path $(MIGRATION_DIR) -database "$(DB_CONN_STR)" -verbose up

migrate-down:
	@echo "Running down migrations..."
	$(MIGRATE) -path $(MIGRATION_DIR) -database "$(DB_CONN_STR)" -verbose down 1

migrate-create:
	@read -p "Enter migration name: " name; \
	$(MIGRATE) create -ext sql -dir $(MIGRATION_DIR) -seq $$name

clean:
	@echo "Cleaning up..."
	rm -rf bin/
