include .env

run:
	@go run ./cmd/api/

generate:
	@sqlc generate

migrateup:
	@goose -dir="./internal/migrations/" postgres ${DATABASE_URL} up
	
migratedown:
	@goose -dir="./internal/migrations/" postgres ${DATABASE_URL} down
