BIN_DIR=./bin

SHELL := env VERSION=$(VERSION) $(SHELL)
VERSION ?= $(shell git describe --tags $(git rev-list --tags --max-count=1))

APP_NAME?=strava-api
SHELL := env APP_NAME=$(APP_NAME) $(SHELL)

API_DOC_URL?=https://developers.strava.com/swagger/swagger.json

COMPOSE_TOOLS_FILE=deployments/docker-compose/go-tools-docker-compose.yml
COMPOSE_TOOLS_CMD_BASE=docker compose -f $(COMPOSE_TOOLS_FILE)
COMPOSE_TOOLS_CMD_UP=$(COMPOSE_TOOLS_CMD_BASE) up --exit-code-from
COMPOSE_TOOLS_CMD_PULL=$(COMPOSE_TOOLS_CMD_BASE) build

GOVERSION:=1.27

TARGET_MAX_CHAR_NUM=20

## Show help
help:
	${call colored, help is running...}
	@echo ''
	@echo 'Usage:'
	@echo '  make <target>'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "  %-$(TARGET_MAX_CHAR_NUM)s %s\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

gen:
	$(COMPOSE_TOOLS_CMD_UP) go-generate go-generate
.PHONY: gen

codegen: gen sync-vendor format-code vet
.PHONY: codegen

sync-vendor:
	./scripts/sync-vendor.sh
.PHONY: sync-vendor

## Fix imports sorting.
imports:
	$(COMPOSE_TOOLS_CMD_UP) fix-imports fix-imports
.PHONY: imports

## Format code with go fmt.
fmt:
	$(COMPOSE_TOOLS_CMD_UP) fix-fmt fix-fmt
.PHONY: fmt

## Format code and sort imports.
format-code: fmt imports
.PHONY: format-code

## Installs vendored tools.
install-tools:
	$(COMPOSE_TOOLS_CMD_PULL)
.PHONY: install-tools

vet:
	go vet ./...
.PHONY: vet

## Run full linting
lint-full:
	$(COMPOSE_TOOLS_CMD_UP) lint-full lint-full
.PHONY: lint-full

bump-go-version:
	./scripts/bump-go.sh $(GOVERSION)
.PHONY: bump-go-version

## Print a Strava access token for integration tests, refreshing it when needed.
strava-token:
	@go run ./internal/tools/strava-token
.PHONY: strava-token

## Run live integration tests against the Strava API.
test-integration:
	STRAVA_ACCESS_TOKEN="$$(go run ./internal/tools/strava-token)" go test -count=1 -tags=integration ./examples ./internal/specaudit
.PHONY: test-integration

## Update the API spec from Strava, keeping the local changes.
update-swagger-spec:
	API_DOC_URL=$(API_DOC_URL) ./scripts/codegen/update-swagger-spec.sh
.PHONY: update-swagger-spec

## Show the local changes of the API spec compared to the Strava one.
swagger-local-changes:
	@git diff --no-index docs/swagger.upstream.json docs/swagger.json; test $$? -le 1
.PHONY: swagger-local-changes


.DEFAULT_GOAL := help
