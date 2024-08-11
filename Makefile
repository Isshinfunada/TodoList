start:
	docker-compose up -d --build

.PHONY: start

stop:
	docker-compose down -v

.PHONY: stop

restart:
	make stop
	make start

.PHONY: restart

.PHONY: migrate-up
migrate-up:
	@echo "Applying migrations..."
	migrate -path server/db/migrations -database "postgres://user:password@localhost:5432/todos?sslmode=disable" up

.PHONY: migrate-down
migrate-down:
	@echo "Reverting migrations..."
	migrate -path server/db/migrations -database "postgres://user:password@localhost:5432/todos?sslmode=disable" down