.PHONY: run build migrate seed clean vet

# Run the server in development
run:
	go run ./cmd/server/

# Build binary
build:
	go build -o bin/server ./cmd/server/

# Run database migrations
migrate:
	go run ./cmd/server/ -migrate

# Run database seed
seed:
	go run ./cmd/server/ -seed

# Run go vet
vet:
	go vet ./...

# Clean build artifacts
clean:
	rm -rf bin/ tmp/
