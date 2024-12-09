build:
	go build -o bin/main cmd/service/main.go

run:
	go run cmd/service/main.go

test:
	go test ./tests

migrate:
	go run cmd/migrate/migrate.go

