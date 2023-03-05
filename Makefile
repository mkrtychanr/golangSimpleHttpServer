.PHONY: build
build:
	go build -v ./cmd/apiserver

.PHONY: migrate
migrate:
	go build -v ./cmd/migrate
	./migrate

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build