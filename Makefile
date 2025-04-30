MAKEFLAGS += --warn-undefined-variables
SHELL := bash
.DEFAULT_GOAL := build
.DELETE_ON_ERROR:

# Project directories
testdata_dir := $(shell realpath ./test/testdata)
script_dir := $(shell realpath ./test/scripts)
template_dir :=  $(shell realpath ./test/templates)

# add easyjson source file targets here
easyjson_sources := basm/dto.go
easyjson_generated := $(easyjson_sources:.go=_easyjson.go)

# Rule to generate *_easyjson.go files
%_easyjson.go: %.go
	@easyjson $<

# Generate all *_easyjson.go files
.PHONY: generate
generate: $(easyjson_generated)

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

sdk_srcs := $(wildcard ./basm/**/*.go ./x/**/*.go)
wasm_src := $(testdata_dir)/main.go
wasm_out := $(testdata_dir)/x.wasm

.PHONY: wasm
wasm: $(wasm_out)

$(wasm_out): $(wasm_src) $(sdk_srcs)
	@echo "Building WASM module..."
	@tinygo build -o $@ -target=wasi $<

# Set default values for BASM_ variables. These can be overridden by the user
# e.g. `BASM_HOST=http://api.bky.sh/staging/delphi make test`
BASM_PLATFORM ?= plain
BASM_CODE_MEASURE ?= plain
BASM_AUTH_TOKEN ?= auth token
BASM_HOST ?= local-server

# Run the integration test and delete the temp dir afterward
.PHONY: test-integration
test-integration: $(wasm_out)
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
