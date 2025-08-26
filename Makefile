# Makefile для Banking API

# Переменные
BINARY_NAME=banking-api
BUILD_DIR=build
MAIN_FILE=main.go

# Go команды
GO=go
GOBUILD=$(GO) build
GOCLEAN=$(GO) clean
GOTEST=$(GO) test
GOGET=$(GO) get
GOMOD=$(GO) mod

# Docker команды
DOCKER=docker
DOCKER_BUILD=$(DOCKER) build
DOCKER_RUN=$(DOCKER) run
DOCKER_STOP=$(DOCKER) stop
DOCKER_RM=$(DOCKER) rm

# Docker Compose
DOCKER_COMPOSE=docker-compose

# Цвета для вывода
GREEN=\033[0;32m
YELLOW=\033[1;33m
RED=\033[0;31m
NC=\033[0m # No Color

.PHONY: all build clean test test-coverage run docker-build docker-run docker-stop docker-clean help

# Цель по умолчанию
all: clean build test

# Сборка приложения
build:
	@echo "$(GREEN)Building $(BINARY_NAME)...$(NC)"
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_FILE)
	@echo "$(GREEN)Build completed!$(NC)"

# Очистка
clean:
	@echo "$(YELLOW)Cleaning...$(NC)"
	$(GOCLEAN)
	@rm -rf $(BUILD_DIR)
	@echo "$(GREEN)Clean completed!$(NC)"

# Запуск тестов
test:
	@echo "$(GREEN)Running tests...$(NC)"
	$(GOTEST) -v ./...
	@echo "$(GREEN)Tests completed!$(NC)"

# Запуск тестов с покрытием
test-coverage:
	@echo "$(GREEN)Running tests with coverage...$(NC)"
	$(GOTEST) -coverprofile=coverage.out ./...
	$(GO) tool cover -html=coverage.out -o coverage.html
	@echo "$(GREEN)Coverage report generated: coverage.html$(NC)"

# Запуск приложения
run: build
	@echo "$(GREEN)Running $(BINARY_NAME)...$(NC)"
	@echo "$(YELLOW)Server will be available at http://localhost:8080$(NC)"
	@echo "$(YELLOW)Press Ctrl+C to stop$(NC)"
	./$(BUILD_DIR)/$(BINARY_NAME)

# Запуск приложения без сборки
run-dev:
	@echo "$(GREEN)Running $(BINARY_NAME) in development mode...$(NC)"
	@echo "$(YELLOW)Server will be available at http://localhost:8080$(NC)"
	@echo "$(YELLOW)Press Ctrl+C to stop$(NC)"
	$(GO) run $(MAIN_FILE)

# Установка зависимостей
deps:
	@echo "$(GREEN)Installing dependencies...$(NC)"
	$(GOMOD) download
	$(GOMOD) tidy
	@echo "$(GREEN)Dependencies installed!$(NC)"

# Проверка зависимостей
deps-check:
	@echo "$(GREEN)Checking dependencies...$(NC)"
	$(GOMOD) verify
	@echo "$(GREEN)Dependencies check completed!$(NC)"

# Сборка Docker образа
docker-build:
	@echo "$(GREEN)Building Docker image...$(NC)"
	$(DOCKER_BUILD) -t $(BINARY_NAME) .
	@echo "$(GREEN)Docker image built!$(NC)"

# Запуск Docker контейнера
docker-run: docker-build
	@echo "$(GREEN)Running Docker container...$(NC)"
	@echo "$(YELLOW)Server will be available at http://localhost:8080$(NC)"
	$(DOCKER_RUN) -d --name $(BINARY_NAME)-container -p 8080:8080 $(BINARY_NAME)
	@echo "$(GREEN)Container started!$(NC)"

# Остановка Docker контейнера
docker-stop:
	@echo "$(YELLOW)Stopping Docker container...$(NC)"
	$(DOCKER_STOP) $(BINARY_NAME)-container || true
	@echo "$(GREEN)Container stopped!$(NC)"

# Удаление Docker контейнера
docker-clean: docker-stop
	@echo "$(YELLOW)Removing Docker container...$(NC)"
	$(DOCKER_RM) $(BINARY_NAME)-container || true
	@echo "$(GREEN)Container removed!$(NC)"

# Запуск через Docker Compose
docker-compose-up:
	@echo "$(GREEN)Starting services with Docker Compose...$(NC)"
	$(DOCKER_COMPOSE) up -d
	@echo "$(GREEN)Services started!$(NC)"

# Остановка Docker Compose
docker-compose-down:
	@echo "$(YELLOW)Stopping services with Docker Compose...$(NC)"
	$(DOCKER_COMPOSE) down
	@echo "$(GREEN)Services stopped!$(NC)"

# Линтинг кода
lint:
	@echo "$(GREEN)Running linter...$(NC)"
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "$(YELLOW)golangci-lint not found, skipping...$(NC)"; \
	fi

# Форматирование кода
fmt:
	@echo "$(GREEN)Formatting code...$(NC)"
	$(GO) fmt ./...
	@echo "$(GREEN)Code formatted!$(NC)"

# Проверка безопасности зависимостей
security-check:
	@echo "$(GREEN)Checking dependencies for security vulnerabilities...$(NC)"
	@if command -v gosec >/dev/null 2>&1; then \
		gosec ./...; \
	else \
		echo "$(YELLOW)gosec not found, skipping...$(NC)"; \
	fi

# Генерация документации
docs:
	@echo "$(GREEN)Generating documentation...$(NC)"
	@if command -v godoc >/dev/null 2>&1; then \
		echo "$(YELLOW)Documentation will be available at http://localhost:6060$(NC)"; \
		godoc -http=:6060; \
	else \
		echo "$(YELLOW)godoc not found, skipping...$(NC)"; \
	fi

# Проверка всех проверок
check: fmt lint test security-check
	@echo "$(GREEN)All checks completed!$(NC)"

# Помощь
help:
	@echo "$(GREEN)Available commands:$(NC)"
	@echo "  $(YELLOW)build$(NC)          - Build the application"
	@echo "  $(YELLOW)clean$(NC)          - Clean build artifacts"
	@echo "  $(YELLOW)test$(NC)           - Run tests"
	@echo "  $(YELLOW)test-coverage$(NC)  - Run tests with coverage"
	@echo "  $(YELLOW)run$(NC)            - Build and run the application"
	@echo "  $(YELLOW)run-dev$(NC)        - Run in development mode"
	@echo "  $(YELLOW)deps$(NC)           - Install dependencies"
	@echo "  $(YELLOW)deps-check$(NC)     - Check dependencies"
	@echo "  $(YELLOW)docker-build$(NC)   - Build Docker image"
	@echo "  $(YELLOW)docker-run$(NC)     - Run Docker container"
	@echo "  $(YELLOW)docker-stop$(NC)    - Stop Docker container"
	@echo "  $(YELLOW)docker-clean$(NC)   - Remove Docker container"
	@echo "  $(YELLOW)docker-compose-up$(NC)   - Start with Docker Compose"
	@echo "  $(YELLOW)docker-compose-down$(NC) - Stop Docker Compose"
	@echo "  $(YELLOW)lint$(NC)           - Run linter"
	@echo "  $(YELLOW)fmt$(NC)            - Format code"
	@echo "  $(YELLOW)security-check$(NC) - Check for security vulnerabilities"
	@echo "  $(YELLOW)docs$(NC)           - Generate documentation"
	@echo "  $(YELLOW)check$(NC)          - Run all checks"
	@echo "  $(YELLOW)help$(NC)           - Show this help message"

# Информация о проекте
info:
	@echo "$(GREEN)Project Information:$(NC)"
	@echo "  Name: $(BINARY_NAME)"
	@echo "  Go version: $(shell $(GO) version)"
	@echo "  Build directory: $(BUILD_DIR)"
	@echo "  Main file: $(MAIN_FILE)"
	@echo "  Dependencies: $(shell $(GOMOD) list -m all | wc -l) packages"
