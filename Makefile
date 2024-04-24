GO = $(shell which go 2>/dev/null)

APP 			:= go-bootstrap
VERSION 		:= v0.1.0


.PHONY: all build clean test

all: clean build

clean:
	$(RM) -rf bin/*
build:
	$(GO) build -o bin/$(APP) cmd/genesis/*.go
run:
	$(GO) run cmd/genesis/*.go
test:
	$(GO) test -v ./...
