MAKEFLAGS += --warn-undefined-variables
SHELL := bash
.DEFAULT_GOAL := generate
.DELETE_ON_ERROR:

# add easyjson source file targets here
easyjson_sources := basm/dto.go x/xbasm/parse.go
easyjson_generated := $(easyjson_sources:.go=_easyjson.go)

# Rule to generate *_easyjson.go files
%_easyjson.go: %.go
	@easyjson $<

# Generate all *_easyjson.go files
.PHONY: generate
generate: $(easyjson_generated)

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
	$(MAKE) -C ./example clean

sdk_srcs := $(wildcard **/*.go)
wasm_src_dir := $(shell realpath ./test/testdata)
wasm_src := $(wasm_src_dir)/main.go
wasm_out := $(wasm_src_dir)/x.wasm

.PHONY: wasm
wasm: $(wasm_out)

$(wasm_out): $(wasm_src) $(sdk_srcs)
	@echo "Building WASM module..."
	@tinygo build -o $@ -target=wasi $<

test_scripts := $(shell realpath $(wildcard ./test/scripts/*.txtar))

.PHONY: test
test: $(wasm_out)
	@for script in $(test_scripts); do \
  		echo "Running test script: $$script..."; \
		testscript -e WASM_FILE=$(wasm_out) $$script; \
	done
