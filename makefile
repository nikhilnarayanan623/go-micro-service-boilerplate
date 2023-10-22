SHELL := /bin/bash

.PHONY: all build test deps deps-cleancache

GOCMD=go
BUILD_DIR=build
BINARY_DIR=$(BUILD_DIR)/bin
CODE_COVERAGE=code-coverage

all: test build

${BINARY_DIR}:
	mkdir -p $(BINARY_DIR)

build: ${BINARY_DIR} ## Compile the code, build Executable File
	$(GOCMD) build -o $(BINARY_DIR)/api-gateway -v ./api-gateway/cmd/api
	$(GOCMD) build -o $(BINARY_DIR)/auth-service -v ./auth-service/cmd/api
	$(GOCMD) build -o $(BINARY_DIR)/employee-service -v ./employee-service/cmd/api

build-api-gateway: ${BINARY_DIR} ## Compile the code, build Executable File For API Gateway Service
	$(GOCMD) build -o $(BINARY_DIR)/api-gateway -v ./api-gateway/cmd/api

build-auth-service: ${BINARY_DIR} ## Compile the code, build Executable File For Auth Service
	$(GOCMD) build -o $(BINARY_DIR)/auth-service -v ./auth-service/cmd/api

build-employee-service: ${BINARY_DIR} ## Compile the code, build Executable File For Employee Service
	$(GOCMD) build -o $(BINARY_DIR)/employee-service -v ./employee-service/cmd/api

run-api-gateway: build-api-gateway ## Start API Gateway
	./$(BINARY_DIR)/api-gateway

run-auth-service: build-auth-service ## Start Auth Service
	./$(BINARY_DIR)/auth-service

run-employee-service: build-employee-service ## Start Employee Service
	./$(BINARY_DIR)/employee-service

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'