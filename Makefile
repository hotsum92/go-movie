.PHONY: run
run:
	@go run cmd/main.go

.PHONY: generate
generate:
	@go generate ./...

.PHONY: test
test:
	@go test ./...
