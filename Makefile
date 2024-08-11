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