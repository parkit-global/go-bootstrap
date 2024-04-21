GO = $(shell which go 2>/dev/null)
DOCKER = $(shell which docker 2>/dev/null)

APP 			:= {{.AppName}}
VERSION 		:= v0.1.0


.PHONY: all build clean test docker

all: clean build

clean:
	$(RM) -rf bin/*
build:
	$(GO) build -o bin/$(APP) cmd/*.go
run:
	$(GO) run cmd/*.go
test:
	$(GO) test -v ./...
docker:
	$(DOCKER) build --build-arg APP=$(APP) -t $(APP):$(VERSION) -t $(APP):latest .
