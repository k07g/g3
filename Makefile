BINARY := ohayo
BIN_DIR := bin

.PHONY: build run install clean help

help:
	@echo "Usage: make <target>"
	@echo ""
	@echo "Targets:"
	@echo "  build   Build the binary to $(BIN_DIR)/$(BINARY)"
	@echo "  run     Run without building"
	@echo "  install Install the binary via go install"
	@echo "  clean   Remove the binary"
	@echo "  help    Show this help"

build:
	go build -o $(BIN_DIR)/$(BINARY) .

install: build
	cp $(BIN_DIR)/$(BINARY) $(shell go env GOPATH)/bin/$(BINARY)

run:
	go run .

clean:
	rm -f $(BIN_DIR)/$(BINARY)
