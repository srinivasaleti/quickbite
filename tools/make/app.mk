help:
	@echo ""
	@echo "Usage:"
	@echo "  make <target>"
	@echo ""
	@echo "Available targets:"
	@echo "  app-dev SERVER_PORT=8080 :-  Start the entire app on given port (default 8080)"
	@echo "  app-down :- Stop the application"

app-dev: app-down ui-dev server-dev ## Run server first, then UI

app-down: server-down ui-down