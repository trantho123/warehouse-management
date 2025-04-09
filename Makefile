DB_URL=postgresql://admin:admin123@host.docker.internal:5432/warehouse_db?sslmode=disable
MIGRATIONS_DIR=$(CURDIR)/db/migrations

up:
	docker-compose up -d 

down:
	docker-compose down -d

new_migration:
	docker run --rm -v "$(MIGRATIONS_DIR)":/migrations \
  		migrate/migrate \
		create -ext sql -dir /migrations -seq "$(name)"

migrateup:
	docker run --rm -v "$(MIGRATIONS_DIR)":/migrations migrate/migrate \
		-source=file:///migrations -database "$(DB_URL)" -verbose up

migrateup1:
	docker run --rm -v "$(MIGRATIONS_DIR)":/migrations migrate/migrate \
		-source=file:///migrations -database "$(DB_URL)" -verbose up 1

sqlc:
	docker run --rm \
		-v $(CURDIR):/app \
		-w /app \
		kjconroy/sqlc generate

migratedown:
	docker run --rm -v "$(MIGRATIONS_DIR)":/migrations migrate/migrate \
		-source=file:///migrations -database "$(DB_URL)" -verbose down -all

migratedown1:
	docker run --rm -v "$(MIGRATIONS_DIR)":/migrations migrate/migrate \
		-source=file:///migrations -database "$(DB_URL)" -verbose down 1

.PHONY: up down new_migration migrateup migrateup1 sqlc migratedown migratedown1