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
pre-pr: tidy generate lint test-integration

.PHONY: clean
clean:
	$(MAKE) -C ./example clean

sdk_srcs := $(wildcard **/*.go)
wasm_src_dir := ./test/integration/testdata
wasm_src := $(wasm_src_dir)/main.go
wasm_out := $(wasm_src_dir)/x.wasm

.PHONY: wasm
wasm: $(wasm_out)

$(wasm_out): $(wasm_src) $(sdk_srcs)
	@echo "Building WASM module..."
	@tinygo build -o $@ -target=wasi $<

.PHONY: test-integration
test-integration: $(wasm_out)
	@go test -v ./test/integration/... -count=1
