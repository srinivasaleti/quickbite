help:
	@echo ""
	@echo "Usage:"
	@echo "  make <target>"
	@echo ""
	@echo "Available targets:"
	@echo "  app-dev SERVER_PORT=8080 :-  Start the entire app on given port (default 8080)"
	@echo "  app-down :- Stop the application"
	@echo "  go-coverage :- Run tests and return coverage report"

app-dev: app-down ui-dev server-dev ## Run server first, then UI

app-down: server-down ui-down