crypto-bot: install build

test: lint test-unit

test-unit:
	@echo "=== Running unit tests ==="
	@go test -mod=vendor -tags=unit -failfast -coverprofile=test.cover ./...
	@go tool cover -func=test.cover
	@rm -f test.cover

lint:
	@echo "=== Checking code ==="
	golangci-lint run --modules-download-mode vendor --timeout=30m

install:
	echo "=== Installing dependencies ==="
	go mod tidy
	go mod vendor

build: 
	go build -mod=vendor .