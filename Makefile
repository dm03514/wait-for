BIN_DIR=bin/
BIN_NAME=wait-for
PKGS = $(shell go list ./... | grep -v /vendor/)

build:
	GO111MODULE=on go build -o $(BIN_DIR)$(BIN_NAME)

test-unit:
	go test $(PKGS) -v -coverprofile=coverage.out -covermode=count

.PHONY: test-unit