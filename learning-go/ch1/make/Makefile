.DEFAULT_GOAL := build

fmt:
		go fmt
.PHONY: fmt

lint: fmt
		go vet
.PHONY: vet

build: vet
		go mod tidy
		go build
.PHONY: build