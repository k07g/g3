BINARY := ohayo
BIN_DIR := bin

.PHONY: build run install test vet fmt-check check clean help

help:
	@echo "Usage: make <target>"
	@echo ""
	@echo "Targets:"
	@echo "  build   Build the binary to $(BIN_DIR)/$(BINARY)"
	@echo "  run     Run without building"
	@echo "  test    Run tests"
	@echo "  vet     Run go vet"
	@echo "  fmt-check Check formatting with gofmt"
	@echo "  check   Run vet and fmt-check"
	@echo "  install Install the binary via go install"
	@echo "  clean   Remove the binary"
	@echo "  help    Show this help"

build:
	go build -o $(BIN_DIR)/$(BINARY) .

test:
	go test -v ./...

vet:
	go vet ./...

fmt-check:
	@test -z "$(shell gofmt -l .)" || (echo "gofmt diff:"; gofmt -d .; exit 1)

check: vet fmt-check

install: build
	cp $(BIN_DIR)/$(BINARY) $(shell go env GOPATH)/bin/$(BINARY)

run:
	go run .

clean:
	rm -f $(BIN_DIR)/$(BINARY)
