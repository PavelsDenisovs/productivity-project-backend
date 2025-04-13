# 🧩 Migrations Guide

Create a new migration:
migrate create -ext sql -dir db/migrations add_column_to_users
Apply all migrations:
migrate -database "postgres://user:password@localhost:5432/postgres?sslmode=disable" -path db/migrations up
Roll back the last migration:
migrate -database "postgres://user:password@localhost:5432/postgres?sslmode=disable" -path db/migrations down 1
Reset database:
migrate -database "..." -path db/migrations force 0
For production (NeonDB):
migrate -database "$PROD_DB_URL" -path db/migrations up