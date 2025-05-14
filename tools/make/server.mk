# server.mk

SERVER_PORT = 8080

.PHONE: install-air
install-air:
	@if [ ! -f ./bin/air ]; then \
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
