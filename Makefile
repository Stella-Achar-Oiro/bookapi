.PHONY: build test run clean

build:
	go build -o bin/server cmd/server/main.go

test:
	go test -v ./...

test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

run:
	go run cmd/server/main.go

clean:
	rm -rf bin/
	rm -f coverage.out