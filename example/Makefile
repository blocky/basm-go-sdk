.PHONY: default run build

SRCS := $(wildcard ../*.go)

default: run

x.wasm: main.go $(SRCS)
	@echo "Building WASM module..."
	@tinygo build -o x.wasm -target=wasi main.go

build: x.wasm

run: build
	@echo "Running WASM module..."
	@cat fn-call.json \
		| BKY_AS_HOST="local-server" bky-as attest-fn-call \
		| jq -r '.transitive_attested_function_call | "Output:\n\t\(.claims.output | @base64d | fromjson | tojson)\nLogs:\n\t\(.logs | @base64d)"'
