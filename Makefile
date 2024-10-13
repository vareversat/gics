# Makefile for Go project

# Go parameters
GOCMD = go

.PHONY: deps-upgrade deps


# Main build target
all: deps test

# Clean the build artifacts
clean:
	$(GOCMD) clean

# Run tests
test:
	$(GOCMD) test -v ./...

# Install project dependencies
deps:
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
	$(GOCMD) tool cover -html=coverage.out -o coverage.html

# Generate documentation using tools like godoc
doc:
	godoc -http=:6060

# Perform a full code quality check (lint, tests, coverage)
check: lint test coverage

.PHONY: all clean test deps run fmt lint coverage doc check