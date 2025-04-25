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

# We manage the rendered template directory explicitly as the rendered templates
# may contain secrets. Use `/tmp` if `TMPDIR` is not set.
rendered_template_dir := $(or $(TMPDIR), /tmp)/bky-basm-rendered

.PHONY: make-rendered-template-dir
make-rendered-template-dir:
	@echo "rendered templates: $(rendered_template_dir)"
	@mkdir -p $(rendered_template_dir)

.PHONY: delete-rendered-template-dir
delete-rendered-template-dir:
	@rm -rf $(rendered_template_dir)

# Specify all the files that are templates to be rendered
template_src := $(wildcard $(template_dir)/*.mustache)
# Specify all the files that we expect to exist after rendering as renames of the template files
template_out := $(patsubst $(template_dir)/%.mustache,$(rendered_template_dir)/%,$(template_src))
# Mark the rendered files as intermediate. Make will delete them after the build
# if they still exist.
.INTERMEDIATE: $(template_out)

# Set default values for BASM_ variables. These can be overridden by the user
# e.g. `BASM_HOST=http://api.bky.sh/staging/delphi make test`
BASM_PLATFORM ?= plain
BASM_CODE_MEASURE ?= plain
BASM_AUTH_TOKEN ?= auth token
BASM_HOST ?= local-server

# Render each template from its mustache source. This rule is used when make sees
# a dependency that matches the `$(rendered_template_dir)/%` pattern.
# Syntax notes: `$<` represents the first dependency of the rule, which is the
# mustache template file name. `$@` represents the target of the rule, which is
# the rendered file name.
$(rendered_template_dir)/%: $(template_dir)/%.mustache | make-rendered-template-dir
	@jq -n \
		--arg BASM_PLATFORM "$(BASM_PLATFORM)" \
		--arg BASM_CODE_MEASURE "$(BASM_CODE_MEASURE)" \
		--arg BASM_AUTH_TOKEN "$(BASM_AUTH_TOKEN)" \
		--arg BASM_HOST "$(BASM_HOST)" \
		'$$ARGS.named' \
		| mustache $< > $@

# Run the integration test and delete the temp dir afterward
.PHONY: test-integration
test-integration: $(wasm_out) $(template_out)
	@txtar-c $(rendered_template_dir) \
		| cat "$(script_dir)/integration.txtar" - \
		| testscript -e WASM_FILE=$(wasm_out)
	@$(MAKE) -s delete-rendered-template-dir

.PHONY: test-local
test-local: $(wasm_out)
	@testscript -e WASM_FILE=$(wasm_out) "$(script_dir)/local.txtar"

.PHONY: test
test: test-local test-integration
