BIN_DIR=bin/
BIN_NAME=wait-for

build:
	GO111MODULE=on go build -o $(BIN_DIR)$(BIN_NAME) ./...