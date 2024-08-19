.PHONY: format
format:
	gofumpt -l -w .

.PHONY: lint
lint:
	golangci-lint run --timeout 3m ./...

.PHONY: list-fix
lint-fix:
	golangci-lint run --timeout 3m --fix ./...

.PHONY: test
test:
	go test ./pkg/...

.PHONY: test-coverage
test-coverage:
	go test -coverprofile=coverage.out ./pkg/...
	go tool cover -html=coverage.out
	open cover.html

.PHONY: docs
docs:
	go doc -all
