include .env

run:
	@go run ./cmd/api/

generate:
	@sqlc generate

migrateup:
	@goose -dir="./internal/migrations/" mysql ${DATABASE_URL} up
	
migratedown:
	@goose -dir="./internal/migrations/" mysql ${DATABASE_URL} down
