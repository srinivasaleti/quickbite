compose-up: ## Start Docker Containers using `docker compose up`	
	docker compose -f tools/docker/quickbite/docker-compose.yaml up --build --watch &

compose-down: ## Stop Docker Containers using `docker compose down`	
	docker compose -f tools/docker/quickbite/docker-compose.yaml down

compose-quickbitedb-up: ## Starts only quickbite db with docker compose
	docker compose -f tools/docker/quickbite/docker-compose.yaml up quickbite-db --build &
