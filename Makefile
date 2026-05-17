build:
	go build -o bin/task-cli ./cmd/main.go

test:
	go test ./... -v -race

lint:
	golangci-lint run

run: build
	./bin/task-cli

docker-build:
	docker build -t task-cli .

.PHONY: build test lint run docker-build