mod:
	go mod tidy

run: mod
	go run ./cmd/app