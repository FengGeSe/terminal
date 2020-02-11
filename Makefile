# info
VERSION=v0.0.1

PROJECT_NAME=$(shell basename "$(PWD)")
GO_VERSION=$(shell go version  | awk '{print $$3}')
BUILD_TIME=$(shell date +%FT%T%z)
OS_ARCH=$(shell go version  | awk '{print $$4}')
GIT_COMMIT=$(shell git rev-parse HEAD)
CGO_ENABLED=0
GO_BIN=$(GOBIN)

GO_MODULE=$(shell sed -n '/module/p'  go.mod | awk '{print $$2}')

## help: Help for this project
help: Makefile
	@echo "Usage:\n  make [command]"
	@echo
	@echo "Available Commands:"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	
## build: Compile the binary.
build-client:
	@go build -o bin/client cmd/client/main.go

## install: build and install.
install:
	@go build -o bin/$(PROJECT_NAME)
	@mv $(PROJECT_NAME) $(GO_BIN)

## build-linux: Compile the linux binary.
build-linux:
	@CGO_ENABLED=$(CGO_ENABLED) GOOS=linux GOARCH=amd64 go build -o $(PROJECT_NAME) -ldflags "$(LDFLAGS)"

## run-client: run client
run-client:
	@go run cmd/client/main.go

## run-server: run server
run-server:
	@go run cmd/server/main.go

## clean: Clean build files.
clean:
	rm -f bin/*
	
