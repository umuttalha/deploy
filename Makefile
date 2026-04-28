BINARY := deploy
PKG    := github.com/umuttalha/deploy
VERSION ?= dev
LDFLAGS := -X $(PKG)/internal/version.Version=$(VERSION)

.PHONY: build install test lint fmt run tidy clean

build:
	go build -ldflags "$(LDFLAGS)" -o bin/$(BINARY) ./cmd/$(BINARY)

install:
	go install -ldflags "$(LDFLAGS)" ./cmd/$(BINARY)

test:
	go test ./...

lint:
	golangci-lint run

fmt:
	gofmt -s -w .
	goimports -w -local $(PKG) .

run: build
	./bin/$(BINARY)

tidy:
	go mod tidy

clean:
	rm -rf bin/ coverage.out coverage.html
