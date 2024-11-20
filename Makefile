DB_URL=postgres://root:root_password@localhost:5432/auth_db?sslmode=disable

# Docker section
run:
	@docker compose up

clear-containers:
	@docker compose down

clear-postgres-volume:
	@docker volume rm auth_postgres_data

# Postgres migrations and seed section

setup-db:
	@echo "Migrating database..."
	@make -s migrate-up
	@echo "Seeding db..."
	@make -s seed-db

migrate-up:
	@goose -dir ./migrations postgres "$(DB_URL)" up

migrate-down:
	@goose -dir ./migrations postgres "$(DB_URL)" down

seed-db:
	@if make check-seed; then \
		psql "$(DB_URL)" -f seed/auth.sql && \
		echo "Seeding completed."; \
	else \
		echo "Database is already seeded."; \
	fi

check-seed:
	@psql "$(DB_URL)" -t -c "SELECT COUNT(*) FROM two_factor_types;" | grep -q '^ *0 *$$'



