
APP_NAME := "geniee-cli"
SHELL := /bin/bash

Version := $(shell git tag --sort=committerdate | tail -1)
LDFLAGS := "-s -w -X github.com/geniee-ai/geniee-cli/version.Version=$(Version)"


gorelease:
	@rm -rf dist
	@goreleaser release --clean
	@rm -rf dist
docker-build:
	@docker build -t geniee-cli .

build-amd64:
	@echo "Compiling amd64 build"
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags $(LDFLAGS) -a -installsuffix cgo -o bin/${APP_NAME}-amd64 ./cmd/cli/main.go

build-armhf:
	@echo "Compiling armhf build"
	@CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=6 go build -ldflags $(LDFLAGS) -a -installsuffix cgo -o bin/${APP_NAME}-armhf ./cmd/cli/main.go

build-arm64:
	@echo "Compiling arm64 build"
	@CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags $(LDFLAGS) -a -installsuffix cgo -o bin/${APP_NAME}-arm64 ./cmd/cli/main.go

build-darwin:
	@echo "Compiling darwin build"
	@CGO_ENABLED=0 GOOS=darwin go build -ldflags $(LDFLAGS) -a -installsuffix cgo -o bin/${APP_NAME}-darwin ./cmd/cli/main.go

build-windows:
	@echo "Compiling windows build"
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags $(LDFLAGS) -a -installsuffix cgo -o bin/${APP_NAME}.exe ./cmd/cli/main.go

help: # Show this help
	@egrep -h '\s#\s' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?# "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'