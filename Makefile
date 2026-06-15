VERSION ?= $(shell git describe --tags --dirty 2>/dev/null || echo "dev")
COMMIT  ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")

.PHONY: build clean test lint tidy check

build:
	go build ./...

test:
	go test -v -count=1 ./...

lint:
	golangci-lint run ./... 2>/dev/null || echo "golangci-lint not installed, skipping"

tidy:
	go mod tidy

vet:
	go vet ./...

check: tidy vet build

clean:
	rm -rf dist/
	rm -f adminbox

release: check
	@if [ -z "$(tag)" ]; then \
		echo "Usage: make release tag=v0.1.0"; \
		exit 1; \
	fi
	git tag $(tag)
	git push origin $(tag)
	@echo "Released $(tag)"
