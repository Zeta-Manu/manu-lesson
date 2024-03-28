# Makefile for API

# parameters
PORT = 8080
DOCKER_IMAGE_NAME = manu-lesson
DOCKER_CONTAINER_NAME = manu-lesson-container
DB_PORT = 3306
DB_NAME = lesson
DB_URL = "mysql://root:password@tcp(localhost:$(DB_PORT))/$(DB_NAME)"
MIGRATIONS_PATH = db/migrations

tidy:
	go mod tidy

vendor:
	go mod vendor

migrate:
	@echo "Please specify 'up' or 'down' as a sub-target"
	@echo $(DB_URL)

migrate-install:
	go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

migrate-up:
	@echo "Running migrations up..."
	migrate -database $(DB_URL) -path $(MIGRATIONS_PATH) up

migrate-down:
	@echo "Running migrations down..."
	migrate -database $(DB_URL) -path $(MIGRATIONS_PATH) down

build-docker:
	docker compose build

run-docker:
	docker compose up

stop-docker:
	docker compose down

rebuild-docker: stop-docker build-docker run-docker

swag: 
	swag init -g cmd/main.go 

.PHONY: migrate migrate-up migrate-down migrate-install build-docker run-docker stop-docker rebuild-docker swag
