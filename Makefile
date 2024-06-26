# Makefile for Go project

# Go parameters
GOCMD = go
BINARY_NAME_WIN = chabo-api.exe
BINARY_NAME_UNIX = chabo-api.o

.PHONY: deps-upgrade deps


# Main build target
all: deps test build

# Build the application
build:
	$(GOCMD) build -o $(BINARY_NAME_WIN) -v

# Build for Linux
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOCMD) build -o $(BINARY_NAME_UNIX) -v

# Clean the build artifacts
clean:
	$(GOCMD) clean
	rm -f $(BINARY_NAME_WIN)
	rm -f $(BINARY_NAME_UNIX)

# Run tests
test:
	$(GOCMD) test -v ./...

# Install project dependencies
deps:
	$(GOCMD) mod download
	$(GOCMD) mod tidy
	$(GOCMD) mod vendor

# Upgrade project dependencies
deps-upgrade:
	$(GOCMD) get -u -t ./...
	$(GOCMD) mod download
	$(GOCMD) mod tidy
	$(GOCMD) mod vendor

# Format the code
fmt:
	$(GOCMD) fmt ./...
	golines . -w --ignored-dirs=vendor

# Lint the code using a linter tool
lint:
	golangci-lint run

# Generate code coverage report
coverage:
	$(GOCMD) test -coverprofile='coverage.out' ./...
	$(GOCMD) tool cover -html=coverage.out

# Generate documentation using tools like godoc
doc:
	godoc -http=:6060

# Perform a full code quality check (lint, tests, coverage)
check: lint test coverage

.PHONY: all build build-linux clean test deps run fmt lint coverage doc check