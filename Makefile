# Makefile for Terraform Rudder provider

.PHONY: build install test clean fmt vet help

PROVIDER_NAME=terraform-provider-rudder
PROVIDER_VERSION=1.0.0
PROVIDER_NAMESPACE=local/cloudducoeur

# Detect OS and architecture
UNAME_S := $(shell uname -s)
UNAME_M := $(shell uname -m)

ifeq ($(UNAME_S),Linux)
    OS := linux
endif
ifeq ($(UNAME_S),Darwin)
    OS := darwin
endif

ifeq ($(UNAME_M),x86_64)
    ARCH := amd64
endif
ifeq ($(UNAME_M),arm64)
    ARCH := arm64
endif
ifeq ($(UNAME_M),aarch64)
    ARCH := arm64
endif

TERRAFORM_PLUGINS_DIR := $(HOME)/.terraform.d/plugins/$(PROVIDER_NAMESPACE)/rudder/$(PROVIDER_VERSION)/$(OS)_$(ARCH)

## Build the provider
build:
	@echo "🚀 Building the provider..."
	go build -o $(PROVIDER_NAME)

## Install the provider locally
install: build
	@echo "📦 Installing the provider locally..."
	mkdir -p $(TERRAFORM_PLUGINS_DIR)
	cp $(PROVIDER_NAME) $(TERRAFORM_PLUGINS_DIR)/

## Format the code
fmt:
	@echo "🎨 Formatting the code..."
	go fmt ./...

## Check the code with vet
vet:
	@echo "🔍 Checking the code..."
	go vet ./...

## Run tests
test:
	@echo "🧪 Running tests..."
	go test ./... -v

## Clean build files
clean:
	@echo "🧹 Cleaning..."
	rm -f $(PROVIDER_NAME)
	rm -rf ./examples/.terraform*
	rm -f ./examples/terraform.tfstate*

## Test in the examples directory
test-examples: install
	@echo "🔍 Testing in the examples directory..."
	cd examples && terraform init && terraform validate

## Generate documentation
docs:
	@echo "📖 Generating documentation..."
	go generate ./...

## Display help
help:
	@echo "Available commands:"
	@grep -E '^## .*' $(MAKEFILE_LIST) | sed 's/## /  /' | column -t -s ':'

# Default command
default: build
