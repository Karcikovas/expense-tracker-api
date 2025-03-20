SHELL := /bin/bash

PHONY: wire
wire:
	wire
PHONY: deps
deps: wire
	go mod tidy
	go mod vendor

PHONY: build
build: deps
	go build -o bin/main

PHONY: serve
serve: build
	bin/main

PHONY: dev
dev: lint
	go run ./

PHONY:lint
lint:
	golangci-lint run ./...