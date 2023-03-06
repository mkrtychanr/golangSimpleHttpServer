.PHONY: server
server:
	go build -v ./cmd/server
	./server

.PHONY: migrate
migrate:
	go build -v ./cmd/migrate
	./migrate

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build