.PHONY: up
up:
	docker compose -f docker/compose.yaml up --build -d

.PHONY: down
down:
	docker compose -f docker/compose.yaml down --remove-orphans