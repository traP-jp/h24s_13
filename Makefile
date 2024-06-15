.PHONY: up
up:
	docker compose up -d --build

.PHONY: down
down:
	docker compose down -v
