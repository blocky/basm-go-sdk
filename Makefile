lint:
	golangci-lint run --config golangci.yaml

.PHONY: pre-pr
pre-pr: lint test-integration

generate:
	easyjson --all basm/dto.go
	easyjson x/xbasm/parse.go

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
	go test -v ./test/integration/... -count=1
