build:
	go build -o bin/main cmd/service/main.go
	./bin/main

test:
	go test ./...