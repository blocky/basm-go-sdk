lint:
	golangci-lint run --config golangci.yaml

.PHONY: pre-pr
pre-pr: lint test-integration

generate:
	easyjson --all basm/dto.go
	easyjson x/xbasm/parse.go

srcs := $(wildcard **/*.go)
wasm_src_dir := ./test/integration/testdata
wasm_src := $(wasm_src_dir)/main.go
wasm_out := $(wasm_src_dir)/x.wasm

$(wasm_out): $(srcs)
	@echo "Building WASM module..."
	@tinygo build -o $(wasm_out) -target=wasi $(wasm_src)

.PHONY: test-integration
test-integration: $(wasm_out)
	go test -v ./test/integration/... -count=1
