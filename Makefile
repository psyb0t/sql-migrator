SERVICE_NAME := "sql-migrator"
MODULE_NAME := "github.com/psyb0t/${SERVICE_NAME}"
PKG_LIST := $(shell go list ${MODULE_NAME}/...)
MIN_TEST_COVERAGE := 80
GO111MODULE=on
CGO_ENABLED=0
ENTRYPOINT := "./cmd/main.go"

.PHONY: build

all: dep lint test build ## Run dep, lint, test, and build

dep: ## Get project dependencies
	@echo "Getting project dependencies..."
	@go mod tidy

lint: ## Lint all Golang files
	@echo "Linting all Golang files..."
	@golangci-lint run --timeout=30m0s

test: ## Run all tests
	@echo "Running all tests..."
	@go test -race $(PKG_LIST)

test-coverage: ## Run tests with coverage check. Fails if coverage is below the threshold.
	@echo "Running tests with coverage check..."
	@trap 'rm -f coverage.txt' EXIT; \
	go test -race -coverprofile=coverage.txt $(PKG_LIST); \
	if [ $$? -ne 0 ]; then \
		echo "Test failed. Exiting."; \
		exit 1; \
	fi; \
	result=$$(go tool cover -func=coverage.txt | grep -oP 'total:\s+\(statements\)\s+\K\d+' || echo "0"); \
	if [ $$result -eq 0 ]; then \
		echo "No test coverage information available."; \
		exit 0; \
	elif [ $$result -lt $(MIN_TEST_COVERAGE) ]; then \
		echo "FAIL: Coverage $$result% is less than the minimum $(MIN_TEST_COVERAGE)%"; \
		exit 1; \
	fi

BUILD_LIST := \
    build-linux-amd64 \
    build-windows-amd64 \
    build-darwin-amd64 \

build: $(BUILD_LIST) ## Build for all specified architectures

build-linux-amd64: ## Build for Linux amd64
	@echo "Building for Linux amd64..."
	@GOOS=linux GOARCH=amd64 go build -a -o build/$(SERVICE_NAME)-linux-amd64 $(ENTRYPOINT)

build-windows-amd64: ## Build for Windows amd64
	@echo "Building for Windows amd64..."
	@GOOS=windows GOARCH=amd64 go build -a -o build/$(SERVICE_NAME)-windows-amd64.exe $(ENTRYPOINT)

build-darwin-amd64: ## Build for Darwin amd64
	@echo "Building for Darwin amd64..."
	@GOOS=darwin GOARCH=amd64 go build -a -o build/$(SERVICE_NAME)-darwin-amd64 $(ENTRYPOINT)

generate: ## Generate all the shit that can be generated
	@go generate ./...

clean: ## Remove build artifacts
	@echo "Removing build artifacts..."
	@rm -rf build

help: ## Display this help message
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
