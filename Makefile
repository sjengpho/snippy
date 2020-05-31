.DEFAULT_GOAL := build
GO_BIN := $(or $(GOBIN),$(HOME)/go/bin)

build:
	go build -o ./bin/snippy cmd/cli/main.go

install:
	go build -o $(GO_BIN)/snippy cmd/cli/main.go

test:
	go test ./...
