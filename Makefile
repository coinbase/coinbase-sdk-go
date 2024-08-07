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
	go test ./...

.PHONY: docs
docs:
	go doc -all
