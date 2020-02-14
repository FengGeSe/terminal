# info
VERSION=v0.0.1

PROJECT_NAME=$(shell basename "$(PWD)")
GO_VERSION=$(shell go version  | awk '{print $$3}')
BUILD_TIME=$(shell date +%FT%T%z)
OS_ARCH=$(shell go version  | awk '{print $$4}')
GIT_COMMIT=$(shell git rev-parse HEAD)
GO_BIN=$(GOBIN)
GO_MODULE=$(shell sed -n '/module/p'  go.mod | awk '{print $$2}')

## help: Help for this project
help: Makefile
	@echo "Usage:\n  make [command]"
	@echo
	@echo "Available Commands:"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'

## build: Compile the binary.
build:
	@go build -o bin/$(OS_ARCH)/cli cmd/client/main.go
	@go build -o bin/$(OS_ARCH)/server cmd/server/main.go

## install: build and install.
install:
	@go build -o bin/$(OS_ARCH)/cli cmd/client/main.go
	@go build -o bin/$(OS_ARCH)/server cmd/server/main.go
	@mv bin/$(OS_ARCH)/* $(GO_BIN)

## build-linux: Compile the linux binary.
build-linux:
	@GOOS=linux GOARCH=amd64 go build -o bin/linux/amd64/cli cmd/client/main.go
	@GOOS=linux GOARCH=amd64 go build -o bin/linux/amd64/server cmd/server/main.go

## build-mac: Compile the linux binary.
build-mac:
	@GOOS=darwin GOARCH=amd64 go build -o bin/darwin/amd64/cli cmd/client/main.go
	@GOOS=darwin GOARCH=amd64 go build -o bin/darwin/amd64/server cmd/server/main.go

## clean: Clean build files.
clean:
	rm -rf bin/darwin
	rm -rf bin/linux
	rm -rf bin/windows
