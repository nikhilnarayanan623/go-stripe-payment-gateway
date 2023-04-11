SHELL := /bin/bash

.PHONY: all build test deps deps-cleancache
 

GOCMD=go 

BUILD_DIR=build

BINARY_DIR=$(BUILD_DIR)/bin  

 # to start the application
run:
	@echo "Stripe Payment-Gateway"

	$(GOCMD) run ./cmd/api


# to install dependencies packges latest version if its not in local package
deps: 
	$(GOCMD) get -u -t -d -v ./...
#remove un used dependencies
	$(GOCMD) mod tidy 
 # to clean cache in the module
dps-cleancache:
	$(GOCMD) clean -modcache

 ## Generate wire_gen.go
wire:
	cd pkg/di && wire

## Generate swagger docs
swag: 
	swag init -g pkg/api/server.go -o ./cmd/api/docs
 
## Display this help screen
help:
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
