compile:
	@echo "Building WASM module..."
	@tinygo build -o x.wasm -target=wasi main.go

run: compile
	@echo "Running WASM module..."
	@cat fn-call.json \
		| bky-as attest-fn-call \
		| jq -r '.transitive_attested_function_call | "Output:\n\t\(.claims.output | @base64d | fromjson | tojson)\nLogs:\n\t\(.logs | @base64d)"'
