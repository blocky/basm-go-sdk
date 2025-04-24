MAKEFLAGS += --warn-undefined-variables
SHELL := bash
.DEFAULT_GOAL := build
.DELETE_ON_ERROR:

# Project directories
testdata_dir := $(shell realpath ./test/testdata)
script_dir := $(shell realpath ./test/scripts)
template_dir :=  $(shell realpath ./test/templates)
rendered_template_dir := $(shell mktemp -d -t bky-basm-rendered.XXXXXX)

# add easyjson source file targets here
easyjson_src := basm/dto.go
easyjson_out := $(easyjson_src:.go=_easyjson.go)

# Rule to generate *_easyjson.go files
%_easyjson.go: %.go
	@easyjson $<

# Generate all *_easyjson.go files
.PHONY: generate
generate: $(easyjson_out)

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

sdk_src := $(wildcard ./basm/**/*.go)
wasm_src := $(testdata_dir)/main.go
wasm_out := $(testdata_dir)/x.wasm

.PHONY: wasm
wasm: $(wasm_out)

$(wasm_out): $(wasm_src) $(sdk_src)
	@echo "Building WASM module..."
	@tinygo build -o $@ -target=wasi $<

# Set default values for BASM_ variables. These can be overridden by the user
# e.g. `BASM_USER_SECRET=mysecret make test`
BASM_PLATFORM ?= plain
BASM_CODE_MEASURE ?= plain
BASM_AUTH_TOKEN ?= auth token
BASM_HOST ?= local-server
BASM_USER_SECRET ?= this is a bearer token
# Export all BASM_ variables for use in recipes
export $(filter BASM_%,$(.VARIABLES))

template_src := $(wildcard $(template_dir)/*.tmpl)

template_out := $(patsubst $(template_dir)/%.tmpl,$(rendered_template_dir)/%,$(template_src))
# Mark the rendered files as intermediate. Make will delete them after the build.
.INTERMEDIATE: $(template_out)

# Rule to render template files, `envsubst <template_file> > <output_file>`
$(rendered_template_dir)/%: $(template_dir)/%.tmpl
	@envsubst < $< > $@

.PHONY: test-integration
test-integration: $(template_out) $(wasm_out)
	@txtar-c $(rendered_template_dir) \
		| cat "$(script_dir)/integration.txtar" - \
		| testscript -e WASM_FILE=$(wasm_out)

.PHONY: test-local
test-local: $(wasm_out)
	@testscript -e WASM_FILE=$(wasm_out) "$(script_dir)/local.txtar"

.PHONY: test
test: test-local test-integration
