compile:
	@echo "Building WASM module..."
	@tinygo build -o x.wasm -target=wasi main.go

run: compile
	@echo "Running WASM module..."
	@cat fn-call.json \
	| cli attest-fn-call \
	| jq -r '.function_calls[0].claims.output | @base64d | fromjson'
