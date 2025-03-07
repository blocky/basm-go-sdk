check:
	@command -v docker >/dev/null 2>&1 || { echo >&2 "docker not found in PATH. Aborting."; exit 1; }
	@command -v bky-as >/dev/null 2>&1 || { echo >&2 "bky-as not found in PATH. Aborting."; exit 1; }
	@command -v jq >/dev/null 2>&1 || { echo >&2 "jq not found in PATH. Aborting."; exit 1; }

x.wasm: main.go | check
	@echo "Building WASM module..."
	@docker run --rm \
        -v .:/src \
        -w /src \
        tinygo/tinygo:0.31.2 \
        tinygo build -o x.wasm -target=wasi ./...

compile: x.wasm

run: compile
	@echo "Running WASM module..."
	@cat fn-call.json \
		| bky-as attest-fn-call \
		| jq -r '.function_calls[0] | "Output:\n\t\(.claims.output | @base64d | fromjson | tojson)\nLogs:\n\t\(.logs | @base64d)"'
