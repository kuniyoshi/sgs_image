.DEFAULT_GOAL := build

fmt:
	cd app && go fmt ./...
	cd scenario && go fmt ./...
.PHONY: fmt

lint: fmt
	staticcheck
.PHONY: lint

vet: lint
	go vet
.PHONY: vet

build: vet
	go mod tidy
	go build
.PHONY: build
