include .env

run:
	@go run ./cmd/api/

generate:
	@sqlc generate

migrateup:
	@goose -dir="./internal/migrations/" mysql ${DATABASE_URL} up
	# @goose -dir="./internal/migrations/" mysql "sparrow:password@/breve?parseTime=true" up
	
migratedown:
	@goose -dir="./internal/migrations/" mysql ${DATABASE_URL} down
	# @goose -dir="./internal/migrations/" mysql "sparrow:password@/breve?parseTime=true" down
