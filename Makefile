mod:
	go mod tidy

dev: mod
	go run ./cmd/app

start:
	go build -o ./bin/app ./cmd/app && ./bin/app