# Makefile for building Go applications

# Variables
APP_NAME = auth-service
VERSION = 1.0.0
BUILD_DIR = bin
OS = $(shell go env GOOS)
ARCH = $(shell go env GOARCH)

# Build the application
build:
	@echo "Building..."
	go build -o bin/main.exe cmd/main.go

# Run the application
run:
	go run cmd/main.go

# Start docker compose development environment
start-dev:
	docker compose -f docker-compose-dev.yml up -d

# Stop docker compose development environment
stop-dev:
	docker compose -f docker-compose-dev.yml down

# Test the application
test:
	@echo "Testing..."
	@go test ./... -v

docs:
	@echo "Generating docs..."
	@echo "Generating docs..."
	@powershell -ExecutionPolicy Bypass -Command "if (Get-Command swag -ErrorAction SilentlyContinue) { \
		swag init -g .\cmd\main.go -o ./docs; \
		Write-Output 'Watching...'; \
	} else { \
		Write-Output 'Installing air...'; \
		go install github.com/swaggo/swag/cmd/swag@latest; \
		swag init -g .\cmd\main.go -o ./docs; \
		Write-Output 'Watching...'; \
	}"

watch:

	@powershell -ExecutionPolicy Bypass -Command "if (Get-Command air -ErrorAction SilentlyContinue) { \
		air; \
		Write-Output 'Watching...'; \
	} else { \
		Write-Output 'Installing air...'; \
		go install github.com/air-verse/air@latest; \
		air; \
		Write-Output 'Watching...'; \
	}"

# Clean build artifacts
clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)

# Targets
.PHONY: all clean build

all: build