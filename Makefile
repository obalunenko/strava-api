gen:
	./scripts/generate.sh
.PHONY: gen

codegen: gen sync-vendor format-code

sync-vendor:
	./scripts/sync-vendor.sh
.PHONY: sync-vendor

format-code: fmt imports
.PHONY: format-code

fmt:
	./scripts/fmt.sh
.PHONY: fmt

imports:
	./scripts/fix-imports.sh
.PHONY: imports

vet:
	go vet ./...
.PHONY: vet