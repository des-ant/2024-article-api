# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
# The sed command extracts the comments from the Makefile and prints them in a nice format.
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## run/api: run the cmd/api application
# It checks if the port and env variables are set and passes them to the application.
.PHONY: run/api
run/api:
	@PORT_FLAG=$$(if [ -n "${port}" ]; then echo "--port=${port}"; fi); \
	ENV_FLAG=$$(if [ -n "${env}" ]; then echo "--env=${env}"; fi); \
	go run ./cmd/api $$PORT_FLAG $$ENV_FLAG

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## test: run all tests
.PHONY: test
test:
	@echo 'Running tests...'
	go test -v -race -cover ./...

## tidy: format all .go files, and tidy and vendor module dependencies
.PHONY: tidy
tidy:
	@echo 'Formatting .go files...'
	go fmt ./...
	@echo 'Tidying module dependencies...'
	go mod tidy
	@echo 'Verifying and vendoring module dependencies...'
	go mod verify
	go mod vendor

## lint: run golangci-lint
.PHONY: lint
lint:
	@echo 'Running linter...'
	golangci-lint run

# ==================================================================================== #
# BUILD
# ==================================================================================== #

## build/api: build the cmd/api application
.PHONY: build/api
build/api:
	@echo 'Building cmd/api...'
	go build -v ./cmd/api

## clean: remove build artifacts and installed packages
.PHONY: clean
clean:
	rm -rf api
	go clean -i .
