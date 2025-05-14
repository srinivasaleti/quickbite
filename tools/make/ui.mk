# ui.mk

ui-install:  ## Install UI package dependencies
	cd ui && yarn install

ui-dev: ui-install  ## Run UI in dev mode
	cd ui && yarn dev &

ui-down: # Kill ui which runs on port 5173 by default
	@PID=$$(lsof -t -i :5173); \
    if [ -n "$$PID" ]; then \
        echo "killing ui running on port 5173..."; \
        kill $$PID; \
    else \
        echo "UI is not running"; \
    fi
