# server.mk

SERVER_PORT = 8080

server-down: ## Kill server running on port $(SERVER_PORT)
	@PID=$$(lsof -t -i :$(SERVER_PORT)); \
	if [ -n "$$PID" ]; then \
		echo "Killing server running on port $(SERVER_PORT)..."; \
		kill $$PID; \
	else \
		echo "Server not running"; \
	fi

server-dev: server-down ## Start the server on given port (default 8080)
	go mod tidy
	go run server/main.go server --port $(SERVER_PORT)