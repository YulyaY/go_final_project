build:
	go build -o bin/main cmd/service/main.go
	./bin/main

run:
	go run cmd/service/main.go

test:
	go test ./tests

