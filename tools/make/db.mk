install-migrate-cli: ## Install migrate cli
	@go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	@echo "migrate CLI tool installed successfully."

create-migration: install-migrate-cli ## Create migration. Syntax create-migration name=<name>
	@if [ -z "$(name)" ]; then \
		echo "Please provide a migration name. Usage: make create-migration name=<migration_name>"; \
	else \
		PATH=$$PATH:$(shell go env GOPATH)/bin migrate create -format unix -ext sql -dir server/internal/database/migrations $(name); \
	fi