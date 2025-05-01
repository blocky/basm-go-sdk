MAKEFLAGS += --warn-undefined-variables
SHELL := bash
.DEFAULT_GOAL := build
.DELETE_ON_ERROR:

# Project directories
testdata_dir := $(shell realpath ./test/testdata)
script_dir := $(shell realpath ./test/scripts)
template_dir :=  $(shell realpath ./test/templates)

.PHONY: generate
generate:
	go generate ./...

.PHONY: build
build: generate wasm

.PHONY: lint
lint:
	golangci-lint run --config golangci.yaml

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: pre-pr
pre-pr: tidy generate lint test

.PHONY: clean
clean:
	-@rm test/testdata/x.wasm

go_sources := $(shell find . -type f -name "*.go" ! -name "*_test.go")
wasm_src := $(testdata_dir)/main.go
wasm_out := $(testdata_dir)/x.wasm

.PHONY: wasm
wasm: $(wasm_out)

$(wasm_out): $(go_sources)
	@echo "Building WASM module..."
	@tinygo build -o $@ -target=wasi $(wasm_src)

# Set default values for BASM_ variables. These can be overridden by the user
# e.g. `BASM_HOST=http://api.bky.sh/staging/delphi make test`
BASM_PLATFORM ?= plain
BASM_CODE_MEASURE ?= plain
BASM_AUTH_TOKEN ?= auth token
BASM_HOST ?= local-server

.PHONY: test-integration
test-integration: $(wasm_out)
# Use jq and mustache to render env vars into the input files for configuration
# and secrets, concatenate with the test script and pipe the whole txtar
# package to testscript.
	@jq -n \
		--arg BASM_PLATFORM "$(BASM_PLATFORM)" \
		--arg BASM_CODE_MEASURE "$(BASM_CODE_MEASURE)" \
		--arg BASM_AUTH_TOKEN "$(BASM_AUTH_TOKEN)" \
		--arg BASM_HOST "$(BASM_HOST)" \
		'$$ARGS.named' \
		| mustache $(template_dir)/integration-files.txtar.mustache \
		| cat "$(script_dir)/integration.txtar" - \
		| testscript -e WASM_FILE=$(wasm_out)

.PHONY: test-local
test-local: $(wasm_out)
	@testscript -e WASM_FILE=$(wasm_out) "$(script_dir)/local.txtar"

.PHONY: test
test: test-local test-integration
