.PHONY: help, test-cover, prepare-cover-report, open-cover-report, update-readme-cover, update-readme-doc, test, test-regression, configure, tidy, imports, fmt, format-project, install-tools, vet, lint-full, lint-pipeline, lint-sonar, release, release-local-snapshot, check-releaser, new-version
BIN_DIR=./bin

SHELL := env VERSION=$(VERSION) $(SHELL)
VERSION ?= $(shell git describe --tags $(git rev-list --tags --max-count=1))

APP_NAME?=getenv
SHELL := env APP_NAME=$(APP_NAME) $(SHELL)

GOVERSION:=1.27.4

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


## Test coverage report.
test-cover:
	./scripts/tests/coverage.sh

prepare-cover-report: test-cover
	./scripts/tests/prepare-cover-report.sh

## Open coverage report.
open-cover-report: prepare-cover-report
	./scripts/open-coverage-report.sh

## Update readme coverage.
update-readme-cover: prepare-cover-report
	./scripts/update-readme-coverage.sh

## Update readme doc.
update-readme-doc:
	./scripts/update-readme-doc.sh

## Run tests.
test:
	./scripts/tests/run.sh

## Run regression tests.
test-regression: test

## Tidy go.mod and install needed tools.
configure: tidy install-tools

## Tidy and verify go.mod.
tidy:
	./scripts/tidy.sh

## Fix imports sorting.
imports:
	./scripts/style/fix-imports.sh

## Format code with go fmt.
fmt:
	./scripts/style/fmt.sh

## Format code and sort imports.
format-project: fmt imports

## Build tools pinned in tools/<name>/go.mod.
install-tools:
	./scripts/install-tools.sh

## vet project
vet:
	./scripts/linting/run-vet.sh

## Run full linting
lint-full:
	./scripts/linting/run-linters.sh

## Run linting for build pipeline
lint-pipeline:
	./scripts/linting/golangci-pipeline.sh

## Run linting for sonar report
lint-sonar:
	./scripts/linting/golangci-sonar.sh

## Release
release:
	./scripts/release/release.sh

## Release local snapshot
release-local-snapshot:
	./scripts/release/local-snapshot-release.sh

## Check goreleaser config.
check-releaser:
	./scripts/release/check.sh

## Issue new release.
new-version: vet test-regression
	./scripts/release/new-version.sh

bump-go-version:
	./scripts/bump-go.sh $(GOVERSION)
.PHONY: bump-go-version


.DEFAULT_GOAL := help

