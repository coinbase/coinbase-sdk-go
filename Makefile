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
	go test ./pkg/coinbase/... ./pkg/auth/...

.PHONY: test-coverage
test-coverage:
	go test -coverprofile=coverage.out ./pkg/coinbase/... ./pkg/auth/...
	go tool cover -html=coverage.out
	open cover.html

.PHONY: mocks
mocks:
	mockery --disable-version-string --name AssetsAPI --keeptree --dir gen/client --output pkg/mocks
	mockery --disable-version-string --name StakeAPI --keeptree --dir gen/client --output pkg/mocks
	mockery --disable-version-string --name ValidatorsAPI --keeptree --dir gen/client --output pkg/mocks

.PHONY: docs
docs:
	go doc -all
