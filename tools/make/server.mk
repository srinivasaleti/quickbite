# server.mk

SERVER_PORT = 8080

install-air:
	@if [ ! -f ./bin/air ]; then \
		echo "installing air-verse package"; \
		curl -sSfL https://raw.githubusercontent.com/air-verse/air/master/install.sh | sh -s; \
	else \
		echo "./bin/air already exists"; \
	fi

server-down: ## Kill server running on port $(SERVER_PORT)
	@PID=$$(lsof -t -i :$(SERVER_PORT)); \
	if [ -n "$$PID" ]; then \
		echo "Killing server running on port $(SERVER_PORT)..."; \
		kill $$PID; \
	else \
		echo "Server not running"; \
	fi

server-up: server-down ## Start the server on given port (default 8080)
	go run server/main.go server --port $(SERVER_PORT)

server-dev: ## Start the server on given port (default 8080) + support live reloading.
	go mod tidy
	./bin/air -c .air.toml

go-test: ## Run go tests
	@go test ./... -race -coverprofile=coverage.out -covermode=atomic

go-coverage: go-test ## Get overall coverage
	@echo "Generating coverage report (file-level)..."
	@go tool cover -func=coverage.out \
		| grep -v "mock" \
		| grep -E "total|^testing" \
		| awk '{ print $$1, $$3 }'



BINARY_NAME := quickbite
OUTPUT_DIR := ./bin
PLATFORMS := linux/amd64 linux/arm64 darwin/amd64 darwin/arm64

server-build: ## Build Go binaries for linux and darwin platforms (amd64 and arm64)
	@mkdir -p $(OUTPUT_DIR)
	@for platform in $(PLATFORMS); do \
		OS=$$(echo $$platform | cut -d'/' -f1); \
		ARCH=$$(echo $$platform | cut -d'/' -f2); \
		echo "Building for $$OS/$$ARCH..."; \
		mkdir -p $(OUTPUT_DIR)/$$OS/$$ARCH; \
		GOOS=$$OS GOARCH=$$ARCH go build -o $(OUTPUT_DIR)/$$OS/$$ARCH/$(BINARY_NAME) ./server; \
	done
