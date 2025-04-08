.PHONY: default run build clean

SRCS := $(wildcard ../*.go)

default: run

x.wasm: main.go $(SRCS)
	@echo "Building WASM module..."
	@tinygo build -o x.wasm -target=wasi main.go

build: x.wasm

result.json: x.wasm
	@echo "Running WASM module..."
	@cat fn-call.json \
		| BKY_AS_HOST="local-server" bky-as attest-fn-call > result.json

run: result.json
	@echo "## Output"
	@jq -r '.transitive_attested_function_call.claims.output | @base64d | fromjson ' result.json
	@echo ""
	@echo "## Log"
	@jq -r '.transitive_attested_function_call.logs | @base64d' result.json

clean:
	rm -f x.wasm result.json
